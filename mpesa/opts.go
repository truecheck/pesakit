package mpesa

import (
	"io"
	"net/http"
)

// ClientOption is a setter func to set DisburseClient details like
// Timeout, context, base and Logger
type ClientOption func(client *Client)

func WithCallbackHandler(handler PushCallbackHandler) ClientOption {
	return func(client *Client) {
		client.pushCallbackFunc = handler
	}
}

// WithApiPlatform .....
func WithApiPlatform(platform Platform) ClientOption {
	return func(client *Client) {
		client.Conf.Platform = platform

	}
}

// WithMarket .....
func WithMarket(market Market) ClientOption {
	return func(client *Client) {
		client.Conf.Market = market

	}
}

// WithDebugMode set debug mode to true or false
func WithDebugMode(debugMode bool) ClientOption {
	return func(client *Client) {
		client.base.DebugMode = debugMode

	}
}

// WithLogger set a Logger of user preference but of type io.Writer
// that will be used for debugging use cases. A default value is os.Stderr
// it can be replaced by any io.Writer unless its nil which in that case
// it will be ignored
func WithLogger(out io.Writer) ClientOption {
	return func(client *Client) {
		if out == nil {
			return
		}
		client.base.Logger = out
	}
}

// WithHTTPClient when called unset the present http.Client and replace it
// with c. In case user tries to pass a nil value referencing the pkg
// i.e WithHTTPClient(nil), it will be ignored and the pkg wont be replaced
// Note: the new pkg Transport will be modified. It will be wrapped by another
// middleware that enables pkg to
func WithHTTPClient(httpClient *http.Client) ClientOption {

	// TODO check if its really necessary to set the default Timeout to 1 minute

	return func(client *Client) {
		if httpClient == nil {
			return
		}

		client.base.Http = httpClient
	}
}
