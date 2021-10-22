package tigo

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/techcraftlabs/base"
	"math"
	"net/http"
	"net/url"
	"time"
)

const (
	SuccessCode                 = "BILLER-30-0000-S"
	FailureCode                 = "BILLER-30-3030-E"
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
	requestType                 = "REQMFCI"
	senderLanguage              = "EN"
)

type (
	CallbackHandler interface {
		Respond(ctx context.Context, request CallbackRequest) (CallbackResponse, error)
	}
	CallbackHandlerFunc func(context.Context, CallbackRequest) (CallbackResponse, error)

	disburseRequest struct {
		XMLName     xml.Name `xml:"COMMAND"`
		Text        string   `xml:",chardata"`
		Type        string   `xml:"TYPE"`
		ReferenceID string   `xml:"REFERENCEID"`
		Msisdn      string   `xml:"MSISDN"`
		PIN         string   `xml:"PIN"`
		Msisdn1     string   `xml:"MSISDN1"`
		Amount      float64  `xml:"AMOUNT"`
		SenderName  string   `xml:"SENDERNAME"`
		Language1   string   `xml:"LANGUAGE1"`
		BrandID     string   `xml:"BRAND_ID"`
	}

	Response struct {
		XMLName     xml.Name `xml:"COMMAND" json:"-"`
		Text        string   `xml:",chardata" json:"-"`
		Type        string   `xml:"TYPE" json:"type"`
		ReferenceID string   `xml:"REFERENCEID" json:"reference,omitempty"`
		TxnID       string   `xml:"TXNID" json:"id,omitempty"`
		TxnStatus   string   `xml:"TXNSTATUS" json:"status,omitempty"`
		Message     string   `xml:"MESSAGE" json:"message,omitempty"`
	}

	TokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int64  `json:"expires_in"`
	}

	Request struct {
		MSISDN      string  `json:"msisdn"`
		Amount      float64 `json:"amount"`
		Remarks     string  `json:"remarks,omitempty"`
		ReferenceID string  `json:"referenceID"`
	}

	// payRequest This is the request expected by tigo with BillerMSISDN hooked up
	// from DisburseConfig
	// Request is used by DisburseClient without the need to specify BillerMSISDN
	payRequest struct {
		CustomerMSISDN string  `json:"CustomerMSISDN"`
		BillerMSISDN   string  `json:"BillerMSISDN"`
		Amount         float64 `json:"Amount"`
		Remarks        string  `json:"Remarks,omitempty"`
		ReferenceID    string  `json:"ReferenceID"`
	}

	PayResponse struct {
		ResponseCode        string `json:"ResponseCode"`
		ResponseStatus      bool   `json:"ResponseStatus"`
		ResponseDescription string `json:"ResponseDescription"`
		ReferenceID         string `json:"ReferenceID"`
		Message             string `json:"Message,omitempty"`
	}

	CallbackRequest struct {
		Status           bool   `json:"Status"`
		Description      string `json:"Description"`
		MFSTransactionID string `json:"MFSTransactionID,omitempty"`
		ReferenceID      string `json:"ReferenceID"`
		Amount           string `json:"Amount"`
	}

	CallbackResponse struct {
		ResponseCode        string `json:"ResponseCode"`
		ResponseDescription string `json:"ResponseDescription"`
		ResponseStatus      bool   `json:"ResponseStatus"`
		ReferenceID         string `json:"ReferenceID"`
	}

	DisburseConfig struct {
		AccountName   string
		AccountMSISDN string
		BrandID       string
		PIN           string
		RequestURL    string
	}

	PushConfig struct {
		Username          string
		Password          string
		PasswordGrantType string
		BaseURL           string
		TokenEndpoint     string
		PushPayEndpoint   string
		BillerMSISDN      string
		BillerCode        string
	}

	Config struct {
		*DisburseConfig
		*PushConfig
	}

	Client struct {
		*Config
		rp base.Replier
		rv base.Receiver
		base            *base.Client
		CallbackHandler CallbackHandler
		token           *string
		tokenExpires    time.Time
	}

	Service interface {
		Token(ctx context.Context) (TokenResponse, error)
		Push(ctx context.Context, request Request) (PayResponse, error)
		CallbackServeHTTP(writer http.ResponseWriter, r *http.Request)
		Disburse(ctx context.Context, request Request) (Response, error)
	}
)

func NewClient(config *Config, opts ...ClientOption) *Client {
	token := new(string)
	client := &Client{
		Config:       config,
		token:        token,
		tokenExpires: time.Now(),
		base:         base.NewClient(),
	}

	for _, opt := range opts {
		opt(client)
	}

	lg := client.base.Logger
	dm := client.base.DebugMode

	client.rp = base.NewReplier(lg, dm)
	client.rv = base.NewReceiver(lg,dm)

	return client
}

func (c *Client) SetCallbackHandler(handler CallbackHandler) {
	c.CallbackHandler = handler
}

