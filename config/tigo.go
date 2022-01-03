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
	"github.com/techcraftlabs/tigopesa"
)

const (
	envTigoDisbursePIN           = "PESAKIT_TIGO_DISBURSE_PIN"
	envTigoDisburseURL           = "PESAKIT_TIGO_DISaBURSE_URL"
	envTigoDisburseBrandID       = "PESAKIT_TIGO_DISBURSE_BRAND_ID"
	envTigoDisburseAccountMSISDN = "PESAKIT_TIGO_DISBURSE_ACCOUNT_MSISDN"
	envTigoDisburseAccountName   = "PESAKIT_TIGO_DISBURSE_ACCOUNT_NAME"
	envTigoPushUsername          = "PESAKIT_TIGO_PUSH_USERNAME"
	envTigoPushPassword          = "PESAKIT_TIGO_PUSH_PASSWORD"
	envTigoPushBillerMSISDN      = "PESAKIT_TIGO_PUSH_BILLER_MSISDN"
	envTigoPushBaseURL           = "PESAKIT_TIGO_PUSH_BASE_URL"
	envTigoPushBillerCode        = "PESAKIT_TIGO_PUSH_BILLER_CODE"
	envTigoPushTokenURL          = "PESAKIT_TIGO_PUSH_TOKEN_URL"
	envTigoPushPayURL            = "PESAKIT_TIGO_PUSH_PAY_URL"
	envTigoPasswordGrantType     = "PESAKIT_TIGO_PASSWORD_GRANT_TYPE"
	defTigoDisburseAccountName   = ""
	defTigoDisburseAccountMSISDN = ""
	defTigoDisburseBrandID       = ""
	defTigoDisbursePIN           = ""
	defTigoDisburseURL           = ""
	defTigoPushUsername          = ""
	defTigoPushPassword          = ""
	defTigoPushBaseURL           = ""
	defTigoPushBillerMSISDN      = ""
	defTigoPushBillerCode        = ""
	defTigoPushTokenURL          = ""
	defTigoPushPayURL            = ""
	defTigoPasswordGrantType     = "password"
)

type TigoPesa struct {
	DisburseAccountName   string
	DisburseAccountMSISDN string
	DisburseBrandID       string
	DisbursePIN           string
	DisburseRequestURL    string
	PushUsername          string
	PushPassword          string
	PasswordGrantType     string
	PushBaseURL           string
	PushTokenEndpoint     string
	PushBillerMSISDN      string
	PushBillerCode        string
	PushPayEndpoint       string
}

func DefaultTigoPesaConfig() *TigoPesa {
	def := env.String
	return &TigoPesa{
		DisburseAccountName:   def(envTigoDisburseAccountName, defTigoDisburseAccountName),
		DisburseAccountMSISDN: def(envTigoDisburseAccountMSISDN, defTigoDisburseAccountMSISDN),
		DisburseBrandID:       def(envTigoDisburseBrandID, defTigoDisburseBrandID),
		DisbursePIN:           def(envTigoDisbursePIN, defTigoDisbursePIN),
		DisburseRequestURL:    def(envTigoDisburseURL, defTigoDisburseURL),
		PushUsername:          def(envTigoPushUsername, defTigoPushUsername),
		PushPassword:          def(envTigoPushPassword, defTigoPushPassword),
		PasswordGrantType:     def(envTigoPasswordGrantType, defTigoPasswordGrantType),
		PushBaseURL:           def(envTigoPushBaseURL, defTigoPushBaseURL),
		PushTokenEndpoint:     def(envTigoPushTokenURL, defTigoPushTokenURL),
		PushBillerMSISDN:      def(envTigoPushBillerMSISDN, defTigoPushBillerMSISDN),
		PushBillerCode:        def(envTigoPushBillerCode, defTigoPushBillerCode),
		PushPayEndpoint:       def(envTigoPushPayURL, defTigoPushPayURL),
	}
}

func (t *TigoPesa) Export() *tigopesa.Config {
	config := &tigopesa.Config{
		Disburse: &tigopesa.DisburseConfig{
			AccountName:   t.DisburseAccountName,
			AccountMSISDN: t.DisburseAccountMSISDN,
			BrandID:       t.DisburseBrandID,
			PIN:           t.DisbursePIN,
			RequestURL:    t.DisburseRequestURL,
		},
		Push: &tigopesa.PushConfig{
			Username:          t.PushUsername,
			Password:          t.PushPassword,
			PasswordGrantType: t.PasswordGrantType,
			BaseURL:           t.PushBaseURL,
			TokenEndpoint:     t.PushTokenEndpoint,
			BillerMSISDN:      t.PushBillerMSISDN,
			BillerCode:        t.PushBillerCode,
			PushPayEndpoint:   t.PushPayEndpoint,
		},
	}

	return config
}
