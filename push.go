package pesakit

import (
	"fmt"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
)

const (
	flagCollectAmount   = "amount"
	pFlagCollectAmount  = "a"
	defCollectAmount    = 0.0
	usageCollectAmount  = "amount to collect from the user"
	flagCollectPhoneNo  = "phone"
	pFlagCollectPhoneNo = "p"
	defCollectPhoneNo   = ""
	usageCollectPhoneNo = "phone number to send collection request"
)

func (app *App) pushCommand() {
	pushCommand := &cobra.Command{
		Use:   "push",
		Short: "send push pay request collect money from payer mobile money wallet",
		Long: `send a push pay request to the payer mobile money wallet
and collect money from the payee mobile money wallet`,
		Run: func(cmd *cobra.Command, args []string) {
			out := app.getWriter()
			// get amount and phone number
			amount, err := cmd.Flags().GetFloat64(flagCollectAmount)
			if err != nil {
				_, _ = fmt.Fprintf(out, "could not obtain amount %v\n", err)
				return
			}

			phone, err := cmd.Flags().GetString(flagCollectPhoneNo)
			if err != nil {
				_, _ = fmt.Fprintf(out, "could not obtain phone %v\n", err)
				return
			}
			operator, phoneNumber, err := mno.PhoneNumberInfo(phone)
			if err != nil {
				return
			}

			if operator == mno.Unknown {
				return
			}

			_, _ = fmt.Fprintf(out, "sending %s push pay request of amount %.2f to %s\n", operator, amount, phoneNumber)

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
	app.root.AddCommand(pushCommand)
}
