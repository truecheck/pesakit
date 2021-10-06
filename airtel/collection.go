package airtel

import (
	"context"
	"fmt"
	"github.com/techcraftlabs/pesakit/internal"
	"github.com/techcraftlabs/pesakit/pkg/countries"
	"net/http"
)

type CollectionService interface {
	Push(ctx context.Context, request PushPayRequest) (PushPayResponse, error)
	Refund(ctx context.Context, request RefundRequest) (RefundResponse, error)
	PushEnquiry(ctx context.Context, request PushEnquiryRequest) (PushEnquiryResponse, error)
	CallbackServeHTTP(writer http.ResponseWriter, request *http.Request)
}

func (c *Client) Push(ctx context.Context, request PushPayRequest) (PushPayResponse, error) {

	pushRequest := c.reqAdapter.ToPushPayRequest(request)
	pushResponse, err := c.push(ctx, pushRequest)
	if err != nil {
		return PushPayResponse{}, err
	}
	response := c.resAdapter.ToPushPayResponse(pushResponse)
	return response, nil
}

func (c *Client) push(ctx context.Context, request PushRequest) (iPushResponse, error) {
	token, err := c.checkToken(ctx)
	if err != nil {
		return iPushResponse{}, err
	}

	transaction := request.Transaction
	countryCodeName := transaction.Country
	currencyCodeName := transaction.Currency

	var opts []internal.RequestOption

	hs := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"X-Country":     countryCodeName,
		"X-Currency":    currencyCodeName,
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	headersOpt := internal.WithRequestHeaders(hs)
	opts = append(opts, headersOpt)
	pe := c.Conf.Endpoints.PushEndpoint
	req := internal.MakeInternalRequest(c.baseURL, pe, UssdPush, request, opts...)
	res := new(iPushResponse)
	_, err = c.base.Do(ctx, UssdPush.String(), req, res)
	if err != nil {
		return iPushResponse{}, err
	}
	return *res, nil
}

func (c *Client) Refund(ctx context.Context, request RefundRequest) (RefundResponse, error) {
	country, err := countries.GetByName(request.CountryOfTransaction)
	if err != nil {
		return RefundResponse{}, err
	}
	token, err := c.checkToken(ctx)
	if err != nil {
		return RefundResponse{}, err
	}
	var opts []internal.RequestOption
	hs := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"X-Country":     country.CodeName,
		"X-Currency":    country.CurrencyCode,
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	headersOpt := internal.WithRequestHeaders(hs)
	opts = append(opts, headersOpt)

	req := c.makeInternalRequest(Refund, request, opts...)

	if err != nil {
		return RefundResponse{}, err
	}

	res := new(RefundResponse)
	//env := c.Conf.Environment
	rn := Refund.String()
	_, err = c.base.Do(ctx, rn, req, res)
	if err != nil {
		return RefundResponse{}, err
	}
	return *res, nil

}

func (c *Client) PushEnquiry(ctx context.Context, request PushEnquiryRequest) (PushEnquiryResponse, error) {

	country, err := countries.GetByName(request.CountryOfTransaction)
	if err != nil {
		return PushEnquiryResponse{}, err
	}
	token, err := c.checkToken(ctx)
	if err != nil {
		return PushEnquiryResponse{}, err
	}
	var opts []internal.RequestOption
	hs := map[string]string{
		"Content-Type":  "application/json",
		"Accept":        "*/*",
		"X-Country":     country.CodeName,
		"X-Currency":    country.CurrencyCode,
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}
	headersOpt := internal.WithRequestHeaders(hs)
	endpointOpt := internal.WithEndpoint(request.ID)
	opts = append(opts, headersOpt, endpointOpt)
	req := c.makeInternalRequest(PushEnquiry, request, opts...)
	if err != nil {
		return PushEnquiryResponse{}, err
	}
	reqName := PushEnquiry.String()
	res := new(PushEnquiryResponse)
	_, err = c.base.Do(ctx, reqName, req, res)
	if err != nil {
		return PushEnquiryResponse{}, err
	}
	return *res, nil
}

func (c *Client) CallbackServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body := new(CallbackRequest)
	err := internal.ReceivePayload(request, body)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	reqBody := *body

	//todo: check the hash if it is OK
	err = c.pushCallbackFunc.Handle(reqBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
