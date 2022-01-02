package pesakit

import (
	"fmt"
	"github.com/pesakit/pesakit/home"
	"path/filepath"

	"github.com/pesakit/pesakit/env"
	"github.com/spf13/cobra"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/mpesa"
	"github.com/techcraftlabs/tigopesa"
)

func (app *App) createRootCommand() {
	var (
		varDebugMode                      = env.Bool(envDebugMode, defDebugMode)
		varAirtelPublicKey                = env.String(envAirtelPublicKey, defAirtelPublicKey)
		varAirtelDisbursePin              = env.String(envAirtelDisbursePin, defAirtelDisbursePin)
		varAirtelCallbackPrivKey          = env.String(envAirtelCallbackPrivKey, defAirtelCallbackPrivKey)
		varAirtelClientID                 = env.String(envAirtelClientId, defAirtelClientId)
		varAirtelClientSecret             = env.String(envAirtelClientSecret, defAirtelClientSecret)
		varAirtelDeploymentEnv            = env.String(envAirtelDeploymentEnv, defAirtelDeploymentEnv)
		varAirtelCallbackAuth             = env.Bool(envAirtelCallbackAuth, defAirtelCallbackAuth)
		varAirtelCountries                = env.String(envAirtelCountries, defAirtelCountries)
		varMpesaPlatform                  = env.String(envMpesaPlatform, defMpesaPlatform)
		varMpesaMarket                    = env.String(envMpesaMarket, defMpesaMarket)
		varMpesaAuthEndpoint              = env.String(envMpesaAuthEndpoint, defMpesaAuthEndpoint)
		varMpesaPushEndpoint              = env.String(envMpesaPushEndpoint, defMpesaPushEndpoint)
		varMpesaDisburseEndpoint          = env.String(envMpesaDisburseEndpoint, defMpesaDisburseEndpoint)
		varMpesaReversalEndpoint          = env.String(envMpesaReversalEndpoint, defMpesaReversalEndpoint)
		varMpesaB2BEndpoint               = env.String(envMpesaB2BEndpoint, defMpesaB2BEndpoint)
		varMpesaDirectDebitCreateEndpoint = env.String(envMpesaDirectDebitCreateEndpoint, defMpesaDirectDebitCreateEndpoint)
		varMpesaDirectDebitPayEndpoint    = env.String(envMpesaDirectDebitPayEndpoint, defMpesaDirectDebitPayEndpoint)
		varMpesaTransactionStatusEndpoint = env.String(envMpesaTransactionStatusEndpoint, defMpesaTransactionStatusEndpoint)
		varMpesaBaseURL                   = env.String(envMpesaBaseURL, defMpesaBaseURL)
		varMpesaAppName                   = env.String(envMpesaAppName, defMpesaAppName)
		varMpesaAppVersion                = env.String(envMpesaAppVersion, defMpesaAppVersion)
		varMpesaAppDescription            = env.String(envMpesaAppDesc, defMpesaAppDesc)
		varMpesaSandboxAPIKey             = env.String(envMpesaSandboxAPIKey, defMpesaSandboxAPIKey)
		varMpesaOpenAPIKey                = env.String(envMpesaOpenAPIKey, defMpesaOpenApiKey)
		varMpesaSandboxPublicKey          = env.String(envMpesaSandboxPubKey, defMpesaSandboxPubKey)
		varMpesaOpenAPIPublicKey          = env.String(envMpesaOpenApiPubKey, defMpesaOpenApiPubKey)
		varMpesaSessionLifetimeMinutes    = env.Int64(envMpesaSessionLifetimeMinutes, defMpesaSessionLifetimeMinutes)
		varMpesaServiceProviderCode       = env.String(envMpesaServiceProviderCode, defMpesaServiceProviderCode)
		varMpesaTrustedSources            = env.String(envMpesaTrustedSources, defMpesaTrustedSources)
		varTigoDisbursePIN                = env.String(envTigoDisbursePIN, defTigoDisbursePIN)
		varTigoDisburseURL                = env.String(envTigoDisburseURL, defTigoDisburseURL)
		varTigoDisburseBrandID            = env.String(envTigoDisburseBrandID, defTigoDisburseBrandID)
		varTigoDisburseAccountMSISDN      = env.String(envTigoDisburseAccountMSISDN, defTigoDisburseAccountMSISDN)
		varTigoDisburseAccountName        = env.String(envTigoDisburseAccountName, defTigoDisburseAccountName)
		varTigoPushUsername               = env.String(envTigoPushUsername, defTigoPushUsername)
		varTigoPushPassword               = env.String(envTigoPushPassword, defTigoPushPassword)
		varTigoPushBillerMSISDN           = env.String(envTigoPushBillerMSISDN, defTigoPushBillerMSISDN)
		varTigoPushBaseURL                = env.String(envTigoPushBaseURL, defTigoPushBaseURL)
		varTigoPushBillerCode             = env.String(envTigoPushBillerCode, defTigoPushBillerCode)
		varTigoPushTokenURL               = env.String(envTigoPushTokenURL, defTigoPushTokenURL)
		varTigoPushPayURL                 = env.String(envTigoPushPayURL, defTigoPushPayURL)
		varTigoPasswordGrantType          = env.String(envTigoPasswordGrantType, defTigoPasswordGrantType)
		varConfigFile                     = env.String(envConfigFile, defConfigFile)
		varHomeDirectory                  = env.String(envHomeDirectory, defHomeDirectory)
	)

	var rootCommand = &cobra.Command{
		Use:              appName,
		Short:            appShortDesc,
		Long:             appLongDescription,
		PersistentPreRun: app.persistentPreRun,
	}

	rootCommand.PersistentFlags().StringVar(&varHomeDirectory, flagHomeDirectory, defHomeDirectory, usageHomeDirectory)
	rootCommand.PersistentFlags().BoolVar(&varDebugMode, flagDebugMode, varDebugMode, usageDebugMode)
	rootCommand.PersistentFlags().StringVar(&varAirtelPublicKey, flagAirtelPublicKey, varAirtelPublicKey, usageAirtelPublicKey)
	rootCommand.PersistentFlags().StringVar(&varAirtelDisbursePin, flagAirtelDisbursePin,
		varAirtelDisbursePin, usageAirtelDisbursePin)
	rootCommand.PersistentFlags().StringVar(&varAirtelCallbackPrivKey, flagAirtelCallbackPrivKey, varAirtelCallbackPrivKey,
		usageAirtelCallbackPrivKey)
	rootCommand.PersistentFlags().StringVar(&varAirtelClientID, flagAirtelClientId, varAirtelClientID, usageAirtelClientId)
	rootCommand.PersistentFlags().StringVar(&varAirtelClientSecret, flagAirtelClientSecret,
		varAirtelClientSecret, usageAirtelClientSecret)
	rootCommand.PersistentFlags().StringVar(&varAirtelDeploymentEnv, flagAirtelDeploymentEnv,
		varAirtelDeploymentEnv, usageAirtelDeploymentEnv)
	rootCommand.PersistentFlags().BoolVar(&varAirtelCallbackAuth, flagAirtelCallbackAuth,
		varAirtelCallbackAuth, usageAirtelCallbackAuth)
	rootCommand.PersistentFlags().StringVar(&varAirtelCountries, flagAirtelCountries, varAirtelCountries, usageAirtelCountries)
	rootCommand.PersistentFlags().StringVar(&varMpesaPlatform, flagMpesaPlatform, varMpesaPlatform, usageMpesaPlatform)
	rootCommand.PersistentFlags().StringVar(&varMpesaMarket, flagMpesaMarket, varMpesaMarket, usageMpesaMarket)
	rootCommand.PersistentFlags().StringVar(&varMpesaAuthEndpoint, flagMpesaAuthEndpoint,
		varMpesaAuthEndpoint, usageMpesaAuthEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaPushEndpoint, flagMpesaPushEndpoint,
		varMpesaPushEndpoint, usageMpesaPushEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaDisburseEndpoint, flagMpesaDisburseEndpoint,
		varMpesaDisburseEndpoint, usageMpesaDisburseEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaReversalEndpoint, flagMpesaReversalEndpoint,
		varMpesaReversalEndpoint, usageMpesaReversalEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaB2BEndpoint, flagMpesaB2BEndpoint,
		varMpesaB2BEndpoint, usageMpesaB2BEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaDirectDebitCreateEndpoint, flagMpesaDirectDebitCreateEndpoint,
		varMpesaDirectDebitCreateEndpoint, usageMpesaDirectDebitCreateEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaDirectDebitPayEndpoint, flagMpesaDirectDebitPayEndpoint,
		varMpesaDirectDebitPayEndpoint, usageMpesaDirectDebitPayEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaTransactionStatusEndpoint, flagMpesaTransactionStatusEndpoint,
		varMpesaTransactionStatusEndpoint, usageMpesaTransactionStatusEndpoint)
	rootCommand.PersistentFlags().StringVar(&varMpesaBaseURL, flagMpesaBaseURL, varMpesaBaseURL, usageMpesaBaseURL)
	rootCommand.PersistentFlags().StringVar(&varMpesaAppName, flagMpesaAppName, varMpesaAppName, usageMpesaAppName)
	rootCommand.PersistentFlags().StringVar(&varMpesaAppVersion, flagMpesaAppVersion, varMpesaAppVersion, usageMpesaAppVersion)
	rootCommand.PersistentFlags().StringVar(&varMpesaAppDescription, flagMpesaAppDesc,
		varMpesaAppDescription, usageMpesaAppDesc)
	rootCommand.PersistentFlags().StringVar(&varMpesaSandboxAPIKey, flagMpesaSandboxApiKey,
		varMpesaSandboxAPIKey, usageMpesaSandboxApiKey)
	rootCommand.PersistentFlags().StringVar(&varMpesaOpenAPIKey, flagMpesaOpenApiKey, varMpesaOpenAPIKey, usageMpesaOpenApiKey)
	rootCommand.PersistentFlags().StringVar(&varMpesaSandboxPublicKey, flagMpesaSandboxPubKey,
		varMpesaSandboxPublicKey, usageMpesaSandboxPubKey)
	rootCommand.PersistentFlags().StringVar(&varMpesaOpenAPIPublicKey, flagMpesaOpenAPIPubKey,
		varMpesaOpenAPIPublicKey, usageMpesaOpenApiPubKey)
	rootCommand.PersistentFlags().Int64Var(&varMpesaSessionLifetimeMinutes, flagMpesaSessionLifetimeMinutes,
		varMpesaSessionLifetimeMinutes, usageMpesaSessionLifetimeMinutes)
	rootCommand.PersistentFlags().StringVar(&varMpesaServiceProviderCode, flagMpesaServiceProviderCode,
		varMpesaServiceProviderCode, usageMpesaServiceProviderCode)
	rootCommand.PersistentFlags().StringVar(&varMpesaTrustedSources, flagMpesaTrustedSources,
		varMpesaTrustedSources, usageMpesaTrustedSources)
	rootCommand.PersistentFlags().StringVar(&varTigoDisbursePIN, flagTigoDisbursePIN, varTigoDisbursePIN, usageTigoDisbursePIN)
	rootCommand.PersistentFlags().StringVar(&varTigoDisburseURL, flagTigoDisburseURL, varTigoDisburseURL, usageTigoDisburseURL)
	rootCommand.PersistentFlags().StringVar(&varTigoPushBillerCode, flagTigoPushBillerCode,
		varTigoPushBillerCode, usageTigoPushBillerCode)
	rootCommand.PersistentFlags().StringVar(&varTigoDisburseBrandID, flagTigoDisburseBrandID,
		varTigoDisburseBrandID, usageTigoDisburseBrandID)
	rootCommand.PersistentFlags().StringVar(&varTigoDisburseAccountMSISDN, flagTigoDisburseAccountMSISDN,
		varTigoDisburseAccountMSISDN, usageTigoDisburseAccountMSISDN)
	rootCommand.PersistentFlags().StringVar(&varTigoDisburseAccountName, flagTigoDisburseAccountName,
		varTigoDisburseAccountName, usageTigoDisburseAccountName)
	rootCommand.PersistentFlags().StringVar(&varTigoPushUsername, flagTigoPushUsername,
		varTigoPushUsername, usageTigoPushUsername)
	rootCommand.PersistentFlags().StringVar(&varTigoPushPassword, flagTigoPushPassword,
		varTigoPushPassword, usageTigoPushPassword)
	rootCommand.PersistentFlags().StringVar(&varTigoPushBaseURL, flagTigoPushBaseURL,
		varTigoPushBaseURL, usageTigoPushBaseURL)
	rootCommand.PersistentFlags().StringVar(&varTigoPushBillerMSISDN, flagTigoPushBillerMSISDN,
		varTigoPushBillerMSISDN, usageTigoPushBillerMSISDN)
	rootCommand.PersistentFlags().StringVar(&varTigoPushTokenURL, flagTigoPushTokenURL,
		varTigoPushTokenURL, usageTigoPushTokenURL)
	rootCommand.PersistentFlags().StringVar(&varTigoPushPayURL, flagTigoPushPayURL, varTigoPushPayURL, usageTigoPushPayURL)
	rootCommand.PersistentFlags().StringVar(&varTigoPasswordGrantType, flagTigoPasswordGrantType,
		varTigoPasswordGrantType, usageTigoPasswordGrantType)
	rootCommand.PersistentFlags().StringVar(&varConfigFile, "config", "",
		"config file (default is $HOME/.pesakit.yaml)")
	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCommand.SetVersionTemplate(versionTemplate)
	markHiddenExcept(rootCommand.Flags(), "help")
	app.root = rootCommand
	loadCommands(
		app.b2bCommand,
		app.configCommand,
		app.pushCommand,
		app.reverseCommand,
		app.sessionCommand,
		app.statusCommand,
		app.callbacksCommand,
		app.encryptCommand,
		app.docsCommand,
		app.disburseCommand,
	)

	app.setDebugMode(varDebugMode)

}

