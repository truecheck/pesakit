package main

import (
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pesakit/pesakit"
	"github.com/pesakit/pesakit/cli"
	"github.com/pesakit/pesakit/pkg/env"
	"github.com/pesakit/pesakit/tigo"
	tigopesa "github.com/techcraftlabs/tigopesa"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa/disburse"
	"github.com/techcraftlabs/tigopesa/push"
	"github.com/techcraftlabs/tigopesa/ussd"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	envAirtelPublicKey                   = "PK_AIRTEL_PUBKEY"
	envAirtelDisbursePin                 = "PK_AIRTEL_DISBURSE_PIN"
	envAirtelClientId                    = "PK_AIRTEL_CLIENT_ID"
	envAirtelClientSecret                = "PK_AIRTEL_CLIENT_SECRET"
	envDebugMode                         = "PK_DEBUG_MODE"
	envAirtelDeploymentEnv               = "PK_AIRTEL_DEPLOYMENT"
	envAirtelCallbackAuth                = "PK_AIRTEL_CALLBACK_AUTH"
	envAirtelCallbackPrivKey             = "PK_AIRTEL_CALLBACK_PRIVKEY"
	envAirtelCountries                   = "PK_AIRTEL_COUNTRIES"
	envAirtelAuthEndpoint                = "PK_AIRTEL_AUTH_ENDPOINT"
	envAirtelPushEndpoint                = "PK_AIRTEL_PUSH_ENDPOINT"
	envAirtelRefundEndpoint              = "PK_AIRTEL_REFUND_ENDPOINT"
	envAirtelPushEnquiryEndpoint         = "PK_AIRTEL_PUSH_ENQUIRY_ENDPOINT"
	envAirtelDisbursementEndpoint        = "PK_AIRTEL_DISBURSE_ENDPOINT"
	envAirtelDisbursementEnquiryEndpoint = "PK_AIRTEL_DISBURSE_ENQUIRY_ENDPOINT"
	envAirtelTransactionSummaryEndpoint  = "PK_AIRTEL_SUMMARY_ENDPOINT"
	envAirtelBalanceEnquiryEndpoint      = "PK_AIRTEL_BALANCE_ENDPOINT"
	envAirtelUserEnquiryEndpoint         = "PK_AIRTEL_USER_ENDPOINT"
	defAirtelPublicKey                   = ""
	defAirtelDisbursePin                 = ""
	defAirtelClientId                    = ""
	defAirtelClientSecret                = ""
	defDebugMode                         = true
	defAirtelDeploymentEnv               = "staging"
	defAirtelCallbackAuth                = false
	defAirtelCallbackPrivKey             = "zITVAAGYSlzl1WkUQJn81kbpT5drH3koffT8jCkcJJA="
	defAirtelCountries                   = "tanzania"
	defAirtelAuthEndpoint                = "/auth/oauth2/token"
	defAirtelPushEndpoint                = "/merchant/v1/payments/"
	defAirtelRefundEndpoint              = "/standard/v1/payments/refund"
	defAirtelPushEnquiryEndpoint         = "/standard/v1/payments/"
	defAirtelDisbursementEndpoint        = "/standard/v1/disbursements/"
	defAirtelDisbursementEnquiryEndpoint = "/standard/v1/disbursements/"
	defAirtelTransactionSummaryEndpoint  = "/merchant/v1/transactions"
	defAirtelBalanceEnquiryEndpoint      = "/standard/v1/users/balance"
	defAirtelUserEnquiryEndpoint         = "/standard/v1/users/"
	envMpesaPlatform                     = "PK_MPESA_PLATFORM"
	envMpesaMarket                       = "PK_MPESA_MARKET"
	envMpesaAuthEndpoint                 = "PK_MPESA_AUTH_ENDPOINT"
	envMpesaPushEndpoint                 = "PK_MPESA_PUSH_ENDPOINT"
	envMpesaDisburseEndpoint             = "PK_MPESA_DISBURSE_ENDPOINT"
	envMpesaBaseURL                      = "PK_MPESA_BASE_URL"
	envMpesaAppName                      = "PK_MPESA_APP_NAME"
	envMpesaAppVersion                   = "PK_MPESA_APP_VERSION"
	envMpesaAppDesc                      = "PK_MPESA_APP_DESCRIPTION"
	envMpesaSandboxApiKey                = "PK_MPESA_SANDBOX_API_KEY"
	envMpesaOpenApiKey                   = "PK_MPESA_OPENAPI_KEY"
	envMpesaSandboxPubKey                = "PK_MPESA_SANDBOX_PUBLIC_KEY"
	envMpesaOpenApiPubKey                = "PK_MPESA_OPENAPI_PUBLIC_KEY"
	envMpesaSessionLifetimeMinutes       = "PK_MPESA_SESSION_LIFETIME_MINUTES"
	envMpesaServiceProvideCode           = "PK_MPESA_SERVICE_PROVIDER_CODE"
	envMpesaTrustedSources               = "PK_MPESA_TRUSTED_SOURCES"
	defMpesaPlatform                     = "sandbox"
	defMpesaMarket                       = "Tanzania"
	defMpesaAuthEndpoint                 = "/getSession/"
	defMpesaPushEndpoint                 = ""
	defMpesaDisburseEndpoint             = ""
	defMpesaBaseURL                      = "openapi.m-pesa.com"
	defMpesaAppName                      = "beanpay"
	defMpesaAppVersion                   = "1.0"
	defMpesaAppDesc                      = "unified payment as a service"
	defMpesaSandboxApiKey                = ""
	defMpesaOpenApiKey                   = ""
	defMpesaSandboxPubKey                = ""
	defMpesaOpenApiPubKey                = ""
	defMpesaSessionLifetimeMinutes       = 60
	defMpesaServiceProvideCode           = ""
	defMpesaTrustedSources               = "openapi.m-pesa.com"
	envTigoDisbursePIN                   = "PK_TIGO_DISBURSE_PIN"
	envTigoDisburseURL                   = "PK_TIGO_DISaBURSE_URL"
	envTigoDisburseBrandID               = "PK_TIGO_DISBURSE_BRAND_ID"
	envTigoDisburseAccountMSISDN         = "PK_TIGO_DISBURSE_ACCOUNT_MSISDN"
	envTigoDisburseAccountName           = "PK_TIGO_DISBURSE_ACCOUNT_NAME"
	envTigoPushUsername                  = "PK_TIGO_PUSH_USERNAME"
	envTigoPushPassword                  = "PK_TIGO_PUSH_PASSWORD"
	envTigoPushBillerMSISDN              = "PK_TIGO_PUSH_BILLER_MSISDN"
	envTigoPushBaseURL                   = "PK_TIGO_PUSH_BASE_URL"
	envTigoPushBillerCode                = "PK_TIGO_PUSH_BILLER_CODE"
	envTigoPushTokenURL                  = "PK_TIGO_PUSH_TOKEN_URL"
	envTigoPushPayURL                    = "PK_TIGO_PUSH_PAY_URL"
	envTigoPasswordGrantType             = "PK_TIGO_PASSWORD_GRANT_TYPE"
	defTigoDisburseAccountName           = ""
	defTigoDisburseAccountMSISDN         = ""
	defTigoDisburseBrandID               = ""
	defTigoDisbursePIN                   = ""
	defTigoDisburseURL                   = ""
	defTigoPushUsername                  = ""
	defTigoPushPassword                  = ""
	defTigoPushBaseURL                   = ""
	defTigoPushBillerMSISDN              = ""
	defTigoPushBillerCode                = ""
	defTigoPushTokenURL                  = ""
	defTigoPushPayURL                    = ""
	defTigoPasswordGrantType             = "password"
)

