package airtel

import (
	"io"
	"net/http"
	"strings"
)

// ClientOption is a setter func to set DisburseClient details like
// Timeout, context, base and Logger
type ClientOption func(client *Client)

func WithCallbackHandler(handler PushCallbackHandler) ClientOption {
	return func(client *Client) {
		client.pushCallbackFunc = handler
	}
}

// WithDebugMode set debug mode to true or false
func WithDebugMode(debugMode bool) ClientOption {
	return func(client *Client) {
		client.base.DebugMode = debugMode

	}
}

// WithEnvironment ....
func WithEnvironment(envString string) ClientOption {

	var (
		env Environment
	)
	isStaging := strings.ToLower(envString) == "staging"
	isProduction := strings.ToLower(envString) == "production"

	if isProduction {
		env = PRODUCTION
	} else if isStaging {
		env = STAGING
	} else {
		env = STAGING
	}
	return func(client *Client) {
		client.environment = env
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
