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
)

const (
	flagAirtelAuthEndpoint                    = "airtel-auth-endpoint"
	flagAirtelPushEndpoint                    = "airtel-push-endpoint"
	flagAirtelRefundEndpoint                  = "airtel-refund-endpoint"
	flagAirtelPushEnquiryEndpoint             = "airtel-push-enquiry-endpoint"
	flagAirtelDisbursementEndpoint            = "airtel-disbursement-endpoint"
	flagAirtelDisbursementEnquiryEndpoint     = "airtel-disbursement-enquiry-endpoint"
	flagAirtelTransactionSummaryEndpoint      = "airtel-transaction-summary-endpoint"
	flagAirtelBalanceEnquiryEndpoint          = "airtel-balance-enquiry-endpoint"
	flagAirtelUserEnquiryEndpoint             = "airtel-user-enquiry-endpoint"
	flagAirtelAuthRegisteredCountries         = "airtel-auth-countries"
	flagAirtelAccountRegisteredCountries      = "airtel-account-countries"
	flagAirtelCollectionRegisteredCountries   = "airtel-collection-countries"
	flagAirtelDisburseRegisteredCountries     = "airtel-disburse-countries"
	flagAirtelKYCRegisteredCountries          = "airtel-kyc-countries"
	flagAirtelTransactionRegisteredCountries  = "airtel-transaction-countries"
	flagAirtelDisbursePIN                     = "airtel-disburse-pin"
	flagAirtelCallbackPrivateKey              = "airtel-callback-private-key"
	flagAirtelCallbackAuth                    = "airtel-callback-auth"
	flagAirtelPublicKey                       = "airtel-public-key"
	flagAirtelEnvironment                     = "airtel-environment"
	flagAirtelClientID                        = "airtel-client-id"
	flagAirtelProdBaseURL                     = "airtel-prod-base-url"
	flagAirtelStagingBaseURL                  = "airtel-staging-base-url"
	flagAirtelSecret                          = "airtel-secret"
	usageAirtelAuthEndpoint                   = "Airtel Auth Endpoint"
	usageAirtelPushEndpoint                   = "Airtel Push Endpoint"
	usageAirtelRefundEndpoint                 = "Airtel Refund Endpoint"
	usageAirtelPushEnquiryEndpoint            = "Airtel Push Enquiry Endpoint"
	usageAirtelDisbursementEndpoint           = "Airtel Disbursement Endpoint"
	usageAirtelDisbursementEnquiryEndpoint    = "Airtel Disbursement Enquiry Endpoint"
	usageAirtelTransactionSummaryEndpoint     = "Airtel Transaction Summary Endpoint"
	usageAirtelBalanceEnquiryEndpoint         = "Airtel Balance Enquiry Endpoint"
	usageAirtelUserEnquiryEndpoint            = "Airtel User Enquiry Endpoint"
	usageAirtelAuthRegisteredCountries        = "Airtel Auth Registered Countries"
	usageAirtelAccountRegisteredCountries     = "Airtel Account Registered Countries"
	usageAirtelCollectionRegisteredCountries  = "Airtel Collection Registered Countries"
	usageAirtelDisburseRegisteredCountries    = "Airtel Disburse Registered Countries"
	usageAirtelKYCRegisteredCountries         = "Airtel KYC Registered Countries"
	usageAirtelTransactionRegisteredCountries = "Airtel Transaction Registered Countries"
	usageAirtelDisbursePIN                    = "Airtel Disburse PIN"
	usageAirtelCallbackPrivateKey             = "Airtel Callback Private Key"
	usageAirtelCallbackAuth                   = "Airtel Callback Auth"
	usageAirtelPublicKey                      = "Airtel Public Key"
	usageAirtelEnvironment                    = "Airtel Environment"
	usageAirtelClientID                       = "Airtel Client ID"
	usageAirtelProdBaseURL                    = "Airtel Prod Base URL"
	usageAirtelStagingBaseURL                 = "Airtel Staging Base URL"
	usageAirtelSecret                         = "Airtel Secret" //nolint:gosec
)

