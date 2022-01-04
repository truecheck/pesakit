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
	"github.com/techcraftlabs/airtel"
	"strings"
)

const (
	envAirtelProdBaseURL                    = "PESAKIT_AIRTEL_PROD_BASE_URL"
	envAirtelStagingBaseURL                 = "PESAKIT_AIRTEL_STAGING_BASE_URL"
	envAirtelAuthEndpoint                   = "PESAKIT_AIRTEL_AUTH_ENDPOINT"
	envAirtelPushEndpoint                   = "PESAKIT_AIRTEL_PUSH_ENDPOINT"
	envAirtelRefundEndpoint                 = "PESAKIT_AIRTEL_REFUND_ENDPOINT"
	envAirtelPushEnquiryEndpoint            = "PESAKIT_AIRTEL_PUSH_ENQUIRY_ENDPOINT"
	envAirtelDisbursementEndpoint           = "PESAKIT_AIRTEL_DISBURSEMENT_ENDPOINT"
	envAirtelDisbursementEnquiryEndpoint    = "PESAKIT_AIRTEL_DISBURSEMENT_ENQUIRY_ENDPOINT"
	envAirtelTransactionSummaryEndpoint     = "PESAKIT_AIRTEL_TRANSACTION_SUMMARY_ENDPOINT"
	envAirtelBalanceEnquiryEndpoint         = "PESAKIT_AIRTEL_BALANCE_ENQUIRY_ENDPOINT"
	envAirtelUserEnquiryEndpoint            = "PESAKIT_AIRTEL_USER_ENQUIRY_ENDPOINT"
	envAirtelAuthRegisteredCountries        = "PESAKIT_AIRTEL_AUTH_REGISTERED_COUNTRIES"
	envAirtelAccountRegisteredCountries     = "PESAKIT_AIRTEL_ACCOUNT_REGISTERED_COUNTRIES"
	envAirtelCollectionRegisteredCountries  = "PESAKIT_AIRTEL_COLLECTION_REGISTERED_COUNTRIES"
	envAirtelDisburseRegisteredCountries    = "PESAKIT_AIRTEL_DISBURSE_REGISTERED_COUNTRIES"
	envAirtelKYCRegisteredCountries         = "PESAKIT_AIRTEL_KYC_REGISTERED_COUNTRIES"
	envAirtelTransactionRegisteredCountries = "PESAKIT_AIRTEL_TRANSACTION_REGISTERED_COUNTRIES"
	envAirtelDisbursePIN                    = "PESAKIT_AIRTEL_DISBURSE_PIN"
	envAirtelCallbackPrivateKey             = "PESAKIT_AIRTEL_CALLBACK_PRIVATE_KEY"
	envAirtelCallbackAuth                   = "PESAKIT_AIRTEL_CALLBACK_AUTH"
	envAirtelPublicKey                      = "PESAKIT_AIRTEL_PUBLIC_KEY"
	envAirtelEnvironment                    = "PESAKIT_AIRTEL_ENVIRONMENT"
	envAirtelClientID                       = "PESAKIT_AIRTEL_CLIENT_ID"
	envAirtelSecret                         = "PESAKIT_AIRTEL_SECRET"
	defAirtelProdBaseURL                    = "https://openapi.airtel.africa"
	defAirtelStagingBaseURL                 = "https://openapiuat.airtel.africa"
	defAirtelAuthEndpoint                   = "/auth/oauth2/token"
	defAirtelPushEndpoint                   = "/merchant/v1/payments/"
	defAirtelRefundEndpoint                 = "/standard/v1/payments/refund"
	defAirtelPushEnquiryEndpoint            = "/standard/v1/payments/"
	defAirtelDisbursementEndpoint           = "/standard/v1/disbursements/"
	defAirtelDisbursementEnquiryEndpoint    = "/standard/v1/disbursements/"
	defAirtelTransactionSummaryEndpoint     = "/merchant/v1/transactions"
	defAirtelBalanceEnquiryEndpoint         = "/standard/v1/users/balance"
	defAirtelUserEnquiryEndpoint            = "/standard/v1/users/"
	defAirtelAuthRegisteredCountries        = "Tanzania"
	defAirtelAccountRegisteredCountries     = "Tanzania"
	defAirtelCollectionRegisteredCountries  = "Tanzania"
	defAirtelDisburseRegisteredCountries    = "Tanzania"
	defAirtelKYCRegisteredCountries         = "Tanzania"
	defAirtelTransactionRegisteredCountries = "Tanzania"
	defAirtelDisbursePIN                    = "123456"
	defAirtelCallbackPrivateKey             = ""
	defAirtelCallbackAuth                   = false
	defAirtelPublicKey                      = ""
	defAirtelEnvironment                    = "staging"
	defAirtelClientID                       = ""
	defAirtelSecret                         = ""
)