func (handler CallbackHandlerFunc) Respond(ctx context.Context, request CallbackRequest) (CallbackResponse, error) {
	return handler(ctx, request)
}

func (c *Client) Push(ctx context.Context, request Request) (response PayResponse, err error) {
	amount := math.Floor(request.Amount * 100 / 100)
	var billPayReq = payRequest{
		CustomerMSISDN: request.MSISDN,
		BillerMSISDN:   c.BillerMSISDN,
		Amount:         amount,
		Remarks:        request.Remarks,
		ReferenceID:    fmt.Sprintf("%s%s", c.PushConfig.BillerCode, request.ReferenceID),
	}

	token, err := c.checkToken(ctx)
	if err != nil {
		return PayResponse{}, err
	}

	authHeader := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", token),
		"Username":      c.PushConfig.Username,
		"Password":      c.PushConfig.Password,
	}
	var requestOpts []base.RequestOption
	moreHeaderOpt := base.WithMoreHeaders(authHeader)
	//basicAuth := base.WithBasicAuth(c.PushConfig.Username, c.PushConfig.Password)
	requestOpts = append(requestOpts, moreHeaderOpt)

	req := base.MakeInternalRequest(c.BaseURL, c.PushConfig.PushPayEndpoint, PushPay, billPayReq, requestOpts...)

	rn := PushPay.String()
	_, err = c.base.Do(context.TODO(), rn, req, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

func (c *Client) CallbackServeHTTP(w http.ResponseWriter, r *http.Request) {
	callbackRequest := new(CallbackRequest)
	statusCode := 200

	err := base.ReceivePayload(r, callbackRequest)
	if err != nil {
		statusCode = http.StatusInternalServerError
		http.Error(w, err.Error(), statusCode)
		return
	}
	req := *callbackRequest

	callbackResponse, err := c.CallbackHandler.Respond(context.TODO(), req)

	if err != nil {
		statusCode = http.StatusInternalServerError
		http.Error(w, err.Error(), statusCode)
		return
	}

	var responseOpts []base.ResponseOption
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	headersOpt := base.WithResponseHeaders(headers)

	responseOpts = append(responseOpts, headersOpt, base.WithResponseError(err))
	response := base.NewResponse(statusCode, callbackResponse, responseOpts...)
	c.rp.Reply(w, response)

}

func (c *Client) Disburse(ctx context.Context, request Request) (response Response, err error) {
	amount := math.Floor(request.Amount * 100 / 100)
	var reqOpts []base.RequestOption
	headers := map[string]string{
		"Content-Type": "application/xml",
	}
	headersOpt := base.WithRequestHeaders(headers)
	reqOpts = append(reqOpts, headersOpt)
	r := disburseRequest{
		Type:        requestType,
		ReferenceID: request.ReferenceID,
		Msisdn:      c.DisburseConfig.AccountMSISDN,
		PIN:         c.DisburseConfig.PIN,
		Msisdn1:     request.MSISDN,
		Amount:      amount,
		SenderName:  c.DisburseConfig.AccountName,
		Language1:   senderLanguage,
		BrandID:     c.DisburseConfig.BrandID,
	}

	req := base.MakeInternalRequest(c.RequestURL, "", Disburse, r, reqOpts...)

	_, err = c.base.Do(ctx, Disburse.String(), req, &response)

	if err != nil {
		return
	}

	return
}

func (c *Client) checkToken(ctx context.Context) (string, error) {
	var token string
	if *c.token == "" {
		str, err := c.Token(ctx)
		if err != nil {
			return "", err
		}
		token = fmt.Sprintf("%s", str.AccessToken)
	}
	//Add Auth Header
	if *c.token != "" {
		if !c.tokenExpires.IsZero() && time.Until(c.tokenExpires) < (60*time.Second) {
			if _, err := c.Token(ctx); err != nil {
				return "", err
			}
		}
		token = *c.token
	}

	return token, nil
}

func (c *Client) Token(ctx context.Context) (TokenResponse, error) {
	var form = url.Values{}
	form.Set("username", c.Username)
	form.Set("password", c.Password)
	form.Set("grant_type", c.PasswordGrantType)

	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Cache-Control": "no-cache",
	}

	var requestOptions []base.RequestOption
	headersOption := base.WithRequestHeaders(headers)
	requestOptions = append(requestOptions, headersOption)

	request := base.MakeInternalRequest(c.PushConfig.BaseURL, c.PushConfig.TokenEndpoint, GetToken, form, requestOptions...)

	var tokenResponse TokenResponse

	rn := GetToken.String()

	_, err := c.base.Do(context.TODO(), rn, request, &tokenResponse)

	if err != nil {
		return TokenResponse{}, err
	}

	token := tokenResponse.AccessToken

	c.token = &token

	//This set the value to when a new token will set above will be expired
	//the minus 10 is an overhead a margin for error.
	tokenExpiresAt := time.Now().Add(time.Duration(tokenResponse.ExpiresIn-10) * time.Second)
	c.tokenExpires = tokenExpiresAt

	return tokenResponse, nil

}