func SetAirtel(cmd *cobra.Command) {
	airtelDefaultConfig := config.DefaultAirtelConfig()
	s := cmd.PersistentFlags().String
	b := cmd.PersistentFlags().Bool
	s(flagAirtelAuthEndpoint, airtelDefaultConfig.AuthEndpoint, usageAirtelAuthEndpoint)
	s(flagAirtelPushEndpoint, airtelDefaultConfig.PushEndpoint, usageAirtelPushEndpoint)
	s(flagAirtelRefundEndpoint, airtelDefaultConfig.RefundEndpoint, usageAirtelRefundEndpoint)
	s(flagAirtelPushEnquiryEndpoint, airtelDefaultConfig.PushEnquiryEndpoint, usageAirtelPushEnquiryEndpoint)
	s(flagAirtelDisbursementEndpoint, airtelDefaultConfig.DisbursementEndpoint, usageAirtelDisbursementEndpoint)
	s(flagAirtelDisbursementEnquiryEndpoint, airtelDefaultConfig.DisbursementEnquiryEndpoint,
		usageAirtelDisbursementEnquiryEndpoint)
	s(flagAirtelTransactionSummaryEndpoint, airtelDefaultConfig.TransactionSummaryEndpoint,
		usageAirtelTransactionSummaryEndpoint)
	s(flagAirtelBalanceEnquiryEndpoint, airtelDefaultConfig.BalanceEnquiryEndpoint, usageAirtelBalanceEnquiryEndpoint)
	s(flagAirtelUserEnquiryEndpoint, airtelDefaultConfig.UserEnquiryEndpoint, usageAirtelUserEnquiryEndpoint)
	s(flagAirtelAuthRegisteredCountries, airtelDefaultConfig.AuthRegisteredCountries, usageAirtelAuthRegisteredCountries)
	s(flagAirtelAccountRegisteredCountries, airtelDefaultConfig.AccountRegisteredCountries,
		usageAirtelAccountRegisteredCountries)
	s(flagAirtelCollectionRegisteredCountries, airtelDefaultConfig.CollectionRegisteredCountries,
		usageAirtelCollectionRegisteredCountries)
	s(flagAirtelDisburseRegisteredCountries, airtelDefaultConfig.DisburseRegisteredCountries,
		usageAirtelDisburseRegisteredCountries)
	s(flagAirtelKYCRegisteredCountries, airtelDefaultConfig.KYCRegisteredCountries, usageAirtelKYCRegisteredCountries)
	s(flagAirtelTransactionRegisteredCountries, airtelDefaultConfig.TransactionRegisteredCountries,
		usageAirtelTransactionRegisteredCountries)
	s(flagAirtelDisbursePIN, airtelDefaultConfig.DisbursePIN, usageAirtelDisbursePIN)
	s(flagAirtelCallbackPrivateKey, airtelDefaultConfig.CallbackPrivateKey, usageAirtelCallbackPrivateKey)
	b(flagAirtelCallbackAuth, airtelDefaultConfig.CallbackAuth, usageAirtelCallbackAuth)
	s(flagAirtelPublicKey, airtelDefaultConfig.PublicKey, usageAirtelPublicKey)
	s(flagAirtelEnvironment, airtelDefaultConfig.Environment, usageAirtelEnvironment)
	s(flagAirtelClientID, airtelDefaultConfig.ClientID, usageAirtelClientID)
	s(flagAirtelProdBaseURL, airtelDefaultConfig.ProdBaseURL, usageAirtelProdBaseURL)
	s(flagAirtelStagingBaseURL, airtelDefaultConfig.StagingBaseURL, usageAirtelStagingBaseURL)
	s(flagAirtelSecret, airtelDefaultConfig.Secret, usageAirtelSecret)
}
