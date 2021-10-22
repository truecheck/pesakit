//

package airtel

import (
	"fmt"
	"github.com/techcraftlabs/base"
	"net/http"
	"strings"
)

var _ base.RequestInformer = (*RequestType)(nil)

const (
	defaultGrantType     = "client_credentials"
	AuthApiGroup         = "authorization"
	CollectionApiGroup   = "collection"
	DisbursementApiGroup = "disbursement"
	AccountApiGroup      = "account"
	KycApiGroup          = "kyc"
	TransactionApiGroup  = "transaction"
)

const (
	Authorization RequestType = iota
	UssdPush
	Refund
	PushEnquiry
	PushCallback
	Disbursement
	DisbursementEnquiry
	BalanceEnquiry
	TransactionSummary
	UserEnquiry
)

func (t RequestType) Method() string {
	switch t {
	case Authorization, UssdPush, Refund, PushCallback, Disbursement:
		return http.MethodPost

	case PushEnquiry, DisbursementEnquiry, UserEnquiry, BalanceEnquiry,
		TransactionSummary:
		return http.MethodGet

	default:
		return ""
	}
}

func (t RequestType) Name() string {
	return []string{"authorization", "ussd push", "refund", "push enquiry", "push callback",
		"disbursement", "disbursement enquiry", "balance enquiry", "transaction summary",
		"user enquiry"}[t]
}

func (t RequestType) Group() string {
	switch t {
	case Authorization:
		return AuthApiGroup

	case PushCallback, Refund, PushEnquiry, UssdPush:
		return CollectionApiGroup

	case Disbursement, DisbursementEnquiry:
		return DisbursementApiGroup

	case BalanceEnquiry:
		return AccountApiGroup

	case UserEnquiry:
		return KycApiGroup

	case TransactionSummary:
		return TransactionApiGroup

	default:
		return "unknown/unsupported api group"
	}
}

func (t RequestType) String() string {
	return fmt.Sprintf("mno=%s:: group=%s:: request=%s:", t.MNO(), t.Group(), t.Name())
}

func (t RequestType) Endpoint() string {
	panic("implement me")
}

func (t RequestType) MNO() string {
	return "airtel"
}

func (t RequestType) endpoint(es Endpoints) string {
	switch t {
	case Authorization:
		return es.AuthEndpoint

	case UssdPush:
		return es.PushEndpoint

	case PushEnquiry:
		return es.PushEnquiryEndpoint

	case Refund:
		return es.RefundEndpoint

	case Disbursement:
		return es.DisbursementEndpoint

	case DisbursementEnquiry:
		return es.DisbursementEnquiryEndpoint

	case UserEnquiry:
		return es.UserEnquiryEndpoint

	case BalanceEnquiry:
		return es.BalanceEnquiryEndpoint

	case TransactionSummary:
		return es.TransactionSummaryEndpoint

	default:
		return ""
	}
}

type (
	RequestType uint
	Endpoints   struct {
		AuthEndpoint                string
		PushEndpoint                string
		RefundEndpoint              string
		PushEnquiryEndpoint         string
		DisbursementEndpoint        string
		DisbursementEnquiryEndpoint string
		TransactionSummaryEndpoint  string
		BalanceEnquiryEndpoint      string
		UserEnquiryEndpoint         string
	}
)

func (c *Client) makeInternalRequest(requestType RequestType, payload interface{}, opts ...base.RequestOption) *base.Request {
	baseURL := c.baseURL
	endpoints := c.Conf.Endpoints
	edps := *endpoints
	url := appendEndpoint(baseURL, requestType.endpoint(edps))
	method := requestType.Method()
	return base.NewRequest(method, url, payload, opts...)
}

func appendEndpoint(url string, endpoint string) string {
	url, endpoint = strings.TrimSpace(url), strings.TrimSpace(endpoint)
	urlHasSuffix, endpointHasPrefix := strings.HasSuffix(url, "/"), strings.HasPrefix(endpoint, "/")

	bothTrue := urlHasSuffix == true && endpointHasPrefix == true
	bothFalse := urlHasSuffix == false && endpointHasPrefix == false
	notEqual := urlHasSuffix != endpointHasPrefix

	if notEqual {
		return fmt.Sprintf("%s%s", url, endpoint)
	}

	if bothFalse {
		return fmt.Sprintf("%s/%s", url, endpoint)
	}

	if bothTrue {
		endp := strings.TrimPrefix(endpoint, "/")
		return fmt.Sprintf("%s%s", url, endp)
	}

	return ""
}