func loadTigoEnv() *tigopesa.Config {

	var (
		disburseAccountName   = env.String(envTigoDisburseAccountName, defTigoDisburseAccountName)
		disburseAccountMSISDN = env.String(envTigoDisburseAccountMSISDN, defTigoDisburseAccountMSISDN)
		disburseBrandID       = env.String(envTigoDisburseBrandID, defTigoDisburseBrandID)
		disbursePIN           = env.String(envTigoDisbursePIN, defTigoDisbursePIN)
		disburseURL           = env.String(envTigoDisburseURL, defTigoDisburseURL)
		pushUsername          = env.String(envTigoPushUsername, defTigoPushUsername)
		pushPassword          = env.String(envTigoPushPassword, defTigoPushPassword)
		pushBaseURL           = env.String(envTigoPushBaseURL, defTigoPushBaseURL)
		pushBillerMSISDN      = env.String(envTigoPushBillerMSISDN, defTigoPushBillerMSISDN)
		pushBillerCode        = env.String(envTigoPushBillerCode, defTigoPushBillerCode)
		pushTokenURL          = env.String(envTigoPushTokenURL, defTigoPushTokenURL)
		pushPayURL            = env.String(envTigoPushPayURL, defTigoPushPayURL)
		pwdGrantType          = env.String(envTigoPasswordGrantType, defTigoPasswordGrantType)
	)

	//conf := &tigopesa.Config{
		
		
		//DisburseConfig: &tigopesa.DisburseConfig{
		//	AccountName:   disburseAccountName,
		//	AccountMSISDN: disburseAccountMSISDN,
		//	BrandID:       disburseBrandID,
		//	PIN:           disbursePIN,
		//	RequestURL:    disburseURL,
		//},
		//PushConfig: &tigopesa.PushConfig{
		//	BaseURL:           pushBaseURL,
		//	Username:          pushUsername,
		//	Password:          pushPassword,
		//	PasswordGrantType: pwdGrantType,
		//	TokenEndpoint:     pushTokenURL,
		//	PushPayEndpoint:   pushPayURL,
		//	BillerMSISDN:      pushBillerMSISDN,
		//	BillerCode:        pushBillerCode,
		//},
//	}
	
	conf := &tigopesa.Config{
		Disburse: &disburse.Config{
				AccountName:   disburseAccountName,
				AccountMSISDN: disburseAccountMSISDN,
				BrandID:       disburseBrandID,
				PIN:           disbursePIN,
				RequestURL:    disburseURL,
		},
		Push:     &push.Config{
				BaseURL:           pushBaseURL,
				Username:          pushUsername,
				Password:          pushPassword,
				PasswordGrantType: pwdGrantType,
				TokenEndpoint:     pushTokenURL,
				PushPayEndpoint:   pushPayURL,
				BillerMSISDN:      pushBillerMSISDN,
				BillerCode:        pushBillerCode,
		},
		Ussd:     &ussd.Config{
			AccountName:   "",
			AccountMSISDN: "",
			BillerNumber:  "",
			RequestURL:    "",
			NamecheckURL:  "",
		},
	}

	return conf
}

