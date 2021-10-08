package pesakit

import (
	"context"
	"errors"
	"fmt"
	"github.com/techcraftlabs/pesakit/airtel"
	"github.com/techcraftlabs/pesakit/mpesa"
	"github.com/techcraftlabs/pesakit/pkg/mno"
	"github.com/techcraftlabs/pesakit/tigo"
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
		ID                    string `json:"id"`
		Amount                float64
		MSISDN                string
		Description           string
		ThirdPartyReferenceID string
		SubscriberCountry     string
		TransactionCountry    string
	}
	Action  int
	Service interface {
		Do(ctx context.Context, action Action, request Request) (interface{}, error)
	}
	Client struct {
		AirtelMoney *airtel.Client
		TigoPesa    *tigo.Client
		Mpesa       *mpesa.Client
	}
)

func (c *Client) Do(ctx context.Context, action Action, request Request) (interface{}, error) {
	ac := c.AirtelMoney
	tc := c.TigoPesa
	mp := c.Mpesa

	operator, fmtPhone, err := c.mnoAutoCheck(request.MSISDN)
	if err != nil {
		return nil, err
	}

	vodaSelectedButNil := mp == nil && operator == mno.Vodacom
	tigoSelectedButNil := tc == nil && operator == mno.Tigo
	airtSelectedButNil := ac == nil && operator == mno.Airtel

	selectedButNil := vodaSelectedButNil || tigoSelectedButNil || airtSelectedButNil

	if selectedButNil{
		return nil, ErrClientNil
	}

	switch action {
	case Push:
		switch operator {
		case mno.Tigo:
			req := tigo.PayRequest{
				CustomerMSISDN: fmtPhone,
				Amount: int64(request.Amount),
				Remarks:        request.Description,
				ReferenceID:    request.ID,
			}
			return tc.Push(ctx,req)

		case mno.Vodacom:
			req := mpesa.Request{
				ThirdPartyID: request.ThirdPartyReferenceID,
				Reference:    request.ID,
				Amount:       request.Amount,
				MSISDN:       request.MSISDN,
				Desc:         request.Description,
			}
			return mp.PushAsync(ctx,req)

		case mno.Airtel:
			req := airtel.PushPayRequest{
				Reference:          request.Description,
				SubscriberCountry:  request.SubscriberCountry,
				SubscriberMsisdn:   request.MSISDN,
				TransactionAmount: int64(request.Amount),
				TransactionCountry: request.TransactionCountry,
				TransactionID:      request.ID,
			}
			return ac.Push(ctx,req)

		default:
			return nil,fmt.Errorf("unknown action only push and disburse allowed for mpesa, tigopesa and airtel")
		}

	case Disburse:
		switch operator {
		case mno.Vodacom:
			req := mpesa.Request{
				ThirdPartyID: request.ThirdPartyReferenceID,
				Reference:    request.ID,
				Amount:       request.Amount,
				MSISDN:       request.MSISDN,
				Desc:         request.Description,
			}
			return mp.Disburse(ctx,req)
		case mno.Airtel:
			req := airtel.DisburseRequest{
				ID:                   request.ID,
				MSISDN:               request.MSISDN,
				Amount: int64(request.Amount),
				Reference:            request.Description,
				CountryOfTransaction: request.TransactionCountry,
			}
			return ac.Disburse(ctx,req)
		case mno.Tigo:
			req := tigo.Request{
				ReferenceID: request.ID,
				MSISDN:      request.MSISDN,
				Amount:      request.Amount,
			}
			return tc.Disburse(ctx, req)
		default:
			return nil,fmt.Errorf("unknown action only push and disburse allowed for mpesa, tigopesa and airtel")
		}

	default:
		return nil,fmt.Errorf("unknown action only push and disburse allowed for mpesa, tigopesa and airtel")
	}
}

func NewClient(airtelMoney *airtel.Client, tigopesa *tigo.Client, vodaMpesa *mpesa.Client) *Client {
	return &Client{
		AirtelMoney: airtelMoney,
		TigoPesa:    tigopesa,
		Mpesa:       vodaMpesa,
	}
}

func (c *Client) mnoAutoCheck(phone string) (mno.Operator, string, error) {
	op, fmtPhone, err := mno.Get(phone)
	if op == mno.Airtel{
		return op, fmtPhone[3:],err
	}
	return op,fmtPhone,err
}
