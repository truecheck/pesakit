package cli

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/pesakit/pesakit"
	"github.com/pesakit/pesakit/tigo"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/airtel/models"
	"github.com/techcraftlabs/mpesa"
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

func AirtelCallbackHandler() airtel.PushCallbackFunc {

	return func(request models.CallbackRequest) error {
		return nil
	}

}

func TigoCallbackHandler() tigopesa.CallbackHandlerFunc {
	return func(ctx context.Context, request tigopesa.CallbackRequest) (tigopesa.CallbackResponse, error) {
		response := tigopesa.CallbackResponse{
			ResponseCode:        tigopesa.SuccessCode,
			ResponseDescription: "successful",
			ResponseStatus:      true,
			ReferenceID:         request.ReferenceID,
		}

		return response, nil
	}
}

type (
	CallbackServer struct {
		server  *http.Server
		Port    string
		Handler http.HandlerFunc
	}
)

func NewCallbackServer(port, path string, handler http.HandlerFunc) *CallbackServer {

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
	return &CallbackServer{
		server: server,
		Port:   port,
	}
}

func run(client *pesakit.Client) clix.ActionFunc {
	return func(ctx *clix.Context) error {
		var (
			handlerFunc http.HandlerFunc
		)
		port, path := ctx.String("port"), ctx.String("path")
		mno := ctx.String("mno")
		if mno == "tigo" {
			handlerFunc = client.TigoPesa.CallbackServeHTTP
		}
		if mno == "airtel" {
			handlerFunc = client.AirtelMoney.CallbackServeHTTP
		}

		if mno == "vodacom" || mno == "mpesa" {
			handlerFunc = client.Mpesa.CallbackServeHTTP
		}
		cs := NewCallbackServer(port, path, handlerFunc)
		return cs.server.ListenAndServe()
	}

}

func callbackCommand(apiClient *pesakit.Client) *clix.Command {
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
		Action:  run(apiClient),
	}
}