type (
	Airtel struct {
		AuthEndpoint                   string
		PushEndpoint                   string
		RefundEndpoint                 string
		PushEnquiryEndpoint            string
		DisbursementEndpoint           string
		DisbursementEnquiryEndpoint    string
		TransactionSummaryEndpoint     string
		BalanceEnquiryEndpoint         string
		UserEnquiryEndpoint            string
		AuthRegisteredCountries        string
		AccountRegisteredCountries     string
		CollectionRegisteredCountries  string
		DisburseRegisteredCountries    string
		KYCRegisteredCountries         string
		TransactionRegisteredCountries string
		DisbursePIN                    string
		CallbackPrivateKey             string
		CallbackAuth                   bool
		PublicKey                      string
		Environment                    string
		ClientID                       string
		ProdBaseURL                    string
		StagingBaseURL                 string
		Secret                         string
	}
)

func DefaultAirtelConfig() *Airtel {
	s := env.String
	b := env.Bool
	return &Airtel{
		AuthEndpoint:                   s(envAirtelAuthEndpoint, defAirtelAuthEndpoint),
		PushEndpoint:                   s(envAirtelPushEndpoint, defAirtelPushEndpoint),
		RefundEndpoint:                 s(envAirtelRefundEndpoint, defAirtelRefundEndpoint),
		PushEnquiryEndpoint:            s(envAirtelPushEnquiryEndpoint, defAirtelPushEnquiryEndpoint),
		DisbursementEndpoint:           s(envAirtelDisbursementEndpoint, defAirtelDisbursementEndpoint),
		DisbursementEnquiryEndpoint:    s(envAirtelDisbursementEnquiryEndpoint, defAirtelDisbursementEnquiryEndpoint),
		TransactionSummaryEndpoint:     s(envAirtelTransactionSummaryEndpoint, defAirtelTransactionSummaryEndpoint),
		BalanceEnquiryEndpoint:         s(envAirtelBalanceEnquiryEndpoint, defAirtelBalanceEnquiryEndpoint),
		UserEnquiryEndpoint:            s(envAirtelUserEnquiryEndpoint, defAirtelUserEnquiryEndpoint),
		AuthRegisteredCountries:        s(envAirtelAuthRegisteredCountries, defAirtelAuthRegisteredCountries),
		AccountRegisteredCountries:     s(envAirtelAccountRegisteredCountries, defAirtelAccountRegisteredCountries),
		CollectionRegisteredCountries:  s(envAirtelCollectionRegisteredCountries, defAirtelCollectionRegisteredCountries),
		DisburseRegisteredCountries:    s(envAirtelDisburseRegisteredCountries, defAirtelDisburseRegisteredCountries),
		KYCRegisteredCountries:         s(envAirtelKYCRegisteredCountries, defAirtelKYCRegisteredCountries),
		TransactionRegisteredCountries: s(envAirtelTransactionRegisteredCountries, defAirtelTransactionRegisteredCountries),
		DisbursePIN:                    s(envAirtelDisbursePIN, defAirtelDisbursePIN),
		CallbackPrivateKey:             s(envAirtelCallbackPrivateKey, defAirtelCallbackPrivateKey),
		CallbackAuth:                   b(envAirtelCallbackAuth, defAirtelCallbackAuth),
		PublicKey:                      s(envAirtelPublicKey, defAirtelPublicKey),
		Environment:                    s(envAirtelEnvironment, defAirtelEnvironment),
		ClientID:                       s(envAirtelClientID, defAirtelClientID),
		ProdBaseURL:                    s(envAirtelProdBaseURL, defAirtelProdBaseURL),
		StagingBaseURL:                 s(envAirtelStagingBaseURL, defAirtelStagingBaseURL),
		Secret:                         s(envAirtelSecret, defAirtelSecret),
	}
}

func (a *Airtel) Export() *airtel.Config {
	s := strings.Split
	return &airtel.Config{
		Endpoints: &airtel.Endpoints{
			AuthEndpoint:                a.AuthEndpoint,
			PushEndpoint:                a.PushEndpoint,
			RefundEndpoint:              a.RefundEndpoint,
			PushEnquiryEndpoint:         a.PushEnquiryEndpoint,
			DisbursementEndpoint:        a.DisbursementEndpoint,
			DisbursementEnquiryEndpoint: a.DisbursementEnquiryEndpoint,
			TransactionSummaryEndpoint:  a.TransactionSummaryEndpoint,
			BalanceEnquiryEndpoint:      a.BalanceEnquiryEndpoint,
			UserEnquiryEndpoint:         a.UserEnquiryEndpoint,
		},
		AllowedCountries: &airtel.RegisteredCountries{
			Auth:        s(a.AuthRegisteredCountries, " "),
			Account:     s(a.AccountRegisteredCountries, " "),
			Collection:  s(a.CollectionRegisteredCountries, " "),
			Disburse:    s(a.DisburseRegisteredCountries, " "),
			KYC:         s(a.KYCRegisteredCountries, " "),
			Transaction: s(a.TransactionRegisteredCountries, " "),
		},
		DisbursePIN:        a.DisbursePIN,
		CallbackPrivateKey: a.CallbackPrivateKey,
		CallbackAuth:       a.CallbackAuth,
		PublicKey:          a.PublicKey,
		Environment:        a.Environment,
		ClientID:           a.ClientID,
		ProdBaseURL:        a.ProdBaseURL,
		StagingBaseURL:     a.StagingBaseURL,
		Secret:             a.Secret,
	}
}
