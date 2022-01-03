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

package config

import (
	"github.com/pesakit/pesakit/env"
	"github.com/techcraftlabs/mpesa"
	"strings"
)

const (
	envMpesaPlatform                  = "PESAKIT_MPESA_PLATFORM"
	defMpesaPlatform                  = "sandbox"
	envMpesaMarket                    = "PESAKIT_MPESA_MARKET"
	defMpesaMarket                    = "Tanzania"
	envMpesaAuthEndpoint              = "PESAKIT_MPESA_AUTH_ENDPOINT"
	defMpesaAuthEndpoint              = "/getSession/"
	envMpesaPushEndpoint              = "PESAKIT_MPESA_PUSH_ENDPOINT"
	defMpesaPushEndpoint              = ""
	envMpesaDisburseEndpoint          = "PESAKIT_MPESA_DISBURSE_ENDPOINT"
	defMpesaDisburseEndpoint          = ""
	envMpesaReversalEndpoint          = "PESAKIT_MPESA_REVERSAL_ENDPOINT"
	defMpesaReversalEndpoint          = ""
	envMpesaB2BEndpoint               = "PESAKIT_MPESA_B2B_ENDPOINT"
	defMpesaB2BEndpoint               = ""
	envMpesaDirectDebitCreateEndpoint = "PESAKIT_MPESA_DIRECT_DEBIT_CREATE_ENDPOINT"
	defMpesaDirectDebitCreateEndpoint = ""
	envMpesaDirectDebitPayEndpoint    = "PESAKIT_MPESA_DIRECT_DEBIT_PAY_ENDPOINT"
	defMpesaDirectDebitPayEndpoint    = ""
	envMpesaTransactionStatusEndpoint = "PESAKIT_MPESA_TRANSACTION_STATUS_ENDPOINT"
	defMpesaTransactionStatusEndpoint = ""
	envMpesaBaseURL                   = "PESAKIT_MPESA_BASE_URL"
	envMpesaAppName                   = "PESAKIT_MPESA_APP_NAME"
	envMpesaAppVersion                = "PESAKIT_MPESA_APP_VERSION"
	envMpesaAppDesc                   = "PESAKIT_MPESA_APP_DESCRIPTION"
	envMpesaSandboxAPIKey             = "PESAKIT_MPESA_SANDBOX_API_KEY"
	envMpesaOpenAPIKey                = "PESAKIT_MPESA_OPENAPI_KEY"
	envMpesaSandboxPubKey             = "PESAKIT_MPESA_SANDBOX_PUBLIC_KEY"
	envMpesaOpenApiPubKey             = "PESAKIT_MPESA_OPENAPI_PUBLIC_KEY"
	envMpesaSessionLifetimeMinutes    = "PESAKIT_MPESA_SESSION_LIFETIME_MINUTES"
	envMpesaServiceProviderCode       = "PESAKIT_MPESA_SERVICE_PROVIDER_CODE"
	envMpesaTrustedSources            = "PESAKIT_MPESA_TRUSTED_SOURCES"
	defMpesaBaseURL                   = "openapi.m-pesa.com"
	defMpesaAppName                   = "mpesa-app"
	defMpesaAppVersion                = "1.0"
	defMpesaAppDesc                   = "unified payment as a service"
	defMpesaSandboxAPIKey             = ""
	defMpesaOpenApiKey                = ""
	defMpesaSandboxPubKey             = ""
	defMpesaOpenApiPubKey             = ""
	defMpesaSessionLifetimeMinutes    = 60
	defMpesaServiceProviderCode       = ""
	defMpesaTrustedSources            = "openapi.m-pesa.com"
)

type (
	Mpesa struct {
		AuthEndpoint               string
		PushEndpoint               string
		DisburseEndpoint           string
		QueryEndpoint              string
		DirectDebitCreateEndpoint  string
		DirectDebitPayEndpoint     string
		TransactionReverseEndpoint string
		B2BEndpoint                string
		Name                       string
		Version                    string
		Description                string
		BasePath                   string
		Market                     string
		Platform                   string
		SandboxApiKey              string
		OpenApiKey                 string
		SandboxPubKey              string
		OpenApiPubKey              string
		SessionLifetimeMinutes     int64
		ServiceProviderCode        string
		TrustedSources             string
	}
)

