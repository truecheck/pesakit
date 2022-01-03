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

package flags

import (
	"github.com/pesakit/pesakit/config"
	"github.com/spf13/cobra"
	"github.com/techcraftlabs/mpesa"
)

const (
	flagMpesaPlatform                   = "mpesa-platform"
	usageMpesaPlatform                  = "M-Pesa platform"
	flagMpesaMarket                     = "mpesa-market"
	usageMpesaMarket                    = "M-Pesa market"
	flagMpesaAuthEndpoint               = "mpesa-auth-endpoint"
	usageMpesaAuthEndpoint              = "M-Pesa auth endpoint"
	flagMpesaPushEndpoint               = "mpesa-push-endpoint"
	usageMpesaPushEndpoint              = "M-Pesa push endpoint"
	flagMpesaDisburseEndpoint           = "mpesa-disburse-endpoint"
	usageMpesaDisburseEndpoint          = "M-Pesa disburse endpoint"
	flagMpesaReversalEndpoint           = "mpesa-reversal-endpoint"
	usageMpesaReversalEndpoint          = "M-Pesa reversal endpoint"
	flagMpesaB2BEndpoint                = "mpesa-b2b-endpoint"
	usageMpesaB2BEndpoint               = "M-Pesa b2b endpoint"
	flagMpesaDirectDebitCreateEndpoint  = "mpesa-direct-debit-create-endpoint"
	usageMpesaDirectDebitCreateEndpoint = "M-Pesa direct debit create endpoint"
	flagMpesaDirectDebitPayEndpoint     = "mpesa-direct-debit-pay-endpoint"
	usageMpesaDirectDebitPayEndpoint    = "M-Pesa direct debit pay endpoint"
	flagMpesaTransactionStatusEndpoint  = "mpesa-transaction-status-endpoint"
	usageMpesaTransactionStatusEndpoint = "M-Pesa transaction status endpoint"
	flagMpesaBaseURL                    = "mpesa-base-url"
	usageMpesaBaseURL                   = "M-Pesa base url"
	flagMpesaAppName                    = "mpesa-App-name"
	usageMpesaAppName                   = "M-Pesa App name"
	flagMpesaAppVersion                 = "mpesa-App-version"
	usageMpesaAppVersion                = "M-Pesa App version"
	flagMpesaAppDesc                    = "mpesa-App-description"
	usageMpesaAppDesc                   = "M-Pesa App description"
	flagMpesaSandboxApiKey              = "mpesa-sandbox-api-key"
	usageMpesaSandboxApiKey             = "M-Pesa sandbox api key"
	flagMpesaOpenApiKey                 = "mpesa-openapi-key"
	usageMpesaOpenApiKey                = "M-Pesa openapi key"
	flagMpesaSandboxPubKey              = "mpesa-sandbox-public-key"
	usageMpesaSandboxPubKey             = "M-Pesa sandbox public key"
	flagMpesaOpenAPIPubKey              = "mpesa-openapi-public-key"
	usageMpesaOpenApiPubKey             = "M-Pesa openapi public key"
	flagMpesaSessionLifetimeMinutes     = "mpesa-session-lifetime-minutes"
	usageMpesaSessionLifetimeMinutes    = "M-Pesa session lifetime minutes"
	flagMpesaServiceProviderCode        = "mpesa-service-provider-code"
	usageMpesaServiceProviderCode       = "M-Pesa service provider code"
	flagMpesaTrustedSources             = "mpesa-trusted-sources"
	usageMpesaTrustedSources            = "M-Pesa trusted sources"
)

func SetMpesa(cmd *cobra.Command) {
	mpesaConfig := config.DefaultMpesaConfig()
	strVar := cmd.PersistentFlags().StringVar
	intVar := cmd.PersistentFlags().Int64Var
	strVar(&mpesaConfig.AuthEndpoint, flagMpesaAuthEndpoint, mpesaConfig.AuthEndpoint, usageMpesaAuthEndpoint)
	strVar(&mpesaConfig.PushEndpoint, flagMpesaPushEndpoint, mpesaConfig.PushEndpoint, usageMpesaPushEndpoint)
	strVar(&mpesaConfig.DisburseEndpoint, flagMpesaDisburseEndpoint, mpesaConfig.DisburseEndpoint,
		usageMpesaDisburseEndpoint)
	strVar(&mpesaConfig.TransactionReverseEndpoint, flagMpesaReversalEndpoint,
		mpesaConfig.TransactionReverseEndpoint, usageMpesaReversalEndpoint)
	strVar(&mpesaConfig.B2BEndpoint, flagMpesaB2BEndpoint, mpesaConfig.B2BEndpoint, usageMpesaB2BEndpoint)
	strVar(&mpesaConfig.DirectDebitCreateEndpoint, flagMpesaDirectDebitCreateEndpoint,
		mpesaConfig.DirectDebitCreateEndpoint, usageMpesaDirectDebitCreateEndpoint)
	strVar(&mpesaConfig.DirectDebitPayEndpoint, flagMpesaDirectDebitPayEndpoint,
		mpesaConfig.DirectDebitPayEndpoint, usageMpesaDirectDebitPayEndpoint)
	strVar(&mpesaConfig.QueryEndpoint, flagMpesaTransactionStatusEndpoint, mpesaConfig.QueryEndpoint,
		usageMpesaTransactionStatusEndpoint)
	strVar(&mpesaConfig.BasePath, flagMpesaBaseURL, mpesaConfig.BasePath, usageMpesaBaseURL)
	strVar(&mpesaConfig.Name, flagMpesaAppName, mpesaConfig.Name, usageMpesaAppName)
	strVar(&mpesaConfig.Version, flagMpesaAppVersion, mpesaConfig.Version, usageMpesaAppVersion)
	strVar(&mpesaConfig.Description, flagMpesaAppDesc, mpesaConfig.Description, usageMpesaAppDesc)
	strVar(&mpesaConfig.SandboxAPIKey, flagMpesaSandboxApiKey, mpesaConfig.SandboxAPIKey, usageMpesaSandboxApiKey)
	strVar(&mpesaConfig.OpenAPIKey, flagMpesaOpenApiKey, mpesaConfig.OpenAPIKey, usageMpesaOpenApiKey)
	strVar(&mpesaConfig.SandboxPubKey, flagMpesaSandboxPubKey, mpesaConfig.SandboxPubKey, usageMpesaSandboxPubKey)
	strVar(&mpesaConfig.OpenAPIPubKey, flagMpesaOpenAPIPubKey, mpesaConfig.OpenAPIPubKey, usageMpesaOpenApiPubKey)
	intVar(&mpesaConfig.SessionLifetimeMinutes, flagMpesaSessionLifetimeMinutes, mpesaConfig.SessionLifetimeMinutes,
		usageMpesaSessionLifetimeMinutes)
	strVar(&mpesaConfig.ServiceProviderCode, flagMpesaServiceProviderCode, mpesaConfig.ServiceProviderCode,
		usageMpesaServiceProviderCode)
	strVar(&mpesaConfig.TrustedSources, flagMpesaTrustedSources, mpesaConfig.TrustedSources,
		usageMpesaTrustedSources)
	strVar(&mpesaConfig.Market, flagMpesaMarket, mpesaConfig.Market, usageMpesaMarket)
	strVar(&mpesaConfig.Platform, flagMpesaPlatform, mpesaConfig.Platform, usageMpesaPlatform)
}