func loadCommands(fns ...func()) {
	for _, fn := range fns {
		fn()
	}
}
func (app *App) persistentPreRun(cmd *cobra.Command, args []string) {
	var (
		configFile          string
		configFileFlagGiven bool
		homeDirFlagGiven    bool
		err                 error
	)

	configFileFlagGiven = cmd.Flags().Changed(flagConfigFile)
	homeDirFlagGiven = cmd.Flags().Changed(flagHomeDirectory)
	// check if config file has been specified in the command line
	if configFileFlagGiven {
		specifiedConfigFile, err := cmd.PersistentFlags().GetString(flagConfigFile)
		if err != nil {
			logger := app.getLogger()
			_, _ = fmt.Fprintf(logger, "error: %v\n", err)
			return
		}
		configFile = specifiedConfigFile
	} else {
		// check if flagHomeDirectory is set
		if homeDirFlagGiven {
			homeDirectory, err := app.root.PersistentFlags().GetString(flagHomeDirectory)
			if err != nil {
				logger := app.getLogger()
				_, _ = fmt.Fprintf(logger, "error: %v\n", err)
				return
			}
			err = home.At(homeDirectory)
			if err != nil {
				logger := app.getLogger()
				_, _ = fmt.Fprintf(logger, "error: %v\n", err)
				return
			}
			appHomePath := filepath.Join(homeDirectory, ".pesakit")
			app.setHomeDir(appHomePath)
			configFile = filepath.Join(appHomePath, ".pesakit.env")
		} else {
			err := home.At("")
			if err != nil {
				logger := app.getLogger()
				_, _ = fmt.Fprintf(logger, "error: %v\n", err)
				return
			}
			homeDir, err := home.Get()
			if err != nil {
				logger := app.getLogger()
				_, _ = fmt.Fprintf(logger, "error: %v\n", err)
				return
			}
			appHomePath := filepath.Join(homeDir, ".pesakit")
			app.setHomeDir(appHomePath)
			configFile = filepath.Join(appHomePath, ".pesakit.env")
		}
	}

	if homeDirFlagGiven && !configFileFlagGiven {
		_ = env.LoadConfigFrom(filepath.Join(app.getHome(), ".env"))
	}
	err = env.LoadConfigFrom(configFile)
	if err != nil {
		logger := app.getLogger()
		_, _ = fmt.Fprintf(logger, "error: %v\n", err)
		return
	}

	err = app.loadConfigAndSetClients(cmd)
	if err != nil {
		logger := app.getLogger()
		_, _ = fmt.Fprintf(logger, "error: %v\n", err)
		return
	}
}

func (app *App) loadConfigAndSetClients(cmd *cobra.Command) error {

	logger := app.getLogger()
	debugMode := app.getDebugMode()

	// loads all configurations there is
	configMpesa, err := loadMpesaConfig(cmd)
	if err != nil {
		return err
	}

	configAirtel, err := loadAirtelConfig(cmd)
	if err != nil {
		return err
	}

	configTigo, err := loadTigoConfig(cmd)
	if err != nil {
		return err
	}
	clientMpesa := mpesa.NewClient(configMpesa, mpesa.WithLogger(logger), mpesa.WithDebugMode(debugMode))
	clientTigo := tigopesa.NewClient(configTigo, tigopesa.WithDebugMode(debugMode), tigopesa.WithLogger(logger))
	clientAirtel := airtel.NewClient(configAirtel, nil, true)
	app.setMpesaClient(clientMpesa)
	app.setTigoClient(clientTigo)
	app.setAirtelClient(clientAirtel)

	return nil
}
