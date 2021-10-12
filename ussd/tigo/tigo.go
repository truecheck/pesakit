package tigo

import (
	"context"
	"encoding/xml"
	"github.com/pesakit/pesakit/internal"
	"net/http"
	"time"
)

const (
	syncLookupResponse  = "SYNC_LOOKUP_RESPONSE"
	syncBillPayResponse = "SYNC_BILLPAY_RESPONSE"

	//NameQuery Error Codes

	ErrNameNotRegistered = "error010"
	ErrNameInvalidFormat = "error030"
	ErrNameUserSuspended = "error030"
	NoNamecheckErr       = "error000"

	//WalletToAccount error codes

	ErrSuccessTxn               = "error000"
	ErrServiceNotAvailable      = "error001"
	ErrInvalidCustomerRefNumber = "error010"
	ErrCustomerRefNumLocked     = "error011"
	ErrInvalidAmount            = "error012"
	ErrAmountInsufficient       = "error013"
	ErrAmountTooHigh            = "error014"
	ErrAmountTooLow             = "error015"
	ErrInvalidPayment           = "error016"
	ErrGeneralError             = "error100"
	ErrRetryConditionNoResponse = "error111"
)

var (
	_ PaymentHandler   = (*PaymentHandleFunc)(nil)
	_ NameQueryHandler = (*NameQueryFunc)(nil)
	_ Service          = (*Client)(nil)
)

type (
	Service interface {
		NameQueryServeHTTP(writer http.ResponseWriter, request *http.Request)
		PaymentServeHTTP(writer http.ResponseWriter, request *http.Request)
	}

	PaymentHandler interface {
		PaymentRequest(ctx context.Context, request PayRequest) (PayResponse, error)
	}

	PaymentHandleFunc func(ctx context.Context, request PayRequest) (PayResponse, error)

	NameQueryHandler interface {
		NameQuery(ctx context.Context, request NameRequest) (NameResponse, error)
	}

	NameQueryFunc func(ctx context.Context, request NameRequest) (NameResponse, error)

	nameRequest struct {
		XMLName             xml.Name `xml:"COMMAND"`
		Text                string   `xml:",chardata"`
		Type                string   `xml:"TYPE"`
		Msisdn              string   `xml:"MSISDN"`
		CompanyName         string   `xml:"COMPANYNAME"`
		CustomerReferenceID string   `xml:"CUSTOMERREFERENCEID"`
	}

	NameRequest struct {
		Msisdn              string `xml:"MSISDN"`
		CompanyName         string `xml:"COMPANYNAME"`
		CustomerReferenceID string `xml:"CUSTOMERREFERENCEID"`
	}

	NameResponse struct {
		Result    string `xml:"RESULT"`
		ErrorCode string `xml:"ERRORCODE"`
		ErrorDesc string `xml:"ERRORDESC"`
		Msisdn    string `xml:"MSISDN"`
		Flag      string `xml:"FLAG"`
		Content   string `xml:"CONTENT"`
	}

	nameResponse struct {
		XMLName   xml.Name `xml:"COMMAND"`
		Text      string   `xml:",chardata"`
		Type      string   `xml:"TYPE"`
		Result    string   `xml:"RESULT"`
		ErrorCode string   `xml:"ERRORCODE"`
		ErrorDesc string   `xml:"ERRORDESC"`
		Msisdn    string   `xml:"MSISDN"`
		Flag      string   `xml:"FLAG"`
		Content   string   `xml:"CONTENT"`
	}

	PayRequest struct {
		TxnID               string  `xml:"TXNID"`
		Msisdn              string  `xml:"MSISDN"`
		Amount              float64 `xml:"AMOUNT"`
		CompanyName         string  `xml:"COMPANYNAME"`
		CustomerReferenceID string  `xml:"CUSTOMERREFERENCEID"`
		SenderName          string  `xml:"SENDERNAME"`
	}

	payRequest struct {
		XMLName             xml.Name `xml:"COMMAND"`
		Text                string   `xml:",chardata"`
		TYPE                string   `xml:"TYPE"`
		TxnID               string   `xml:"TXNID"`
		Msisdn              string   `xml:"MSISDN"`
		Amount              float64  `xml:"AMOUNT"`
		CompanyName         string   `xml:"COMPANYNAME"`
		CustomerReferenceID string   `xml:"CUSTOMERREFERENCEID"`
		SenderName          string   `xml:"SENDERNAME"`
	}

	PayResponse struct {
		TxnID            string `xml:"TXNID"`
		RefID            string `xml:"REFID"`
		Result           string `xml:"RESULT"`
		ErrorCode        string `xml:"ERRORCODE"`
		ErrorDescription string `xml:"ERRORDESCRIPTION"`
		Msisdn           string `xml:"MSISDN"`
		Flag             string `xml:"FLAG"`
		Content          string `xml:"CONTENT"`
	}

	payResponse struct {
		XMLName          xml.Name `xml:"COMMAND"`
		Text             string   `xml:",chardata"`
		Type             string   `xml:"TYPE"`
		TxnID            string   `xml:"TXNID"`
		RefID            string   `xml:"REFID"`
		Result           string   `xml:"RESULT"`
		ErrorCode        string   `xml:"ERRORCODE"`
		ErrorDescription string   `xml:"ERRORDESCRIPTION"`
		Msisdn           string   `xml:"MSISDN"`
		Flag             string   `xml:"FLAG"`
		Content          string   `xml:"CONTENT"`
	}

	Config struct {
		AccountName   string
		AccountMSISDN string
		BillerNumber  string
		RequestURL    string
		NameCheckURL  string
	}

	Client struct {
		rv internal.Receiver
		rp internal.Replier
		base *internal.BaseClient
		*Config
		PaymentHandler   PaymentHandler
		NameQueryHandler NameQueryHandler
	}
)

