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
	"fmt"
	"github.com/pesakit/pesakit/mno"
	"github.com/spf13/cobra"
)

type TransactionRequest struct {
	ThirdPartyID string
	Reference    string
	Phone        string
	Mno          mno.Mno
	Description  string
	Amount       float64
}

const (
	flagRequestThirdPartyID  = "third-party-id"
	usageRequestThirdPartyID = "The third party ID"
	flagRequestReference     = "reference"
	usageRequestReference    = "The reference of the request"
	flagRequestDescription   = "description"
	usageRequestDescription  = "The description of the request"
	flagRequestAmount        = "amount"
	usageAmount              = "amount of money involved in the transaction"
	flagRequestPhone         = "phone"
	usageRequestPhone        = "phone number of the receiver"
)

func SetTransactionRequestFlags(cmd *cobra.Command) {
	strFlag := cmd.PersistentFlags().String
	floatFlag := cmd.PersistentFlags().Float64
	strFlag(flagRequestThirdPartyID, "", usageRequestThirdPartyID)
	strFlag(flagRequestReference, "", usageRequestReference)
	strFlag(flagRequestDescription, "", usageRequestDescription)
	floatFlag(flagRequestAmount, 0, usageAmount)
	strFlag(flagRequestPhone, "", usageRequestPhone)
}

func GetTransactionRequest(cmd *cobra.Command) (*TransactionRequest, error) {

	strFlag := cmd.PersistentFlags().GetString
	floatFlag := cmd.PersistentFlags().GetFloat64
	phone, err := strFlag(flagRequestPhone)
	if err != nil {
		return nil, err
	}
	mnoValue, phone, err := mno.PhoneNumberInfo(phone)
	if err != nil {
		return nil, err
	}
	if mnoValue == mno.Unknown {
		return nil, fmt.Errorf("unknown mno")
	}

	amount, err := floatFlag(flagRequestAmount)
	if err != nil {
		return nil, err
	}

	thirdID, err := strFlag(flagRequestThirdPartyID)
	if err != nil {
		return nil, err
	}
	desc, err := strFlag(flagRequestDescription)
	if err != nil {
		return nil, err
	}

	ref, err := strFlag(flagRequestReference)
	if err != nil {
		return nil, err
	}

	return &TransactionRequest{
		ThirdPartyID: thirdID,
		Reference:    ref,
		Phone:        phone,
		Mno:          mnoValue,
		Description:  desc,
		Amount:       amount,
	}, nil
}
