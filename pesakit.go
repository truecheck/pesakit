package pesakit

import (
	"context"
	"errors"
	"fmt"
	"github.com/pesakit/pesakit/pkg/countries"
	"github.com/pesakit/pesakit/pkg/mno"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/base"
	"github.com/techcraftlabs/base/io"
	"github.com/techcraftlabs/mpesa"
	tigo "github.com/techcraftlabs/tigopesa"
	"github.com/techcraftlabs/tigopesa/disburse"
	"github.com/techcraftlabs/tigopesa/push"
	clix "github.com/urfave/cli/v2"
	stdio "io"
)

var (
	_            doer = (*Client)(nil)
	ErrClientNil      = errors.New("the client is nil")
)

const (
	pushAction action = iota
	disburseAction
)

type (
	request struct {
		ID                    string  `json:"id"`
		Amount                float64 `json:"amount"`
		MSISDN                string  `json:"msisdn"`
		Description           string  `json:"description"`
		thirdPartyReferenceID string
		subscriberCountry     string
		transactionCountry    string
	}

	requestOption func(request *request)

	action int
	doer   interface {
		do(ctx context.Context, action action, request *request) (interface{}, error)
	}
	Client struct {
		rv      base.Receiver
		cli     *clix.App
		verbose bool
		logger  stdio.Writer
		debug   bool
		airtel  *airtel.Client
		tigo    *tigo.Client
		mpesa   *mpesa.Client
		configFile string
	}

	ClientOption func(*Client)
)

func Logger(writer stdio.Writer) ClientOption {
    return func(c *Client) {
		c.logger = writer
		c.rv = base.NewReceiver(c.logger, c.debug)
	}
}

func Debug(debug bool) ClientOption {
    return func(c *Client) {
        c.debug = debug
		c.rv = base.NewReceiver(c.logger, c.debug)
    }
}

func Verbose(verbose bool) ClientOption {
    return func(c *Client) {
        c.verbose = verbose
    }
}

func makeRequest(id string, amount float64, msisdn string, desc string, opts ...requestOption) *request {
	r := &request{
		ID:                    id,
		Amount:                amount,
		MSISDN:                msisdn,
		Description:           desc,
		thirdPartyReferenceID: id,
		subscriberCountry:     countries.TANZANIA,
		transactionCountry:    countries.TANZANIA,
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func requestThirdPartyReferenceID(ref string) requestOption {
	return func(request *request) {
		request.thirdPartyReferenceID = ref
	}
}

func requestSubscriberCountry(country string) requestOption {
	return func(request *request) {
		request.subscriberCountry = country
	}
}

func requestTransactionCountry(country string) requestOption {
	return func(request *request) {
		request.transactionCountry = country
	}
}

func (c *Client) do(ctx context.Context, action action, req *request) (interface{}, error) {
	ac := c.airtel
	tc := c.tigo
	mp := c.mpesa

	operator, fmtPhone, err := mnoAutoCheck(req.MSISDN)
	if err != nil {
		return nil, err
	}

	req.MSISDN = fmtPhone

	mpesaSelectedButNil := (mp == nil) && (operator == mno.Vodacom)
	tigoSelectedButNil := (tc == nil) && (operator == mno.Tigo)
	airtelSelectedButNil := (ac == nil) && (operator == mno.Airtel)

	selectedButNil := mpesaSelectedButNil || tigoSelectedButNil || airtelSelectedButNil

	if selectedButNil {
		return nil, ErrClientNil
	}

	switch action {
	case pushAction:
		switch operator {
		case mno.Vodacom:
			re := mpesa.Request{
				ThirdPartyID: req.thirdPartyReferenceID,
				Reference:    req.ID,
				Amount:       req.Amount,
				MSISDN:       req.MSISDN,
				Description:  req.Description,
			}
			return c.mpesa.PushAsync(ctx, re)
		case mno.Tigo:
			re := push.Request{
				MSISDN:      req.MSISDN,
				Amount:      req.Amount,
				Remarks:     req.Description,
				ReferenceID: req.ID,
			}
			return c.tigo.Pay(ctx, re)
		case mno.Airtel:
			re := airtel.PushPayRequest{
				Reference:          req.Description,
				SubscriberCountry:  req.subscriberCountry,
				SubscriberMsisdn:   req.MSISDN,
				TransactionAmount:  req.Amount,
				TransactionCountry: req.transactionCountry,
				TransactionID:      req.ID,
			}
			return c.airtel.Push(ctx, re)
		}

	case disburseAction:
		switch operator {
		case mno.Vodacom:
			re := mpesa.Request{
				ThirdPartyID: req.thirdPartyReferenceID,
				Reference:    req.ID,
				Amount:       req.Amount,
				MSISDN:       req.MSISDN,
				Description:  req.Description,
			}
			return c.mpesa.Disburse(ctx, re)

		case mno.Tigo:
			re := disburse.Request{
				MSISDN:      req.MSISDN,
				Amount:      req.Amount,
				ReferenceID: req.ID,
			}
			return c.tigo.Disburse(ctx, re)

		case mno.Airtel:
			re := airtel.DisburseRequest{
				ID:                   req.ID,
				MSISDN:               req.MSISDN,
				Amount:               req.Amount,
				Reference:            req.Description,
				CountryOfTransaction: req.transactionCountry,
			}
			return c.airtel.Disburse(ctx, re)
		}
	}

	return nil, fmt.Errorf("error: bad request")
}

func NewClient(airtelMoney *airtel.Client, tc *tigo.Client, vodaMpesa *mpesa.Client, opts...ClientOption) *Client {

	c := &Client{
		rv:      nil,
		cli:     nil,
		verbose: false,
		logger:  io.Stderr,
		debug:   false,
		airtel:  airtelMoney,
		tigo:    tc,
		mpesa:   vodaMpesa,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.rv = base.NewReceiver(c.logger, c.debug)

	app := cliApp(c)
	c.cli = app
	return c
}

func (c *Client)Run(args []string)error{
	return c.cli.Run(args)
}

func mnoAutoCheck(phone string) (mno.Operator, string, error) {
	op, fmtPhone, err := mno.Get(phone)
	if op == mno.Airtel {
		return op, fmtPhone[3:], err
	}
	return op, fmtPhone, err
}
