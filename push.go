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
)

func (app *App) pushCommand() {
	pushCommand := &cobra.Command{
		Use:   "push",
		Short: "send push pay request collect money from payer mobile money wallet",
		Long: `send a push pay request to the payer mobile money wallet
and collect money from the payee mobile money wallet`,
		Run: func(cmd *cobra.Command, args []string) {
			request, err := flags.GetTransactionRequest(cmd)
			if err != nil {
				app.logger.Fatal(err)
				return
			}
			app.Logger().Printf("%+v\n", request)
		},
	}
	//pushCommand.PersistentFlags().Float64P(flagCollectAmount, pFlagCollectAmount, defCollectAmount, usageCollectAmount)
	//pushCommand.PersistentFlags().StringP(flagCollectPhoneNo, pFlagCollectPhoneNo, defCollectPhoneNo, usageCollectPhoneNo)
	//_ = markFlagsRequired(pushCommand, globalFlagType, flagCollectAmount, flagCollectPhoneNo)
	//pushCommand.SetHelpFunc(func(command *cobra.Command, strings []string) {
	//	markHiddenExcept(app.root.PersistentFlags(),
	//		flagMpesaSandboxPubKey,
	//		flagMpesaSandboxApiKey,
	//		flagMpesaOpenApiKey,
	//		flagMpesaOpenAPIPubKey,
	//		flagMpesaBaseURL,
	//		flagMpesaMarket,
	//		flagMpesaPlatform,
	//		flagMpesaAuthEndpoint,
	//		flagMpesaPushEndpoint,
	//		flagMpesaServiceProviderCode,
	//		flagAirtelPublicKey,
	//		flagAirtelDeploymentEnv,
	//		flagAirtelClientId,
	//		flagAirtelClientSecret,
	//		flagAirtelCountries,
	//		flagTigoPasswordGrantType,
	//		flagTigoPushBaseURL,
	//		flagTigoPushTokenURL,
	//		flagTigoPushPayURL,
	//		flagTigoPushUsername,
	//		flagTigoPushPassword,
	//	)
	//	command.Parent().HelpFunc()(command, strings)
	//})
	setFlags := func() {
		flags.SetMpesa(pushCommand)
		flags.SetTransactionRequestFlags(pushCommand)
	}

	setFlags()
	app.root.AddCommand(pushCommand)
}
