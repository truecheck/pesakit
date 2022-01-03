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
	"strings"
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
		flags.SetMno(configCommand, flags.PERSISTENT)
		flags.SetMpesa(configCommand)
		flags.SetTigoPesa(configCommand)
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

func airtelEnv(value string) airtel.Environment {
	if value == "production" || value == "prod" {
		return airtel.PRODUCTION
	} else if value == "sandbox" || value == "staging" {
		return airtel.STAGING
	} else {
		return airtel.STAGING
	}
}