func loadMpesaEnv() *mpesa.Config {

	var (
		name                = env.String(envMpesaAppName, defMpesaAppName)
		version             = env.String(envMpesaAppVersion, defMpesaAppVersion)
		desc                = env.String(envMpesaAppDesc, defMpesaAppDesc)
		basePath            = env.String(envMpesaBaseURL, defMpesaBaseURL)
		sandboxApiKey       = env.String(envMpesaSandboxApiKey, defMpesaSandboxApiKey)
		sandboxPubKey       = env.String(envMpesaSandboxPubKey, defMpesaSandboxPubKey)
		openApiKey          = env.String(envMpesaOpenApiKey, defMpesaOpenApiKey)
		openApiPubKey       = env.String(envMpesaOpenApiPubKey, defMpesaOpenApiPubKey)
		spc                 = env.String(envMpesaServiceProvideCode, defMpesaServiceProvideCode)
		trustArray          = strings.Split(env.String(envMpesaTrustedSources, defMpesaTrustedSources), " ")
		sessLifeTimeMinutes = env.Int64(envMpesaSessionLifetimeMinutes, defMpesaSessionLifetimeMinutes)
		authEndpoint        = env.String(envMpesaAuthEndpoint, defMpesaAuthEndpoint)
		pushEndpoint        = env.String(envMpesaPushEndpoint, defMpesaPushEndpoint)
		disburseEndpoint    = env.String(envMpesaDisburseEndpoint, defMpesaDisburseEndpoint)
		platformEnv         = env.String(envMpesaPlatform, defMpesaPlatform)
		marketEnv           = env.String(envMpesaMarket, defMpesaMarket)
	)

	market := mpesa.MarketFmt(marketEnv)
	platform := mpesa.PlatformFmt(platformEnv)

	var apiKey, pubKey string

	if platform == mpesa.OPENAPI {
		apiKey, pubKey = openApiKey, openApiPubKey
	} else {
		apiKey, pubKey = sandboxApiKey, sandboxPubKey
	}

	conf := &mpesa.Config{
		Market:   market,
		Platform: platform,
		Endpoints: &mpesa.Endpoints{
			AuthEndpoint:     authEndpoint,
			PushEndpoint:     pushEndpoint,
			DisburseEndpoint: disburseEndpoint,
		},
		Name:                   name,
		Version:                version,
		Description:            desc,
		BasePath:               basePath,
		APIKey:                 apiKey,
		PublicKey:              pubKey,
		SessionLifetimeMinutes: sessLifeTimeMinutes,
		ServiceProvideCode:     spc,
		TrustedSources:         trustArray,
	}

	return conf
}

