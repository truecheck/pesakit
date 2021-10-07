package cli

import (
	"fmt"
	"github.com/techcraftlabs/pesakit"
	"github.com/techcraftlabs/pesakit/airtel"
	"github.com/techcraftlabs/pesakit/mpesa"
	"github.com/techcraftlabs/pesakit/tigo"
	clix "github.com/urfave/cli/v2"
	"os"
	"text/tabwriter"
)

const (
	envAirtelPublicKey                   = "PK_AIRTEL_PUBKEY"
	envAirtelDisbursePin                 = "PK_AIRTEL_DISBURSE_PIN"
	envAirtelClientId                    = "PK_AIRTEL_CLIENT_ID"
	envAirtelClientSecret                = "PK_AIRTEL_CLIENT_SECRET"
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
	envMpesaPlatform                     = "PK_MPESA_PLATFORM"
	envMpesaMarket                       = "PK_MPESA_MARKET"
	envMpesaAuthEndpoint                 = "PK_MPESA_AUTH_ENDPOINT"
	envMpesaPushEndpoint                 = "PK_MPESA_PUSH_ENDPOINT"
	envMpesaDisburseEndpoint             = "PK_MPESA_DISBURSE_ENDPOINT"
	envMpesaBaseURL                      = "PK_MPESA_BASE_URL"
	envMpesaAppName                      = "PK_MPESA_APP_NAME"
	envMpesaAppVersion                   = "PK_MPESA_APP_VERSION"
	envMpesaAppDesc                      = "PK_MPESA_APP_DESCRIPTION"
	envMpesaSandboxApiKey                = "PK_MPESA_API_KEY"
	envMpesaSandboxPubKey                = "PK_MPESA_PUBLIC_KEY"
	envMpesaSessionLifetimeMinutes       = "PK_MPESA_SESSION_LIFETIME_MINUTES"
	envMpesaServiceProvideCode           = "PK_MPESA_SERVICE_PROVIDER_CODE"
	envMpesaTrustedSources               = "PK_MPESA_TRUSTED_SOURCES"
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
)

func ConfigCommand(client *pesakit.Client) *Cmd {
	flags := []clix.Flag{
		&clix.BoolFlag{
			Name:    "tigo",
			Aliases: []string{"t"},
			Usage:   "set to true if you want to print tigo config",
		},
		&clix.BoolFlag{
			Name:    "mpesa",
			Aliases: []string{"v", "vodacom","voda"},
			Usage:   "set to true if you want to print vodacom mpesa config",
		},
		&clix.BoolFlag{
			Name:    "airtel",
			Aliases: []string{"a"},
			Usage:   "set to true if you want to print airtel money config",
		},
	}
	return &Cmd{
		ApiClient:   client,
		RequestType: Config,
		Name:        "config",
		Usage:       "show all the running configurations",
		Description: "prints all the mno configurations",
		Flags:       flags,
		SubCommands: nil,
	}
}

func printMpesaConf(w *tabwriter.Writer, v *mpesa.Config){

	_, _ = fmt.Fprintf(w, "\n %s\t", "-------------------------")
	_, _ = fmt.Fprintf(w, "\n %s\t", "MPESA CONFIGURATIONS")
	_, _ = fmt.Fprintf(w, "\n %s\t", "-------------------------")
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaAppName, v.Name)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaAppVersion, v.Version)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaSandboxPubKey, v.PublicKey)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaSandboxApiKey, v.APIKey)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaPlatform, v.Platform.String())
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaMarket, v.Market.Description())
	_, _ = fmt.Fprintf(w, "\n %s: \t%d\t", envMpesaSessionLifetimeMinutes, v.SessionLifetimeMinutes)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaTrustedSources, v.TrustedSources)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaServiceProvideCode, v.ServiceProvideCode)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaBaseURL, v.BasePath)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaAppDesc, v.Description)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaAuthEndpoint, v.Endpoints.AuthEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaDisburseEndpoint, v.Endpoints.DisburseEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envMpesaPushEndpoint, v.Endpoints.PushEndpoint)
	_, _ = fmt.Fprintf(w, "\n")
}

