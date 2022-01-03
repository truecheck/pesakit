package pesakit

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa"
	"io"
	"strings"
)

const (
	flagConfigMpesa  = "mpesa"
	flagConfigAirtel = "airtel"
	flagConfigTigo   = "tigo"
	defConfigValue   = false
)

func (app *App) configCommand() {
	rootCommand := app.root
	out := app.getWriter()
	var (
		airtelConfig bool
		mpesaConfig  bool
		tigoConfig   bool
	)

	configCommand := &cobra.Command{
		Use:   "config",
		Short: "manages clients configuration",
		Long: `manages clients configuration. This command is used to create, view,update and delete
clients configurations`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	configCommand.PersistentFlags().BoolVar(&airtelConfig, flagConfigAirtel, defConfigValue, "configure airtel client")
	configCommand.PersistentFlags().BoolVar(&mpesaConfig, flagConfigMpesa, defConfigValue, "configure mpesa client")
	configCommand.PersistentFlags().BoolVar(&tigoConfig, flagConfigTigo, defConfigValue, "configure tigo client")

	initConfigPrintCommand(configCommand, out)
	rootCommand.AddCommand(configCommand)
}

func initConfigPrintCommand(parentCommand *cobra.Command, out io.Writer) {
	configPrintCommand := &cobra.Command{
		Use:   "print",
		Short: "prints the configuration",
		Long:  `prints the configuration if not any is specified all configurations are printed`,
		Run: func(cmd *cobra.Command, args []string) {
			airtelConfigBool, err := cmd.Flags().GetBool(flagConfigAirtel)
			if err != nil {
				return
			}
			mpesaConfigBool, err := cmd.Flags().GetBool(flagConfigMpesa)
			if err != nil {
				return
			}
			tigoConfigBool, err := cmd.Flags().GetBool(flagConfigTigo)
			if err != nil {
				return
			}

			notAnySpecified := !(airtelConfigBool || mpesaConfigBool || tigoConfigBool)
			if notAnySpecified {
				mpesaConfig, err := loadMpesaConfig(cmd)
				if err != nil {
					return
				}
				_, _ = fmt.Fprintf(out, "Mpesa Config: %v\n", mpesaConfig)
				airtelConfig, err := loadAirtelConfig(cmd)
				if err != nil {
					return
				}
				_, _ = fmt.Fprintf(out, "Airtel Config: %v\n", airtelConfig)
				tigoConfig, err := loadTigoConfig(cmd)
				if err != nil {
					return
				}
				_, _ = fmt.Fprintf(out, "Tigo Config: %v\n", tigoConfig)
			}

		},
	}

	//configPrintCommand.SetHelpFunc(func(command *cobra.Command, strings []string) {
	//	set := parentCommand.Parent().PersistentFlags()
	//	markHiddenExcept(set, flagDebugMode, flagConfigAirtel, flagConfigMpesa, flagConfigTigo)
	//	command.Parent().HelpFunc()(command, strings)
	//})

	parentCommand.AddCommand(configPrintCommand)
}

// loadMpesaConfig loads all the mpesa configurations after a command
func loadMpesaConfig(command *cobra.Command) (*mpesa.Config, error) {
	authEndpoint, err := command.Flags().GetString(flagMpesaAuthEndpoint)
	if err != nil {
		return nil, err
	}
	pushEndpoint, err := command.Flags().GetString(flagMpesaPushEndpoint)
	if err != nil {
		return nil, err
	}
	disburseEndpoint, err := command.Flags().GetString(flagMpesaDisburseEndpoint)
	if err != nil {
		return nil, err
	}
	queryEndpoint, err := command.Flags().GetString(flagMpesaTransactionStatusEndpoint)
	if err != nil {
		return nil, err
	}
	directCreateEndpoint, err := command.Flags().GetString(flagMpesaDirectDebitCreateEndpoint)
	if err != nil {
		return nil, err
	}
	directPayEndpoint, err := command.Flags().GetString(flagMpesaDirectDebitPayEndpoint)
	if err != nil {
		return nil, err
	}
	reversalEndpoint, err := command.Flags().GetString(flagMpesaReversalEndpoint)
	if err != nil {
		return nil, err
	}
	b2bEndpoint, err := command.Flags().GetString(flagMpesaB2BEndpoint)
	if err != nil {
		return nil, err
	}
	trustedSourcesStr, err := command.Flags().GetString(flagMpesaTrustedSources)
	if err != nil {
		return nil, err
	}
	trustedSources := strings.Split(trustedSourcesStr, " ")
	providerCode, err := command.Flags().GetString(flagMpesaServiceProviderCode)
	if err != nil {
		return nil, err
	}
	sessionLifetime, err := command.Flags().GetInt64(flagMpesaSessionLifetimeMinutes)
	if err != nil {
		return nil, err
	}
	name, err := command.Flags().GetString(flagMpesaAppName)
	if err != nil {
		return nil, err
	}
	version, err := command.Flags().GetString(flagMpesaAppVersion)
	if err != nil {
		return nil, err
	}
	desc, err := command.Flags().GetString(flagMpesaAppDesc)
	if err != nil {
		return nil, err
	}
	basePath, err := command.Flags().GetString(flagMpesaAppDesc)
	if err != nil {
		return nil, err
	}
	marketString, err := command.Flags().GetString(flagMpesaMarket)
	if err != nil {
		return nil, err
	}
	market := mpesa.MarketFmt(marketString)
	platformStr, err := command.Flags().GetString(flagMpesaPlatform)
	if err != nil {
		return nil, err
	}
	platform := mpesa.PlatformFmt(platformStr)

	endpoints := &mpesa.Endpoints{
		AuthEndpoint:               authEndpoint,
		PushEndpoint:               pushEndpoint,
		DisburseEndpoint:           disburseEndpoint,
		QueryEndpoint:              queryEndpoint,
		DirectDebitCreateEndpoint:  directCreateEndpoint,
		DirectDebitPayEndpoint:     directPayEndpoint,
		TransactionReverseEndpoint: reversalEndpoint,
		B2BEndpoint:                b2bEndpoint,
	}
	config := &mpesa.Config{
		Endpoints:              endpoints,
		Name:                   name,
		Version:                version,
		Description:            desc,
		BasePath:               basePath,
		Market:                 market,
		Platform:               platform,
		APIKey:                 "",
		PublicKey:              "",
		SessionLifetimeMinutes: sessionLifetime,
		ServiceProvideCode:     providerCode,
		TrustedSources:         trustedSources,
	}

	if platform == mpesa.OPENAPI {
		// get production public key and api-key
		apiKey, err := command.Flags().GetString(flagMpesaOpenApiKey)
		if err != nil {
			return nil, err
		}
		config.APIKey = apiKey

		pubKey, err := command.Flags().GetString(flagMpesaOpenAPIPubKey)
		if err != nil {
			return nil, err
		}
		config.PublicKey = pubKey
	} else {
		// get sandbox public key and api-key
		apiKey, err := command.Flags().GetString(flagMpesaSandboxApiKey)
		if err != nil {
			return nil, err
		}
		config.APIKey = apiKey

		pubKey, err := command.Flags().GetString(flagMpesaSandboxPubKey)
		if err != nil {
			return nil, err
		}
		config.PublicKey = pubKey
	}

	return config, nil
}

func loadAirtelConfig(command *cobra.Command) (*airtel.Config, error) {
	countriesStr, err := command.Flags().GetString(flagAirtelCountries)
	if err != nil {
		return nil, err
	}
	countries := strings.Split(countriesStr, " ")
	disbursePin, err := command.Flags().GetString(flagAirtelDisbursePin)
	if err != nil {
		return nil, err
	}

	callbackPrivateKey, err := command.Flags().GetString(flagAirtelCallbackPrivKey)
	if err != nil {
		return nil, err
	}

	callbackAuth, err := command.Flags().GetBool(flagAirtelCallbackAuth)
	if err != nil {
		return nil, err
	}

	clientID, err := command.Flags().GetString(flagAirtelClientId)
	if err != nil {
		return nil, err
	}

	clientSecret, err := command.Flags().GetString(flagAirtelClientSecret)
	if err != nil {
		return nil, err
	}

	environmentStr, err := command.Flags().GetString(flagAirtelDeploymentEnv)
	if err != nil {
		return nil, err
	}

	environment := airtelEnv(environmentStr)

	publicKey, err := command.Flags().GetString(flagAirtelPublicKey)
	if err != nil {
		return nil, err
	}

	config := &airtel.Config{
		AllowedCountries: map[airtel.ApiGroup][]string{
			airtel.Transaction: countries,
			airtel.Collection:  countries,
			airtel.Disburse:    countries,
			airtel.Account:     countries,
			airtel.KYC:         countries,
		},
		DisbursePIN:        disbursePin,
		CallbackPrivateKey: callbackPrivateKey,
		CallbackAuth:       callbackAuth,
		PublicKey:          publicKey,
		Environment:        environment,
		ClientID:           clientID,
		Secret:             clientSecret,
	}

	return config, nil
}

func loadTigoConfig(command *cobra.Command) (*tigopesa.Config, error) {
	disburseAccountName, err := command.Flags().GetString(flagTigoDisburseAccountName)
	if err != nil {
		return nil, err
	}
	disburseAccountMSISDN, err := command.Flags().GetString(flagTigoDisburseAccountMSISDN)
	if err != nil {
		return nil, err
	}
	disburseBrandID, err := command.Flags().GetString(flagTigoDisburseBrandID)
	if err != nil {
		return nil, err
	}
	disbursePIN, err := command.Flags().GetString(flagTigoDisbursePIN)
	if err != nil {
		return nil, err
	}
	disburseRequestURL, err := command.Flags().GetString(flagTigoDisburseURL)
	if err != nil {
		return nil, err
	}

	pushUsername, err := command.Flags().GetString(flagTigoPushUsername)
	if err != nil {
		return nil, err
	}
	pushPassword, err := command.Flags().GetString(flagTigoPushPassword)
	if err != nil {
		return nil, err
	}
	pushPasswordGrantType, err := command.Flags().GetString(flagTigoPasswordGrantType)
	if err != nil {
		return nil, err
	}
	pushBillerCode, err := command.Flags().GetString(flagTigoPushBillerCode)
	if err != nil {
		return nil, err
	}
	pushBillerMSISDN, err := command.Flags().GetString(flagTigoPushBillerMSISDN)
	if err != nil {
		return nil, err
	}
	pushBaseURl, err := command.Flags().GetString(flagTigoPushBaseURL)
	if err != nil {
		return nil, err
	}
	pushTokenEndpoint, err := command.Flags().GetString(flagTigoPushTokenURL)
	if err != nil {
		return nil, err
	}

	pushPayEndpoint, err := command.Flags().GetString(flagTigoPushPayURL)
	if err != nil {
		return nil, err
	}

	config := &tigopesa.Config{
		Disburse: &tigopesa.DisburseConfig{
			AccountName:   disburseAccountName,
			AccountMSISDN: disburseAccountMSISDN,
			BrandID:       disburseBrandID,
			PIN:           disbursePIN,
			RequestURL:    disburseRequestURL,
		},
		Push: &tigopesa.PushConfig{
			Username:          pushUsername,
			Password:          pushPassword,
			PasswordGrantType: pushPasswordGrantType,
			BaseURL:           pushBaseURl,
			TokenEndpoint:     pushTokenEndpoint,
			BillerMSISDN:      pushBillerMSISDN,
			BillerCode:        pushBillerCode,
			PushPayEndpoint:   pushPayEndpoint,
		},
	}
	return config, nil
}

func airtelEnv(value string) airtel.Environment {
	if value == "production" || value == "prod" {
		return airtel.PRODUCTION
	} else if value == "sandbox" || value == "staging" {
		return airtel.STAGING
	} else {
		return airtel.STAGING
	}
}