func (m *Mpesa) Export() *mpesa.Config {

	var (
		apiKey string
		pubKey string
	)

	market := mpesa.MarketFmt(m.Market)
	platform := mpesa.PlatformFmt(m.Platform)

	if platform == mpesa.OPENAPI {
		apiKey = m.OpenApiKey
		pubKey = m.OpenApiPubKey
	} else {
		apiKey = m.SandboxApiKey
		pubKey = m.SandboxPubKey
	}

	return &mpesa.Config{
		Endpoints: &mpesa.Endpoints{
			AuthEndpoint:               m.AuthEndpoint,
			PushEndpoint:               m.PushEndpoint,
			DisburseEndpoint:           m.DisburseEndpoint,
			QueryEndpoint:              m.QueryEndpoint,
			DirectDebitCreateEndpoint:  m.DirectDebitCreateEndpoint,
			DirectDebitPayEndpoint:     m.DirectDebitPayEndpoint,
			TransactionReverseEndpoint: m.TransactionReverseEndpoint,
			B2BEndpoint:                m.B2BEndpoint,
		},
		Name:                   m.Name,
		Version:                m.Version,
		Description:            m.Description,
		BasePath:               m.BasePath,
		Market:                 market,
		Platform:               platform,
		APIKey:                 apiKey,
		PublicKey:              pubKey,
		SessionLifetimeMinutes: m.SessionLifetimeMinutes,
		ServiceProvideCode:     m.ServiceProviderCode,
		TrustedSources:         strings.Split(m.TrustedSources, " "),
	}
}

func DefaultMpesaConfig() *Mpesa {
	return &Mpesa{
		AuthEndpoint:               env.String(envMpesaAuthEndpoint, defMpesaAuthEndpoint),
		PushEndpoint:               env.String(envMpesaPushEndpoint, defMpesaPushEndpoint),
		DisburseEndpoint:           env.String(envMpesaDisburseEndpoint, defMpesaDisburseEndpoint),
		QueryEndpoint:              env.String(envMpesaTransactionStatusEndpoint, defMpesaTransactionStatusEndpoint),
		DirectDebitCreateEndpoint:  env.String(envMpesaDirectDebitCreateEndpoint, defMpesaDirectDebitCreateEndpoint),
		DirectDebitPayEndpoint:     env.String(envMpesaDirectDebitPayEndpoint, defMpesaDirectDebitPayEndpoint),
		TransactionReverseEndpoint: env.String(envMpesaReversalEndpoint, defMpesaReversalEndpoint),
		B2BEndpoint:                env.String(envMpesaB2BEndpoint, defMpesaB2BEndpoint),
		Name:                       env.String(envMpesaAppName, defMpesaAppName),
		Version:                    env.String(envMpesaAppVersion, defMpesaAppVersion),
		Description:                env.String(envMpesaAppDesc, defMpesaAppDesc),
		BasePath:                   env.String(envMpesaBaseURL, defMpesaBaseURL),
		Market:                     env.String(envMpesaMarket, defMpesaMarket),
		Platform:                   env.String(envMpesaPlatform, defMpesaPlatform),
		SandboxApiKey:              env.String(envMpesaSandboxAPIKey, defMpesaSandboxAPIKey),
		OpenApiKey:                 env.String(envMpesaOpenAPIKey, defMpesaOpenApiKey),
		SandboxPubKey:              env.String(envMpesaSandboxPubKey, defMpesaSandboxPubKey),
		OpenApiPubKey:              env.String(envMpesaOpenApiPubKey, defMpesaOpenApiPubKey),
		SessionLifetimeMinutes:     env.Int64(envMpesaSessionLifetimeMinutes, defMpesaSessionLifetimeMinutes),
		ServiceProviderCode:        env.String(envMpesaServiceProviderCode, defMpesaServiceProviderCode),
		TrustedSources:             env.String(envMpesaTrustedSources, defMpesaTrustedSources),
	}
}