func printAirtelConf(w *tabwriter.Writer, a *airtel.Config)  {
	_, _ = fmt.Fprintf(w, "\n %s\t", "-------------------------")
	_, _ = fmt.Fprintf(w, "\n %s\t", "AIRTEL CONFIGURATIONS")
	_, _ = fmt.Fprintf(w, "\n %s\t", "-------------------------")
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelPublicKey, a.PublicKey)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelDisbursePin, a.DisbursePIN)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelCallbackPrivKey, a.CallbackPrivateKey)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelClientId, a.ClientID)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelClientSecret, a.Secret)
	_, _ = fmt.Fprintf(w, "\n %s: \t%v\t", envAirtelCountries, a.AllowedCountries)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelAuthEndpoint, a.Endpoints.AuthEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelPushEndpoint, a.Endpoints.PushEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelRefundEndpoint, a.Endpoints.RefundEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelPushEnquiryEndpoint, a.Endpoints.PushEnquiryEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelDisbursementEndpoint, a.Endpoints.DisbursementEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelTransactionSummaryEndpoint, a.Endpoints.TransactionSummaryEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelBalanceEnquiryEndpoint, a.Endpoints.BalanceEnquiryEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelUserEnquiryEndpoint, a.Endpoints.UserEnquiryEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%t\t", envAirtelCallbackAuth, a.CallbackAuth)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envAirtelDisbursementEnquiryEndpoint, a.Endpoints.DisbursementEnquiryEndpoint)
}

func printTigoConf(w *tabwriter.Writer, tc *tigo.Config){
	p := tc.PushConfig
	d := tc.DisburseConfig


	_, _ = fmt.Fprintf(w, "\n %s\t", "-------------------------")
	_, _ = fmt.Fprintf(w, "\n %s\t", "TIGO CONFIGURATIONS")
	_, _ = fmt.Fprintf(w, "\n %s\t", "-------------------------")
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushUsername, p.Username)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushPassword, p.Password)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushBaseURL, tc.BaseURL)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushTokenURL, p.TokenEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushPayURL, p.PushPayEndpoint)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushBillerCode, p.BillerCode)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPushBillerMSISDN, p.BillerMSISDN)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoDisburseURL, d.RequestURL)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoDisbursePIN, d.PIN)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoDisburseBrandID, d.BrandID)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoDisburseAccountMSISDN, d.AccountMSISDN)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoDisburseAccountName, d.AccountName)
	_, _ = fmt.Fprintf(w, "\n %s: \t%s\t", envTigoPasswordGrantType, p.PasswordGrantType)

}



func printConfigs(w *tabwriter.Writer,a *airtel.Config, v *mpesa.Config, t *tigo.Config) {
	printMpesaConf(w,v)
	printAirtelConf(w,a)
	printTigoConf(w,t)
}

func (c *Cmd) configAction(ctx *clix.Context) error {
	// initialize tabwriter
	w := new(tabwriter.Writer)

	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer func(w *tabwriter.Writer) {
		err := w.Flush()
		if err != nil {
			fmt.Printf("error while closing tabwriter: %v\n", err)
		}
	}(w)

	mc := c.ApiClient.Mpesa.Conf
	tc := c.ApiClient.TigoPesa.Config
	ac := c.ApiClient.AirtelMoney.Conf

	isPrintTigo := ctx.Bool("tigo")
	isPrintVoda := ctx.Bool("mpesa")
	isPrintAirtel := ctx.Bool("airtel")
	isPrintAll := (!isPrintVoda && !isPrintAirtel && !isPrintTigo) || (isPrintVoda && isPrintAirtel && isPrintTigo)
	if isPrintAll {
		printConfigs(w,ac,mc,tc)
		return nil
	}

	if isPrintVoda {
		printMpesaConf(w,mc)
	}

	if isPrintAirtel {
		printAirtelConf(w,ac)
	}

	if isPrintTigo {
		printTigoConf(w,tc)
	}

	return nil
}
