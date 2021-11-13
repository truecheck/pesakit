package pesakit

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/techcraftlabs/airtel"

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
		Addr:    port,
		Handler: router,
	}
	return &callbackServer{
		server: server,
		Port:   port,
	}
}

func run(client *Client) clix.ActionFunc {
	return func(ctx *clix.Context) error {
		var (
			handlerFunc http.HandlerFunc
		)
		port, path := ctx.String("port"), ctx.String("path")
		mno := ctx.String("mno")
		if mno == "tigo" {
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
