package mpesa

import (
	"fmt"
	"github.com/pesakit/pesakit/internal"
	"strings"
)

func (eps *Endpoints) Get(requestType RequestType) string {
	switch requestType {
	case SessionID:
		return eps.AuthEndpoint

	case PushPay:
		return eps.PushEndpoint

	case Disburse:
		return eps.DisburseEndpoint
	}
	return ""
}

func (c *Client) makeInternalRequest(requestType RequestType, payload interface{}, opts ...internal.RequestOption) *internal.Request {
	baseURL := c.Conf.BasePath
	endpoints := c.Conf.Endpoints
	edps := endpoints
	url := appendEndpoint(baseURL, edps.Get(requestType))
	method := requestType.Method()
	return internal.NewRequest(method, url, payload, opts...)
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
