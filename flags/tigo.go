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
	"github.com/techcraftlabs/tigopesa"
)

const (
	flagTigoDisbursePIN            = "tigo-disburse-pin"
	usageTigoDisbursePIN           = "Tigo disburse PIN"
	flagTigoDisburseURL            = "tigo-disburse-url"
	usageTigoDisburseURL           = "Tigo disburse URL"
	flagTigoDisburseBrandID        = "tigo-disburse-brand-id"
	usageTigoDisburseBrandID       = "Tigo disburse brand id"
	flagTigoDisburseAccountMSISDN  = "tigo-disburse-account-msisdn"
	usageTigoDisburseAccountMSISDN = "Tigo disburse account MSISDN"
	flagTigoDisburseAccountName    = "tigo-disburse-account-name"
	usageTigoDisburseAccountName   = "Tigo disburse account name"
	flagTigoPushUsername           = "tigo-push-username"
	usageTigoPushUsername          = "Tigo push username"
	flagTigoPushPassword           = "tigo-push-password"
	usageTigoPushPassword          = "Tigo push password"
	flagTigoPushBillerMSISDN       = "tigo-push-biller-msisdn"
	usageTigoPushBillerMSISDN      = "Tigo push biller MSISDN"
	flagTigoPushBaseURL            = "tigo-push-base-url"
	usageTigoPushBaseURL           = "Tigo push base URL"
	flagTigoPushBillerCode         = "tigo-push-biller-code"
	usageTigoPushBillerCode        = "Tigo push biller code"
	flagTigoPushTokenURL           = "tigo-push-token-url" //nolint:gosec
	usageTigoPushTokenURL          = "Tigo push token URL" //nolint:gosec
	flagTigoPushPayURL             = "tigo-push-pay-url"
	usageTigoPushPayURL            = "Tigo push pay URL"
	flagTigoPasswordGrantType      = "tigo-password-grant-type"
	usageTigoPasswordGrantType     = "Tigo password grant type"
)

func SetTigoPesa(cmd *cobra.Command) {
	t := config.DefaultTigoPesaConfig()
	s := cmd.PersistentFlags().String
	s(flagTigoDisburseAccountName, t.DisburseAccountName, usageTigoDisburseAccountName)
	s(flagTigoDisburseAccountMSISDN, t.DisburseAccountMSISDN, usageTigoDisburseAccountMSISDN)
	s(flagTigoDisburseBrandID, t.DisburseBrandID, usageTigoDisburseBrandID)
	s(flagTigoDisbursePIN, t.DisbursePIN, usageTigoDisbursePIN)
	s(flagTigoDisburseURL, t.DisburseRequestURL, usageTigoDisburseURL)
	s(flagTigoPushBillerCode, t.PushBillerCode, usageTigoPushBillerCode)
	s(flagTigoPushBillerMSISDN, t.PushBillerMSISDN, usageTigoPushBillerMSISDN)
	s(flagTigoPushBaseURL, t.PushBaseURL, usageTigoPushBaseURL)
	s(flagTigoPushPassword, t.PushPassword, usageTigoPushPassword)
	s(flagTigoPushUsername, t.PushUsername, usageTigoPushUsername)
	s(flagTigoPushTokenURL, t.PushTokenEndpoint, usageTigoPushTokenURL)
	s(flagTigoPushPayURL, t.PushPayEndpoint, usageTigoPushPayURL)
	s(flagTigoPasswordGrantType, t.PasswordGrantType, usageTigoPasswordGrantType)
}

func GetTigoPesaConfig(cmd *cobra.Command) (*tigopesa.Config, error) {
	str := cmd.PersistentFlags().GetString
	disburseAccountName, err := str(flagTigoDisburseAccountName)
	if err != nil {
		return nil, err
	}
	disburseAccountMSISDN, err := str(flagTigoDisburseAccountMSISDN)
	if err != nil {
		return nil, err
	}
	disburseBrandID, err := str(flagTigoDisburseBrandID)
	if err != nil {
		return nil, err
	}
	disbursePIN, err := str(flagTigoDisbursePIN)
	if err != nil {
		return nil, err
	}
	disburseRequestURL, err := str(flagTigoDisburseURL)
	if err != nil {
		return nil, err
	}
	pushUsername, err := str(flagTigoPushUsername)
	if err != nil {
		return nil, err
	}
	pushPassword, err := str(flagTigoPushPassword)
	if err != nil {
		return nil, err
	}
	pushPasswordGrantType, err := str(flagTigoPasswordGrantType)
	if err != nil {
		return nil, err
	}
	pushBillerCode, err := str(flagTigoPushBillerCode)
	if err != nil {
		return nil, err
	}
	pushBillerMSISDN, err := str(flagTigoPushBillerMSISDN)
	if err != nil {
		return nil, err
	}
	pushBaseURl, err := str(flagTigoPushBaseURL)
	if err != nil {
		return nil, err
	}
	pushTokenEndpoint, err := str(flagTigoPushTokenURL)
	if err != nil {
		return nil, err
	}

	pushPayEndpoint, err := str(flagTigoPushPayURL)
	if err != nil {
		return nil, err
	}

	t := &config.TigoPesa{
		DisburseAccountName:   disburseAccountName,
		DisburseAccountMSISDN: disburseAccountMSISDN,
		DisburseBrandID:       disburseBrandID,
		DisbursePIN:           disbursePIN,
		DisburseRequestURL:    disburseRequestURL,
		PushUsername:          pushUsername,
		PushPassword:          pushPassword,
		PasswordGrantType:     pushPasswordGrantType,
		PushBaseURL:           pushBaseURl,
		PushTokenEndpoint:     pushTokenEndpoint,
		PushBillerMSISDN:      pushBillerMSISDN,
		PushBillerCode:        pushBillerCode,
		PushPayEndpoint:       pushPayEndpoint,
	}

	return t.Export(), nil
}
