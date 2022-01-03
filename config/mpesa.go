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
	envMpesaMarket                    = "PESAKIT_MPESA_MARKET"
	envMpesaAuthEndpoint              = "PESAKIT_MPESA_AUTH_ENDPOINT"
	envMpesaPushEndpoint              = "PESAKIT_MPESA_PUSH_ENDPOINT"
	envMpesaDisburseEndpoint          = "PESAKIT_MPESA_DISBURSE_ENDPOINT"
	envMpesaReversalEndpoint          = "PESAKIT_MPESA_REVERSAL_ENDPOINT"
	envMpesaB2BEndpoint               = "PESAKIT_MPESA_B2B_ENDPOINT"
	envMpesaDirectDebitCreateEndpoint = "PESAKIT_MPESA_DIRECT_DEBIT_CREATE_ENDPOINT"
	envMpesaDirectDebitPayEndpoint    = "PESAKIT_MPESA_DIRECT_DEBIT_PAY_ENDPOINT"
	envMpesaTransactionStatusEndpoint = "PESAKIT_MPESA_TRANSACTION_STATUS_ENDPOINT"
	envMpesaBaseURL                   = "PESAKIT_MPESA_BASE_URL"
	envMpesaAppName                   = "PESAKIT_MPESA_APP_NAME"
	envMpesaAppVersion                = "PESAKIT_MPESA_APP_VERSION"
	envMpesaAppDesc                   = "PESAKIT_MPESA_APP_DESCRIPTION"
	envMpesaSandboxAPIKey             = "PESAKIT_MPESA_SANDBOX_API_KEY" //nolint:gosec
	envMpesaOpenAPIKey                = "PESAKIT_MPESA_OPENAPI_KEY"     //nolint:gosec
	envMpesaSandboxPubKey             = "PESAKIT_MPESA_SANDBOX_PUBLIC_KEY"
	envMpesaOpenAPIPubKey             = "PESAKIT_MPESA_OPENAPI_PUBLIC_KEY"
	envMpesaSessionLifetimeMinutes    = "PESAKIT_MPESA_SESSION_LIFETIME_MINUTES"
	envMpesaServiceProviderCode       = "PESAKIT_MPESA_SERVICE_PROVIDER_CODE"
	envMpesaTrustedSources            = "PESAKIT_MPESA_TRUSTED_SOURCES"
	defMpesaBaseURL                   = "openapi.m-pesa.com"
	defMpesaAppName                   = "mpesa-app"
	defMpesaAppVersion                = "1.0"
	defMpesaAppDesc                   = "payment service"
	defMpesaPlatform                  = "sandbox"
	defMpesaMarket                    = "Tanzania"
	defMpesaAuthEndpoint              = "/getSession/"
	defMpesaPushEndpoint              = "/c2bPayment/singleStage/"
	defMpesaDisburseEndpoint          = "/b2cPayment/"
	defMpesaB2BEndpoint               = "/b2bPayment/"
	defMpesaReversalEndpoint          = "/reversal/"
	defMpesaTransactionStatusEndpoint = "/queryTransactionStatus"
	defMpesaDirectDebitCreateEndpoint = "/directDebitCreation/"
	defMpesaDirectDebitPayEndpoint    = "/directDebitPayment/"
	defMpesaSandboxAPIKey             = ""
	defMpesaOpenAPIKey                = ""
	defMpesaSandboxPubKey             = "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEArv9yxA69XQKBo24BaF/D+fvlqmGdYjqLQ5WtNBb5tquqGvAvG3WMFETVUSow/LizQalxj2ElMVrUmzu5mGGkxK08bWEXF7a1DEvtVJs6nppIlFJc2SnrU14AOrIrB28ogm58JjAl5BOQawOXD5dfSk7MaAA82pVHoIqEu0FxA8BOKU+RGTihRU+ptw1j4bsAJYiPbSX6i71gfPvwHPYamM0bfI4CmlsUUR3KvCG24rB6FNPcRBhM3jDuv8ae2kC33w9hEq8qNB55uw51vK7hyXoAa+U7IqP1y6nBdlN25gkxEA8yrsl1678cspeXr+3ciRyqoRgj9RD/ONbJhhxFvt1cLBh+qwK2eqISfBb06eRnNeC71oBokDm3zyCnkOtMDGl7IvnMfZfEPFCfg5QgJVk1msPpRvQxmEsrX9MQRyFVzgy2CWNIb7c+jPapyrNwoUbANlN8adU1m6yOuoX7F49x+OjiG2se0EJ6nafeKUXw/+hiJZvELUYgzKUtMAZVTNZfT8jjb58j8GVtuS+6TM2AutbejaCV84ZK58E2CRJqhmjQibEUO6KPdD7oTlEkFy52Y1uOOBXgYpqMzufNPmfdqqqSM4dU70PO8ogyKGiLAIxCetMjjm6FCMEA3Kc8K0Ig7/XtFm9By6VxTJK1Mg36TlHaZKP6VzVLXMtesJECAwEAAQ==" //nolint:lll
	defMpesaOpenAPIPubKey             = "MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAietPTdEyyoV/wvxRjS5pSn3ZBQH9hnVtQC9SFLgM9IkomEX9Vu9fBg2MzWSSqkQlaYIGFGH3d69Q5NOWkRo+Y8p5a61sc9hZ+ItAiEL9KIbZzhnMwi12jUYCTff0bVTsTGSNUePQ2V42sToOIKCeBpUtwWKhhW3CSpK7S1iJhS9H22/BT/pk21Jd8btwMLUHfVD95iXbHNM8u6vFaYuHczx966T7gpa9RGGXRtiOr3ScJq1515tzOSOsHTPHLTun59nxxJiEjKoI4Lb9h6IlauvcGAQHp5q6/2XmxuqZdGzh39uLac8tMSmY3vC3fiHYC3iMyTb7eXqATIhDUOf9mOSbgZMS19iiVZvz8igDl950IMcelJwcj0qCLoufLE5y8ud5WIw47OCVkD7tcAEPmVWlCQ744SIM5afw+Jg50T1SEtu3q3GiL0UQ6KTLDyDEt5BL9HWXAIXsjFdPDpX1jtxZavVQV+Jd7FXhuPQuDbh12liTROREdzatYWRnrhzeOJ5Se9xeXLvYSj8DmAI4iFf2cVtWCzj/02uK4+iIGXlX7lHP1W+tycLS7Pe2RdtC2+oz5RSSqb5jI4+3iEY/vZjSMBVk69pCDzZy4ZE8LBgyEvSabJ/cddwWmShcRS+21XvGQ1uXYLv0FCTEHHobCfmn2y8bJBb/Hct53BaojWUCAwEAAQ==" //nolint:lll
	defMpesaSessionLifetimeMinutes    = 60
	defMpesaServiceProviderCode       = ""
	defMpesaTrustedSources            = "https://openapi.m-pesa.com"
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
		SandboxAPIKey              string
		OpenAPIKey                 string
		SandboxPubKey              string
		OpenAPIPubKey              string
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
		apiKey = m.OpenAPIKey
		pubKey = m.OpenAPIPubKey
	} else {
		apiKey = m.SandboxAPIKey
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
		SandboxAPIKey:              env.String(envMpesaSandboxAPIKey, defMpesaSandboxAPIKey),
		OpenAPIKey:                 env.String(envMpesaOpenAPIKey, defMpesaOpenAPIKey),
		SandboxPubKey:              env.String(envMpesaSandboxPubKey, defMpesaSandboxPubKey),
		OpenAPIPubKey:              env.String(envMpesaOpenAPIPubKey, defMpesaOpenAPIPubKey),
		SessionLifetimeMinutes:     env.Int64(envMpesaSessionLifetimeMinutes, defMpesaSessionLifetimeMinutes),
		ServiceProviderCode:        env.String(envMpesaServiceProviderCode, defMpesaServiceProviderCode),
		TrustedSources:             env.String(envMpesaTrustedSources, defMpesaTrustedSources),
	}
}
