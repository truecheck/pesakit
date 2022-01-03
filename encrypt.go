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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
	"io"
)

const (
	flagEncryptEnv  = "env"
	pFlagEncryptEnv = "e"
	defEncryptEnv   = "sandbox"
	usageEncryptEnv = "environment to use for encryption"
	flagEncryptKey  = "key"
	pFlagEncryptKey = "k"
	defEncryptKey   = ""
	usageEncryptKey = "RSA public key to use for encryption"
	flagEncryptPin  = "pin"
	pFlagEncryptPin = "p"
	defEncryptPin   = ""
	usageEncryptPin = "pin to be encrypted"
	flagEncryptMno  = "mno"
	pFlagEncryptMno = "m"
	defEncryptMno   = ""
	usageEncryptMno = "mobile network operator"
)

func (app *App) encryptCommand() {
	// encryptCmd represents the encrypt command
	out := app.getWriter()
	rootCommand := app.root
	var encryptCmd = &cobra.Command{
		Use:   "encrypt",
		Short: "encrypt the api-key or pin using the provided public key",
		Long: `encrypt command takes the public key in base64 format decodes it and create
RSA public key from it. It then encrypts the api-key or pin using PKCS1v15.
The encrypted data is then encoded in base64 format and printed to stdout.

EXAMPLE:
	pesakit encrypt --key=<public-key> --pin=<pin>
	pesakit encrypt --mno=mpesa --env=sandbox
	pesakit encrypt --mno=airtel --env=production

As seen above you can specify key and pin or mno. If you specify MNO make sure
you have set the public key and pin/api-key for that MNO.
It is good to specify the environment as well. if nil, it will default to sandbox.
`,
		Run: func(cmd *cobra.Command, args []string) {
			encryptRun(cmd, args, out)
		},
	}

	//// flags
	//encryptCmd.PersistentFlags().StringP(flagEncryptEnv, pFlagEncryptEnv,
	//	defEncryptEnv, usageEncryptEnv)
	//encryptCmd.PersistentFlags().StringP(flagEncryptKey, pFlagEncryptKey,
	//	defEncryptKey, usageEncryptKey)
	//encryptCmd.PersistentFlags().StringP(flagEncryptPin, pFlagEncryptPin,
	//	defEncryptPin, usageEncryptPin)
	//encryptCmd.PersistentFlags().StringP(flagEncryptMno, pFlagEncryptMno,
	//	defEncryptMno, usageEncryptMno)
	//
	//encryptCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
	//	markHiddenExcept(rootCommand.PersistentFlags(), flagMpesaSandboxPubKey,
	//		flagMpesaSandboxApiKey, flagMpesaOpenApiKey, flagMpesaOpenAPIPubKey,
	//		flagAirtelPublicKey,
	//	)
	//	command.Parent().HelpFunc()(command, strings)
	//})

	rootCommand.AddCommand(encryptCmd)
}

func encryptRun(cmd *cobra.Command, args []string, out io.Writer) {
	var (
		//varEncryptEnv = cmd.Flag(flagEncryptEnv).Value.String()
		varEncryptKey = cmd.Flag(flagEncryptKey).Value.String()
		varEncryptPin = cmd.Flag(flagEncryptPin).Value.String()
		varEncryptMno = cmd.Flag(flagEncryptMno).Value.String()
	)
	// check if both key and pin are provided
	if varEncryptKey != "" && varEncryptPin != "" {
		encryptionRSA, err := pinEncryptionRSA(varEncryptPin, varEncryptKey)
		if err != nil {
			_, _ = fmt.Fprintf(out, "error: %v\n", err)

			return
		}
		_, _ = fmt.Fprintf(out, "%s\n", encryptionRSA)
	}

	// check if mno is provided
	if varEncryptMno != "" {
		mnoProvided := mno.Which(varEncryptMno)
		incorrectMno := mnoProvided == mno.Unknown || (mnoProvided != mno.Airtel && mnoProvided != mno.Vodacom)
		if incorrectMno {
			_, _ = fmt.Fprintf(out, "error: %v\n", "unsupported mno (supported: airtel, vodacom)")
			return
		}

		var (
			pin string
			key string
		)

		if mnoProvided == mno.Airtel {
			pin = varEncryptPin
			key = varEncryptKey
		} else {
			// this is mpesa
			// --mpesa-openapi-key string          M-Pesa openapi key
			// --mpesa-openapi-public-key string   M-Pesa openapi public key
			// --mpesa-sandbox-api-key string      M-Pesa sandbox api key
			// --mpesa-sandbox-public-key string   M-Pesa sandbox public key
			switch {
			//case varEncryptEnv == "sandbox":
			//	pin = varMpesaSandboxAPIKey
			//	key = varMpesaSandboxPublicKey
			//case varEncryptEnv == "production":
			//	pin = varMpesaOpenAPIKey
			//	key = varMpesaOpenAPIPublicKey
			default:
				_, _ = fmt.Fprintf(out, "error: %v\n", "unsupported env (supported: sandbox, production)")
			}
		}

		encryptionRSA, err := pinEncryptionRSA(pin, key)
		if err != nil {
			_, _ = fmt.Fprintf(out, "error: %v\n", err)

			return
		}

		_, _ = fmt.Fprintf(out, "%s\n", encryptionRSA)
	}
}

// pinEncryptionRSA encrypts the pin with the public key
// returns a base64 encoded string of the encrypted pin
func pinEncryptionRSA(pin string, pubKey string) (string, error) {
	decodedBase64, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return "", fmt.Errorf("could not decode pub key to Base64 string: %w", err)
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(decodedBase64)
	if err != nil {
		return "", fmt.Errorf("could not parse encoded public key (encryption key) : %w", err)
	}

	// check if the public key is RSA public key
	publicKey, isRSAPublicKey := publicKeyInterface.(*rsa.PublicKey)
	if !isRSAPublicKey {
		return "", fmt.Errorf("public key parsed is not an RSA public key : %w", err)
	}
	msg := []byte(pin)
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, msg)
	if err != nil {
		return "", fmt.Errorf("could not encrypt api key using generated public key: %w", err)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}
