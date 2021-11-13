package mno

import (
	"fmt"
	"github.com/techcraftlabs/mna"
)

const (
	Airtel Operator = iota
	Tigo
	Vodacom
)

type Operator int

func FromString(s string) Operator {
	switch s {
	case "Airtel", "airtel", "airtel money", "AIRTEL", "AIRTEL MONEY":
		return Airtel
	case "Tigo", "TIGO", "tigopesa", "TIGO PESA", "tigo pesa":
		return Tigo
	case "Vodacom", "VODA", "MPESA", "M-PESA", "VODACOM", "VODACOM M-PESA", "mpesa", "m-pesa":
		return Vodacom
	default:
		panic(fmt.Sprintf("Unknown operator: %s", s))
	}
}

// Get takes a phone number figure out the operator and then formats the
// number if the operator is in the list.
func Get(phone string) (Operator, string, error) {

	data, err := mna.Details(phone)
	if err != nil {
		return -1, "", fmt.Errorf("could not figure the mno: %w", err)
	}

	isAirtel := data.CommonName == mna.Airtel
	isTigo := data.CommonName == mna.Tigo
	isVodacom := data.CommonName == mna.Vodacom

	fmtNumber, err := mna.Format(phone)
	if err != nil {
		return -1, "", fmt.Errorf("could not format the number: %w", err)
	}

	if isAirtel {
		return Airtel, fmtNumber, nil
	}

	if isTigo {
		return Tigo, fmtNumber, nil
	}

	if isVodacom {
		return Vodacom, fmtNumber, nil
	}

	return -1, "", fmt.Errorf("unsupported provider")
}
