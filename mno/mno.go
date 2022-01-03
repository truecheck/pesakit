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

package mno

import (
	"github.com/techcraftlabs/mna"
	"strings"
)

const (
	Tigo    Mno = tigo
	Airtel  Mno = airtel
	Vodacom Mno = vodacom
	Unknown Mno = unknown
	tigo        = "tigo"
	airtel      = "airtel"
	vodacom     = "vodacom"
	unknown     = "unknown"
)

type (
	Mno string
)

func PhoneNumberInfo(phone string) (Mno, string, error) {
	info, err := mna.Information(phone)
	if err != nil {
		return unknown, phone, err
	}

	operator := info.Operator
	formattedNumber := info.FormattedNumber

	commonName := operator.CommonName()
	if strings.EqualFold(commonName, tigo) {
		return Tigo, formattedNumber, nil
	}
	if strings.EqualFold(commonName, airtel) {
		return Airtel, formattedNumber, nil
	}
	if strings.EqualFold(commonName, vodacom) {
		return Vodacom, formattedNumber, nil
	}

	return Unknown, formattedNumber, nil
}

func Which(mnoStr string) Mno {
	switch strings.ToLower(mnoStr) {
	case tigo:
		return Tigo
	case airtel:
		return Airtel
	case vodacom:
		return Vodacom
	default:
		return Unknown
	}
}
