package tigo

import (
	"fmt"
	"github.com/techcraftlabs/pesakit/internal"
	"net/http"
)

var _ internal.RequestInformer = (*RequestType)(nil)

const (
	GetToken RequestType = iota
	PushPay
	Disburse
	Callback
)

type RequestTypeInfo interface {
	Name() string
	MNO() string
	Group() string
	Endpoint() string
	fmt.Stringer
}

type RequestType int

func (r RequestType) Method() string {
	switch r {
	default:
		return http.MethodPost
	}
}

func (r RequestType) Name() string {
	names := []string{
		"get token", "pushpay", "disburse", "callbnack",
	}

	return names[r]
}

func (r RequestType) MNO() string {
	return "tigo"
}

func (r RequestType) Group() string {
	switch r {
	case Callback, PushPay:
		return "collection"

	case GetToken:
		return "authorization"

	case Disburse:
		return "disbursement"

	default:
		return ""
	}
}

func (r RequestType) Endpoint() string {
	panic("implement me")
}

func (r RequestType) String() string {
	return fmt.Sprintf("mno=%s:: group=%s:: request=%s:", r.MNO(), r.Group(), r.Name())
}