func loadAirtelEnv() *airtel.Config {
	var (
		//baseURL         = env.String(envAirtelBaseURL, defAirtelBaseURL)
		pubKey          = env.String(envAirtelPublicKey, defAirtelPublicKey)
		disbursePin     = env.String(envAirtelDisbursePin, defAirtelDisbursePin)
		callbackPrivKey = env.String(envAirtelCallbackPrivKey, defAirtelCallbackPrivKey)
		clientID        = env.String(envAirtelClientId, defAirtelClientId)
		secret          = env.String(envAirtelClientSecret, defAirtelClientSecret)
		//	debugMode       = env.Bool(envDebugMode, defAirtelDebugMode)
		callbackAuth = env.Bool(envAirtelCallbackAuth, defAirtelCallbackAuth)
		countries    = strings.Split(env.String(envAirtelCountries, defAirtelCountries), " ")
		endpoints    = &airtel.Endpoints{
			AuthEndpoint:                env.String(envAirtelAuthEndpoint, defAirtelAuthEndpoint),
			PushEndpoint:                env.String(envAirtelPushEndpoint, defAirtelPushEndpoint),
			RefundEndpoint:              env.String(envAirtelRefundEndpoint, defAirtelRefundEndpoint),
			PushEnquiryEndpoint:         env.String(envAirtelPushEnquiryEndpoint, defAirtelPushEnquiryEndpoint),
			DisbursementEndpoint:        env.String(envAirtelDisbursementEndpoint, defAirtelDisbursementEndpoint),
			DisbursementEnquiryEndpoint: env.String(envAirtelDisbursementEnquiryEndpoint, defAirtelDisbursementEnquiryEndpoint),
			TransactionSummaryEndpoint:  env.String(envAirtelTransactionSummaryEndpoint, defAirtelTransactionSummaryEndpoint),
			BalanceEnquiryEndpoint:      env.String(envAirtelBalanceEnquiryEndpoint, defAirtelBalanceEnquiryEndpoint),
			UserEnquiryEndpoint:         env.String(envAirtelUserEnquiryEndpoint, defAirtelUserEnquiryEndpoint),
		}
	)

	config := &airtel.Config{
		Endpoints: endpoints,
		AllowedCountries: map[string][]string{
			airtel.TransactionApiGroup:  countries,
			airtel.CollectionApiGroup:   countries,
			airtel.DisbursementApiGroup: countries,
			airtel.AccountApiGroup:      countries,
			airtel.KycApiGroup:          countries,
		},
		DisbursePIN:        disbursePin,
		CallbackPrivateKey: callbackPrivKey,
		CallbackAuth:       callbackAuth,
		PublicKey:          pubKey,
		ClientID:           clientID,
		Secret:             secret,
	}

	return config
}

func init() {
	_ = godotenv.Load(".env")
}

func main(){
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("could not get current working directory %v\n", err)
		os.Exit(1)
		return
	}

	// automatically load in current path the files named .env or pesakit.env
	f1 := fmt.Sprintf(filepath.Join(wd, ".env"))
	f2 := fmt.Sprintf(filepath.Join(wd, "pesakit.env"))
	_ = godotenv.Load(f1, f2)
	airtelDeployEnv := strings.ToLower(env.String(envAirtelDeploymentEnv, defAirtelDeploymentEnv))

	debugMode := env.Bool(envDebugMode, defDebugMode)
	aConfig := loadAirtelEnv()
	vConfig := loadMpesaEnv()
	tConfig := loadTigoEnv()

	var airtelOptions []airtel.ClientOption
	airtelDebugOption := airtel.WithDebugMode(debugMode)
	airtelCbOption := airtel.WithCallbackHandler(cli.AirtelCallbackHandler())
	deployEnvOption := airtel.WithEnvironment(airtelDeployEnv)
	airtelOptions = append(airtelOptions, airtelDebugOption, deployEnvOption,airtelCbOption)
	a := airtel.NewClient(aConfig, airtelOptions...)

	// create mpesa client
	var mpesaOptions []mpesa.ClientOption

	debugOption := mpesa.WithDebugMode(debugMode)
	cbOption := mpesa.WithCallbackHandler(cli.MpesaCallbackHandler())
	mpesaOptions = append(mpesaOptions, debugOption,cbOption)
	m := mpesa.NewClient(vConfig, mpesaOptions...)

	t := tigopesa.NewClient(tConfig, tigopesa.WithDebugMode(debugMode))
	t.SetCallbackHandler(cli.TigoCallbackHandler())
	client := pesakit.NewClient(a, t, m)

	app := cli.New(client)

	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}