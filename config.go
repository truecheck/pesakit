/*
 * MIT License
 *
 * Copyright (c) 2021 PESAKIT - MOBILE MONEY TOOLBOX
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package pesakit

import (
	"github.com/pesakit/pesakit/flags"
	"github.com/spf13/cobra"
	"github.com/techcraftlabs/airtel"
	"github.com/techcraftlabs/tigopesa"
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
	configCommand := &cobra.Command{
		Use:   "config",
		Short: "manages clients configuration",
		Long: `manages clients configuration. This command is used to create, view,update and delete
clients configurations`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	setFlagsFunc := func() {
		flags.SetMnoFlag(configCommand, flags.PERSISTENT)
		flags.SetMpesa(configCommand)
	}
	setFlagsFunc()

	rootCommand.AddCommand(configCommand)
}

//func initConfigPrintCommand(parentCommand *cobra.Command, out io.Writer) {
//	configPrintCommand := &cobra.Command{
//		Use:   "print",
//		Short: "prints the configuration",
//		Long:  `prints the configuration if not any is specified all configurations are printed`,
//		Run: func(cmd *cobra.Command, args []string) {
//			airtelConfigBool, err := cmd.Flags().GetBool(flagConfigAirtel)
//			if err != nil {
//				return
//			}
//			mpesaConfigBool, err := cmd.Flags().GetBool(flagConfigMpesa)
//			if err != nil {
//				return
//			}
//			tigoConfigBool, err := cmd.Flags().GetBool(flagConfigTigo)
//			if err != nil {
//				return
//			}
//
//			notAnySpecified := !(airtelConfigBool || mpesaConfigBool || tigoConfigBool)
//			if notAnySpecified {
//				mpesaConfig, err := flags.GetMpesaConfig(cmd)
//				if err != nil {
//					return
//				}
//				_, _ = fmt.Fprintf(out, "Mpesa Config: %v\n", mpesaConfig)
//				airtelConfig, err := loadAirtelConfig(cmd)
//				if err != nil {
//					return
//				}
//				_, _ = fmt.Fprintf(out, "Airtel Config: %v\n", airtelConfig)
//				tigoConfig, err := loadTigoConfig(cmd)
//				if err != nil {
//					return
//				}
//				_, _ = fmt.Fprintf(out, "Tigo Config: %v\n", tigoConfig)
//			}
//
//		},
//	}
//
//	//configPrintCommand.SetHelpFunc(func(command *cobra.Command, strings []string) {
//	//	set := parentCommand.Parent().PersistentFlags()
//	//	markHiddenExcept(set, flagDebugMode, flagConfigAirtel, flagConfigMpesa, flagConfigTigo)
//	//	command.Parent().HelpFunc()(command, strings)
//	//})
//
//	parentCommand.AddCommand(configPrintCommand)
//}
//

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
