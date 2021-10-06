package airtel

import (
	"context"
	"net/http"
)

type (
	Service interface {
		Push(ctx context.Context, request PushPayRequest) (PushPayResponse, error)
		Disburse(ctx context.Context, request DisburseRequest) (DisburseResponse, error)
		Summary(ctx context.Context, params Params) (ListTransactionsResponse, error)
		CallbackServeHTTP(writer http.ResponseWriter, request *http.Request)
		Token(ctx context.Context) (TokenResponse, error)
	}
)