func transformNameRequest(request nameRequest) NameRequest {
	return NameRequest{
		Msisdn:              request.Msisdn,
		CompanyName:         request.CompanyName,
		CustomerReferenceID: request.CustomerReferenceID,
	}
}

func transformPayRequest(request payRequest) PayRequest {
	return PayRequest{
		TxnID:               request.TxnID,
		Msisdn:              request.Msisdn,
		Amount:              request.Amount,
		CompanyName:         request.CompanyName,
		CustomerReferenceID: request.CustomerReferenceID,
		SenderName:          request.SenderName,
	}
}

func transformToXMLNameResponse(response NameResponse) nameResponse {
	return nameResponse{
		Type:      syncLookupResponse,
		Result:    response.Result,
		ErrorCode: response.ErrorCode,
		ErrorDesc: response.ErrorDesc,
		Msisdn:    response.Msisdn,
		Flag:      response.Flag,
		Content:   response.Content,
	}
}

func transformToXMLPayResponse(response PayResponse) payResponse {
	return payResponse{
		Type:             syncBillPayResponse,
		TxnID:            response.TxnID,
		RefID:            response.RefID,
		Result:           response.Result,
		ErrorCode:        response.ErrorCode,
		ErrorDescription: response.ErrorDescription,
		Msisdn:           response.Msisdn,
		Flag:             response.Flag,
		Content:          response.Content,
	}
}

func NewClient(config *Config, handler PaymentHandler, queryHandler NameQueryHandler, opts ...ClientOption) *Client {
	client := &Client{
		Config:           config,
		PaymentHandler:   handler,
		NameQueryHandler: queryHandler,
		base:             internal.NewBaseClient(),
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (client *Client)SetNameQueryHandler(nameQueryHandler NameQueryHandler){
	client.NameQueryHandler = nameQueryHandler
}

func (client *Client)SetPaymentHandler(paymentHandler PaymentHandler){
	client.PaymentHandler = paymentHandler
}

func (handler PaymentHandleFunc) PaymentRequest(ctx context.Context, request PayRequest) (PayResponse, error) {
	return handler(ctx, request)
}

func (handler NameQueryFunc) NameQuery(ctx context.Context, request NameRequest) (NameResponse, error) {
	return handler(ctx, request)
}

func (client *Client) NameQueryServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()
	var req nameRequest

	err := internal.ReceivePayload(request, &req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	response, err := client.NameQueryHandler.NameQuery(ctx, transformNameRequest(req))

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	var opts []internal.ResponseOption
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	headersOpts := internal.WithResponseHeaders(headers)
	opts = append(opts, headersOpts)

	payload := transformToXMLNameResponse(response)
	res := internal.NewResponse(200, payload, opts...)

	internal.Reply(writer, res)

}

func (client *Client) PaymentServeHTTP(writer http.ResponseWriter, request *http.Request) {

	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Second)
	defer cancel()

	var req payRequest

	err := internal.ReceivePayload(request, &req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := client.PaymentHandler.PaymentRequest(ctx, transformPayRequest(req))

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	var opts []internal.ResponseOption
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	headersOpts := internal.WithResponseHeaders(headers)
	opts = append(opts, headersOpts)
	payload := transformToXMLPayResponse(response)
	res := internal.NewResponse(200, payload, opts...)

	internal.Reply(writer, res)
}