func GetMpesaConfig(command *cobra.Command) (*mpesa.Config, error) {
	getStr := command.PersistentFlags().GetString
	authEndpoint, err := getStr(flagMpesaAuthEndpoint)
	if err != nil {
		return nil, err
	}
	pushEndpoint, err := getStr(flagMpesaPushEndpoint)
	if err != nil {
		return nil, err
	}
	disburseEndpoint, err := getStr(flagMpesaDisburseEndpoint)
	if err != nil {
		return nil, err
	}
	queryEndpoint, err := getStr(flagMpesaTransactionStatusEndpoint)
	if err != nil {
		return nil, err
	}
	directCreateEndpoint, err := getStr(flagMpesaDirectDebitCreateEndpoint)
	if err != nil {
		return nil, err
	}
	directPayEndpoint, err := getStr(flagMpesaDirectDebitPayEndpoint)
	if err != nil {
		return nil, err
	}
	reversalEndpoint, err := getStr(flagMpesaReversalEndpoint)
	if err != nil {
		return nil, err
	}
	b2bEndpoint, err := getStr(flagMpesaB2BEndpoint)
	if err != nil {
		return nil, err
	}
	trustedSourcesStr, err := getStr(flagMpesaTrustedSources)
	if err != nil {
		return nil, err
	}

	providerCode, err := getStr(flagMpesaServiceProviderCode)
	if err != nil {
		return nil, err
	}
	sessionLifetime, err := command.PersistentFlags().GetInt64(flagMpesaSessionLifetimeMinutes)
	if err != nil {
		return nil, err
	}
	name, err := getStr(flagMpesaAppName)
	if err != nil {
		return nil, err
	}
	version, err := getStr(flagMpesaAppVersion)
	if err != nil {
		return nil, err
	}
	desc, err := getStr(flagMpesaAppDesc)
	if err != nil {
		return nil, err
	}
	basePath, err := getStr(flagMpesaAppDesc)
	if err != nil {
		return nil, err
	}
	marketString, err := getStr(flagMpesaMarket)
	if err != nil {
		return nil, err
	}

	platformStr, err := getStr(flagMpesaPlatform)
	if err != nil {
		return nil, err
	}

	// get production public key and api-key
	openAapiKey, err := getStr(flagMpesaOpenApiKey)
	if err != nil {
		return nil, err
	}

	openApiPubKey, err := getStr(flagMpesaOpenAPIPubKey)
	if err != nil {
		return nil, err
	}

	sandboxApiKey, err := getStr(flagMpesaSandboxApiKey)
	if err != nil {
		return nil, err
	}

	sandboxPubKey, err := getStr(flagMpesaSandboxPubKey)
	if err != nil {
		return nil, err
	}

	c := &config.Mpesa{
		AuthEndpoint:               authEndpoint,
		PushEndpoint:               pushEndpoint,
		DisburseEndpoint:           disburseEndpoint,
		QueryEndpoint:              queryEndpoint,
		DirectDebitCreateEndpoint:  directCreateEndpoint,
		DirectDebitPayEndpoint:     directPayEndpoint,
		TransactionReverseEndpoint: reversalEndpoint,
		B2BEndpoint:                b2bEndpoint,
		Name:                       name,
		Version:                    version,
		Description:                desc,
		BasePath:                   basePath,
		Market:                     marketString,
		Platform:                   platformStr,
		SandboxAPIKey:              sandboxApiKey,
		OpenAPIKey:                 openAapiKey,
		SandboxPubKey:              sandboxPubKey,
		OpenAPIPubKey:              openApiPubKey,
		SessionLifetimeMinutes:     sessionLifetime,
		ServiceProviderCode:        providerCode,
		TrustedSources:             trustedSourcesStr,
	}

	return c.Export(), nil
}
