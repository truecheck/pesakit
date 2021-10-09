package pesakit

import (
	"context"
	"errors"
	"fmt"
	"github.com/pesakit/pesakit/airtel"
	"github.com/pesakit/pesakit/mpesa"
	"github.com/pesakit/pesakit/pkg/countries"
	"github.com/pesakit/pesakit/pkg/mno"
	"github.com/pesakit/pesakit/tigo"
)

var (
	_            Service = (*Client)(nil)
	ErrClientNil         = errors.New("the client is nil")
)

const (
	Push Action = iota
	Disburse
)

type (
	Request struct {
		ID                    string  `json:"id"`
		Amount                float64 `json:"amount"`
		MSISDN                string  `json:"msisdn"`
		Description           string  `json:"description"`
		thirdPartyReferenceID string
		subscriberCountry     string
		transactionCountry    string
	}

	RequestOption func(request *Request)

	Action  int
	Service interface {
		Do(ctx context.Context,operator mno.Operator, action Action, request Request, opts... RequestOption) (interface{}, error)
	}
	Client struct {
		AirtelMoney *airtel.Client
		TigoPesa    *tigo.Client
		Mpesa       *mpesa.Client
	}
)

func RequestThirdPartyReferenceID(ref string)RequestOption  {
	return func(request *Request) {
		request.thirdPartyReferenceID = ref
	}
}

func RequestSubscriberCountry(country string)RequestOption  {
	return func(request *Request) {
		request.subscriberCountry = country
	}
}

func RequestTransactionCountry(country string)RequestOption  {
	return func(request *Request) {
		request.transactionCountry = country
	}
}

func (c *Client) Do(ctx context.Context,operator mno.Operator, action Action, req Request, opts...RequestOption) (interface{}, error) {

	request := new(Request)
	request = &Request{
		ID:                    req.ID,
		Amount:                req.Amount,
		MSISDN:                req.MSISDN,
		Description:           req.Description,
		thirdPartyReferenceID: req.ID,
		subscriberCountry:     countries.TANZANIA,
		transactionCountry:    countries.TANZANIA,
	}
	for _, opt := range opts {
		opt(request)
	}
	ac := c.AirtelMoney
	tc := c.TigoPesa
	mp := c.Mpesa

	_ ,fmtPhone, err :=  MnoAutoCheck(request.MSISDN)

	if err != nil {
		return nil, err
	}

	mpesaSelectedButNil := (mp == nil) && (operator == mno.Vodacom)
	tigoSelectedButNil := (tc == nil) && (operator == mno.Tigo)
	airtelSelectedButNil := (ac == nil) && (operator == mno.Airtel)

	selectedButNil := mpesaSelectedButNil || tigoSelectedButNil || airtelSelectedButNil

	if selectedButNil {
		return nil, ErrClientNil
	}

	switch action {
	case Push:
		switch operator {
		case mno.Tigo:
			req := tigo.PayRequest{
				CustomerMSISDN: fmtPhone,
				Amount:         int64(request.Amount),
				Remarks:        request.Description,
				ReferenceID:    request.ID,
			}
			return tc.Push(ctx, req)

		case mno.Vodacom:
			req := mpesa.Request{
				ThirdPartyID: request.thirdPartyReferenceID,
				Reference:    request.ID,
				Amount:       request.Amount,
				MSISDN:       request.MSISDN,
				Desc:         request.Description,
			}
			return mp.PushAsync(ctx, req)

		case mno.Airtel:
			req := airtel.PushPayRequest{
				Reference:          request.Description,
				SubscriberCountry:  request.subscriberCountry,
				SubscriberMsisdn:   request.MSISDN,
				TransactionAmount:  int64(request.Amount),
				TransactionCountry: request.transactionCountry,
				TransactionID:      request.ID,
			}
			return ac.Push(ctx, req)

		default:
			return nil, fmt.Errorf("unknown action only push and disburse allowed for mpesa, tigopesa and airtel")
		}

	case Disburse:
		switch operator {
		case mno.Vodacom:
			req := mpesa.Request{
				ThirdPartyID: request.thirdPartyReferenceID,
				Reference:    request.ID,
				Amount:       request.Amount,
				MSISDN:       request.MSISDN,
				Desc:         request.Description,
			}
			return mp.Disburse(ctx, req)
		case mno.Airtel:
			req := airtel.DisburseRequest{
				ID:                   request.ID,
				MSISDN:               request.MSISDN,
				Amount:               int64(request.Amount),
				Reference:            request.Description,
				CountryOfTransaction: request.transactionCountry,
			}
			return ac.Disburse(ctx, req)
		case mno.Tigo:
			req := tigo.Request{
				ReferenceID: request.ID,
				MSISDN:      request.MSISDN,
				Amount:      request.Amount,
			}
			return tc.Disburse(ctx, req)
		default:
			return nil, fmt.Errorf("unknown action only push and disburse allowed for mpesa, tigopesa and airtel")
		}

	default:
		return nil, fmt.Errorf("unknown action only push and disburse allowed for mpesa, tigopesa and airtel")
	}
}

func NewClient(airtelMoney *airtel.Client, tigopesa *tigo.Client, vodaMpesa *mpesa.Client) *Client {
	return &Client{
		AirtelMoney: airtelMoney,
		TigoPesa:    tigopesa,
		Mpesa:       vodaMpesa,
	}
}

func MnoAutoCheck(phone string) (mno.Operator, string, error) {
	op, fmtPhone, err := mno.Get(phone)
	if op == mno.Airtel {
		return op, fmtPhone[3:], err
	}
	return op, fmtPhone, err
}
