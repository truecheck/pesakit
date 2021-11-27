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
	"strings"
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
	Config struct {
		MpesaAuthEndpoint           string `conf:"flag:mpesa_auth_endpoint,env:PK_MPESA_AUTH_ENDPOINT"`
		MpesaPushEndpoint           string `conf:"flag:mpesa_push_endpoint,env:PK_MPESA_PUSH_ENDPOINT"`
		MpesaDisburseEndpoint       string `conf:"flag:mpesa_disburse_endpoint,env:PK_MPESA_DISBURSE_ENDPOINT"`
		MpesaCallbackEndpoint       string `conf:"flag:mpesa_callback_endpoint,env:PK_MPESA_CALLBACK_ENDPOINT"`
		MpesaAppName                string `conf:"flag:mpesa_app_name,env:PK_MPESA_APP_NAME"`
		MpesaAppVersion             string `conf:"flag:mpesa_app_version,env:PK_MPESA_APP_VERSION"`
		MpesaAppDescription         string `conf:"flag:mpesa_app_description,env:PK_MPESA_APP_DESCRIPTION"`
		MpesaBaseBasePath           string `conf:"flag:mpesa_base_base_path,env:PK_MPESA_BASE_PATH"`
		MpesaMarket                 string `conf:"flag:mpesa_market,env:PK_MPESA_MARKET"`
		MpesaPlatform               string `conf:"flag:mpesa_platform,env:PK_MPESA_PLATFORM"`
		MpesaAPIKey                 string `conf:"flag:mpesa_api_key,env:PK_MPESA_API_KEY"`
		MpesaPublicKey              string `conf:"flag:mpesa_public_key,env:PK_MPESA_PUBLIC_KEY"`
		MpesaSessionLifetimeMinutes int64  `conf:"flag:mpesa_session_lifetime_minutes,env:PK_MPESA_SESSION_LIFETIME"`
		MpesaServiceProvideCode     string `conf:"flag:mpesa_service_provide_code,env:PK_MPESA_SERVICE_PROVIDER_CODE"`
		MpesaTrustedSources         string `conf:"flag:mpesa_trusted_sources,env:PK_MPESA_TRUSTED_SOURCES"`
		MpesaCallbackPath           string `conf:"flag:mpesa_callback_path,env:PK_MPESA_CALLBACK_PATH"`
		AirtelAuthCountries         string `conf:"default:tanzania,flag:airtel_auth_countries,env:PK_AIRTEL_AUTH_COUNTRIES"`
		AirtelAccountCountries      string `conf:"default:tanzania,flag:airtel_account_countries,env:PK_AIRTEL_ACCOUNT_COUNTRIES"`
		AirtelCollectionCountries   string `conf:"default:tanzania,flag:airtel_collection_countries,env:PK_AIRTEL_COLLECTION_COUNTRIES"`
		AirtelDisburseCountries     string `conf:"default:tanzania,flag:airtel_disburse_countries,env:PK_AIRTEL_DISBURSE_COUNTRIES"`
		AirtelKYCCountries          string `conf:"default:tanzania,flag:airtel_kyc_countries,env:PK_AIRTEL_KYC_COUNTRIES"`
		AirtelTransactionCountries  string `conf:"default:tanzania,flag:airtel_transaction_countries,env:PK_AIRTEL_TRANSACTION_COUNTRIES"`
		AirtelDisbursePIN           string `conf:"flag:airtel_disburse_pin,env:PK_AIRTEL_DISBURSE_PIN"`
		AirtelCallbackPrivateKey    string `conf:"flag:airtel_callback_private_key,env:PK_AIRTEL_CALLBACK_PRIVATE_KEY"`
		AirtelCallbackAuth          bool   `conf:"default:false,flag:airtel_callback_auth,env:PK_AIRTEL_CALLBACK_AUTH"`
		AirtelPublicKey             string `conf:"flag:airtel_public_key,env:PK_AIRTEL_PUBLIC_KEY"`
		AirtelEnvironment           string `conf:"flag:airtel_environment,env:PK_AIRTEL_ENVIRONMENT"`
		AirtelClientID              string `conf:"flag:airtel_client_id,env:PK_AIRTEL_CLIENT_ID"`
		AirtelSecret                string `conf:"flag:airtel_secret,env:PK_AIRTEL_SECRET"`
		AirtelCallbackPath          string `conf:"flag:airtel_callback_path,env:PK_AIRTEL_CALLBACK_PATH"`
		TigoPushUsername            string `conf:"flag:tigo_push_username,env:PK_TIGO_PUSH_USERNAME"`
		TigoPushPassword            string `conf:"flag:tigo_push_password,env:PK_TIGO_PUSH_PASSWORD"`
		TigoPushPasswordGrantType   string `conf:"default:password,flag:tigo_push_password_grant_type,env:PK_TIGO_PUSH_PASSWORD_GRANT_TYPE"`
		TigoPushBaseURL             string `conf:"flag:tigo_push_base_url,env:PK_TIGO_PUSH_BASE_URL"`
		TigoPushTokenEndpoint       string `conf:"flag:tigo_push_token_endpoint,env:PK_TIGO_PUSH_TOKEN_ENDPOINT"`
		TigoPushBillerMSISDN        string `conf:"flag:tigo_push_biller_msisdn,env:PK_TIGO_PUSH_BILLER_MSISDN"`
		TigoPushBillerCode          string `conf:"flag:tigo_push_biller_code,env:PK_TIGO_PUSH_BILLER_CODE"`
		TigoPushPayEndpoint         string `conf:"flag:tigo_push_pay_endpoint,env:PK_TIGO_PUSH_ENDPOINT"`
		TigoDisburseAccountName     string `conf:"flag:tigo_disburse_account_name,env:PK_TIGO_DISBURSE_ACCOUNT_NAME"`
		TigoDisburseAccountMSISDN   string `conf:"flag:tigo_disburse_account_msisdn,env:PK_TIGO_DISBURSE_ACCOUNT_MSISDN"`
		TigoDisburseBrandID         string `conf:"flag:tigo_disburse_brand_id,env:PK_TIGO_DISBURSE_BRAND_ID"`
		TigoDisbursePIN             string `conf:"flag:tigo_disburse_pin,env:PK_TIGO_DISBURSE_PIN"`
		TigoDisburseRequestURL      string `conf:"flag:tigo_disburse_request_url,env:PK_TIGO_DISBURSE_URL"`
		TigoCallbackPath            string `conf:"flag:tigo_callback_path,env:PK_TIGO_CALLBACK_PATH"`
	}
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
		rv         base.Receiver
		cli        *clix.App
		verbose    bool
		logger     stdio.Writer
		debug      bool
		airtel     *airtel.Client
		tigo       *tigo.Client
		mpesa      *mpesa.Client
		configFile string
	}

	ClientOption func(*Client)
)

