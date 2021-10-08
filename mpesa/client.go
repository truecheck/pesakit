package mpesa

import (
	"fmt"
	"github.com/techcraftlabs/pesakit/internal"
	"net/http"
	"strings"
)

var (
	_ internal.RequestInformer = (*RequestType)(nil)
	_ market                   = (*Market)(nil)
)

const (
	GhanaMarket    = Market(0)
	TanzaniaMarket = Market(1)
)

const (
	SessionID RequestType = iota
	PushPay
	Disburse
)

const (
	SANDBOX Platform = iota
	OPENAPI
)

type (
	market interface {
		URLContextValue() string
		Country() string
		Currency() string
		Description() string
	}
	Platform    int
	Market      int
	RequestType int
)

func MarketFmt(marketString string) Market {
	if strings.ToLower(marketString) == "ghana" {
		return GhanaMarket
	}

	if strings.ToLower(marketString) == "tanzania" {
		return TanzaniaMarket
	}

	return Market(-1)
}

func PlatformFmt(platformString string) Platform {
	if strings.ToLower(platformString) == "openapi" {
		return OPENAPI
	}

	if strings.ToLower(platformString) == "sandbox" {
		return SANDBOX
	}

	return Platform(-1)
}

func (p Platform) String() string {
	if p == OPENAPI {
		return "openapi"
	}

	return "sandbox"
}

func (m Market) URLContextValue() string {
	switch m {

	//ghana
	case 0:
		return "vodafoneGHA"
		//tanzania
	case 1:
		return "vodacomTZN"
	default:
		return ""
	}
}

func (m Market) Country() string {
	switch m {

	//ghana
	case 0:
		return "GHA"
		//tanzania
	case 1:
		return "TZN"
	default:
		return ""
	}
}

func (m Market) Currency() string {
	switch m {

	//ghana
	case 0:
		return "GHS"
		//tanzania
	case 1:
		return "TZS"
	default:
		return ""
	}
}

func (m Market) Description() string {
	switch m {

	//ghana
	case 0:
		return "Vodafone Ghana"
		//tanzania
	case 1:
		return "Vodacom Tanzania"
	default:
		return ""
	}
}

func (r RequestType) String() string {

	return fmt.Sprintf("mno=%s:: group=%s:: request=%s:", r.MNO(), r.Group(), r.Name())
}

func (r RequestType) Endpoint() string {
	switch r {

	case SessionID:
		return "/getSession/"

	case PushPay:
		return "/c2bPayment/singleStage/"

	case Disburse:
		return "/b2cPayment/"

	}
	return ""
}

func (r RequestType) Method() string {
	switch r {

	case SessionID:
		return http.MethodGet

	case PushPay:
		return http.MethodPost

	default:
		return http.MethodPost

	}
}

func (r RequestType) Name() string {
	return []string{"get session id", "ussd push",
		"disbursement"}[r]
}

func (r RequestType) MNO() string {
	return "vodacom"
}

func (r RequestType) Group() string {
	switch r {
	case SessionID:
		return "Authorization"

	case PushPay:
		return "collection"

	case Disburse:
		return "disbursement"

	default:
		return ""
	}
}
