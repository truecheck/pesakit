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

const (
	envAirtelPublicKey             = "PK_AIRTEL_PUBLIC_KEY"
	flagAirtelPublicKey            = "airtel-public-key"
	usageAirtelPublicKey           = "Airtel public key"
	defAirtelPublicKey             = ""
	envAirtelDisbursePin           = "PK_AIRTEL_DISBURSE_PIN"
	flagAirtelDisbursePin          = "airtel-disburse-pin"
	usageAirtelDisbursePin         = "Airtel disburse pin"
	defAirtelDisbursePin           = ""
	envAirtelClientId              = "PK_AIRTEL_CLIENT_ID"
	flagAirtelClientId             = "airtel-client-id"
	usageAirtelClientId            = "Airtel client id"
	defAirtelClientId              = ""
	envAirtelClientSecret          = "PK_AIRTEL_CLIENT_SECRET"
	flagAirtelClientSecret         = "airtel-client-secret"
	usageAirtelClientSecret        = "Airtel client secret"
	defAirtelClientSecret          = ""
	envAirtelDeploymentEnv         = "PK_AIRTEL_DEPLOYMENT"
	flagAirtelDeploymentEnv        = "airtel-deployment"
	usageAirtelDeploymentEnv       = "Airtel deployment environment"
	defAirtelDeploymentEnv         = "staging"
	envAirtelCallbackAuth          = "PK_AIRTEL_CALLBACK_AUTH"
	flagAirtelCallbackAuth         = "airtel-callback-auth"
	usageAirtelCallbackAuth        = "Airtel callback auth"
	defAirtelCallbackAuth          = false
	envAirtelCallbackPrivKey       = "PK_AIRTEL_CALLBACK_PRIVATE_KEY"
	flagAirtelCallbackPrivKey      = "airtel-callback-private-key"
	usageAirtelCallbackPrivKey     = "Airtel callback private key"
	defAirtelCallbackPrivKey       = "zITVAAGYSlzl1WkUQJn81kbpT5drH3koffT8jCkcJJA="
	envAirtelCountries             = "PK_AIRTEL_COUNTRIES"
	flagAirtelCountries            = "airtel-countries"
	usageAirtelCountries           = "Airtel countries"
	defAirtelCountries             = "tanzania"
	envTigoDisbursePIN             = "PK_TIGO_DISBURSE_PIN"
	flagTigoDisbursePIN            = "tigo-disburse-pin"
	usageTigoDisbursePIN           = "Tigo disburse PIN"
	envTigoDisburseURL             = "PK_TIGO_DISBURSE_URL"
	flagTigoDisburseURL            = "tigo-disburse-url"
	usageTigoDisburseURL           = "Tigo disburse URL"
	envTigoDisburseBrandID         = "PK_TIGO_DISBURSE_BRAND_ID"
	flagTigoDisburseBrandID        = "tigo-disburse-brand-id"
	usageTigoDisburseBrandID       = "Tigo disburse brand id"
	envTigoDisburseAccountMSISDN   = "PK_TIGO_DISBURSE_ACCOUNT_MSISDN"
	flagTigoDisburseAccountMSISDN  = "tigo-disburse-account-msisdn"
	usageTigoDisburseAccountMSISDN = "Tigo disburse account MSISDN"
	envTigoDisburseAccountName     = "PK_TIGO_DISBURSE_ACCOUNT_NAME"
	flagTigoDisburseAccountName    = "tigo-disburse-account-name"
	usageTigoDisburseAccountName   = "Tigo disburse account name"
	envTigoPushUsername            = "PK_TIGO_PUSH_USERNAME"
	flagTigoPushUsername           = "tigo-push-username"
	usageTigoPushUsername          = "Tigo push username"
	envTigoPushPassword            = "PK_TIGO_PUSH_PASSWORD" //nolint:gosec
	flagTigoPushPassword           = "tigo-push-password"
	usageTigoPushPassword          = "Tigo push password"
	envTigoPushBillerMSISDN        = "PK_TIGO_PUSH_BILLER_MSISDN"
	flagTigoPushBillerMSISDN       = "tigo-push-biller-msisdn"
	usageTigoPushBillerMSISDN      = "Tigo push biller MSISDN"
	envTigoPushBaseURL             = "PK_TIGO_PUSH_BASE_URL"
	flagTigoPushBaseURL            = "tigo-push-base-url"
	usageTigoPushBaseURL           = "Tigo push base URL"
	envTigoPushBillerCode          = "PK_TIGO_PUSH_BILLER_CODE"
	flagTigoPushBillerCode         = "tigo-push-biller-code"
	usageTigoPushBillerCode        = "Tigo push biller code"
	envTigoPushTokenURL            = "PK_TIGO_PUSH_TOKEN_URL" //nolint:gosec
	flagTigoPushTokenURL           = "tigo-push-token-url"    //nolint:gosec
	usageTigoPushTokenURL          = "Tigo push token URL"    //nolint:gosec
	envTigoPushPayURL              = "PK_TIGO_PUSH_PAY_URL"
	flagTigoPushPayURL             = "tigo-push-pay-url"
	usageTigoPushPayURL            = "Tigo push pay URL"
	envTigoPasswordGrantType       = "PK_TIGO_PASSWORD_GRANT_TYPE"
	flagTigoPasswordGrantType      = "tigo-password-grant-type"
)