func ExtractMpesaConfig(c *Config) *mpesa.Config {
	endpoints := &mpesa.Endpoints{
		AuthEndpoint:     c.MpesaAuthEndpoint,
		PushEndpoint:     c.MpesaPushEndpoint,
		DisburseEndpoint: c.MpesaDisburseEndpoint,
	}

	ts := strings.Split(c.MpesaTrustedSources," ")
	return &mpesa.Config{
		Endpoints:              endpoints,
		Name:                   c.MpesaAppName,
		Version:                c.MpesaAppVersion,
		Description:            c.MpesaAppDescription,
		BasePath:               c.MpesaBaseBasePath,
		Market:                 mpesa.MarketFmt(c.MpesaMarket),
		Platform:               mpesa.PlatformFmt(c.MpesaPlatform),
		APIKey:                 c.MpesaAPIKey,
		PublicKey:              c.MpesaPublicKey,
		SessionLifetimeMinutes: c.MpesaSessionLifetimeMinutes,
		ServiceProvideCode:     c.MpesaServiceProvideCode,
		TrustedSources:         ts,
	}
}

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

func NewClient(airtelMoney *airtel.Client, tc *tigo.Client, vodaMpesa *mpesa.Client, opts ...ClientOption) *Client {

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

func (c *Client) Run(args []string) error {
	return c.cli.Run(args)
}

func mnoAutoCheck(phone string) (mno.Operator, string, error) {
	op, fmtPhone, err := mno.Get(phone)
	if op == mno.Airtel {
		return op, fmtPhone[3:], err
	}
	return op, fmtPhone, err
}
