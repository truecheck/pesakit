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
