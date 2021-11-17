package pesakit

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/manifoldco/promptui"
	"github.com/techcraftlabs/airtel"
	"strconv"
	"strings"

	//"github.com/techcraftlabs/airtel"
	//	"github.com/techcraftlabs/airtel/models"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa/push"
	clix "github.com/urfave/cli/v2"
	"net/http"
)

func MpesaCallbackHandler() mpesa.PushCallbackFunc {
	return func(request mpesa.PushCallbackRequest) (mpesa.PushCallbackResponse, error) {
		response := mpesa.PushCallbackResponse{
			OriginalConversationID:   request.OriginalConversationID,
			ResponseCode:             request.ResultCode,
			ResponseDesc:             request.ResultDesc,
			ThirdPartyConversationID: request.ThirdPartyConversationID,
		}

		return response, nil
	}
}

func AirtelCallbackHandler()airtel.PushCallbackFunc{
	return func(request airtel.CallbackRequest) error {
		return nil
	}
}

func TigoCallbackHandler() push.CallbackHandlerFunc {
	return func(ctx context.Context, request push.CallbackRequest) (push.CallbackResponse, error) {
		response := push.CallbackResponse{
			ResponseCode:        push.SuccessCode,
			ResponseDescription: "successful",
			ResponseStatus:      true,
			ReferenceID:         request.ReferenceID,
		}

		return response, nil
	}
}

type (
	callbackServer struct {
		server  *http.Server
		Port    string
		Handler http.HandlerFunc
	}
)

func newCallbackServer(port, path string, handler http.HandlerFunc) *callbackServer {

	handle := func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		handler(writer, request)
		return
	}
	router := httprouter.New()

	router.POST(path, handle)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s",port),
		Handler: router,
	}
	return &callbackServer{
		server: server,
		Port:   port,
	}
}

func validateMno(input string)error{
	names := [...]string{
		"airtel",
        "tigo",
        "vodacom",
        "tigopesa",
        "mpesa",
	}

	// check if input matches any string in the names
	for _, name := range names {
        if input == name {
            return nil
        }
    }


	//concatenate all values in the names slice into msg string
	msg := strings.Join(names[:], ", ")
	return fmt.Errorf("%s is not a valid mno. Valid mnos are %s", input, msg)
}

func validatePort(input string)error{
	// check if input is int64
	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil{
		return err
	}
	return nil
}

func run(client *Client) clix.ActionFunc {
	promptMno:= promptui.Prompt{
		Label:       "mno",
		Validate: validateMno,
	}

	promptPort:= promptui.Prompt{
		Label:       "port",
		Validate: validatePort,
	}

	promptPath:= promptui.Prompt{
		Label:       "path",
		Validate: validateNil,
	}
	return func(ctx *clix.Context) error {
		var (
			mno string
			port string
			path string
			err error
			handlerFunc http.HandlerFunc
		)

		mnoIsSet := ctx.IsSet("mno")
		portIsSet := ctx.IsSet("port")
		pathIsSet := ctx.IsSet("path")

		allIsSet := mnoIsSet && portIsSet && pathIsSet

		if allIsSet {
			mno, port, path =ctx.String("mno"), ctx.String("port"), ctx.String("path")
        }else {
			if mnoIsSet {
                mno = ctx.String("mno")
            }else {
				mno, err = promptMno.Run()
				if err != nil{
					return err
				}
			}

			if portIsSet {
                port = ctx.String("port")
            }else {
                port, err = promptPort.Run()
                if err != nil{
                    return err
                }
            }

			if pathIsSet {
				path = ctx.String("path")
            }else {
                path, err = promptPath.Run()
                if err != nil{
                    return err
                }
			}
		}

		if mno == "tigo" || mno == "tigopesa" {
			handlerFunc = client.tigo.CallbackServeHTTP
		}
		if mno == "airtel" {
			handlerFunc = client.airtel.CallbackServeHTTP
		}

		if mno == "vodacom" || mno == "mpesa" {
			handlerFunc = client.mpesa.CallbackServeHTTP
		}
		cs := newCallbackServer(port, path, handlerFunc)
		return cs.server.ListenAndServe()
	}

}

func (c *Client)callbackCommand() *clix.Command {
	flags := []clix.Flag{
		&clix.StringFlag{
			Name:  "mno",
			Usage: "mobile money provider (tigo, airtel, vodacom)",
		},
		&clix.StringFlag{
			Name:  "port",
			Usage: "callback listener port",
		},
		&clix.StringFlag{
			Name:  "path",
			Usage: "callback path",
		},
	}

	return &clix.Command{
		Name:    "callbacks",
		Aliases: []string{"cb"},
		Usage:   "monitor callbacks from mno",
		Flags:   flags,
		Action:  run(c),
	}
}
