package internal

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/pesakit/pesakit/internal/io"
	stdio "io"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

const (
	defaultTimeout = 60 * time.Second
)

type (
	BaseClient struct {
		Http      *http.Client
		Logger    stdio.Writer // for logging purposes
		DebugMode bool
		certPool  *x509.CertPool
	}

	ClientOption func(client *BaseClient)
)

func NewBaseClient(opts ...ClientOption) *BaseClient {
	defClient := &http.Client{
		Timeout: defaultTimeout,
	}
	client := &BaseClient{
		Http:      defClient,
		Logger:    io.Stderr,
		DebugMode: true,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (client *BaseClient) logPayload(t PayloadType, prefix string, payload interface{}) {
	buf, _ := MarshalPayload(t, payload)
	_, _ = client.Logger.Write([]byte(fmt.Sprintf("%s: %s\n\n", prefix, buf.String())))
}

func (client *BaseClient) log(name string, request *http.Request) {

	if request != nil {
		reqDump, _ := httputil.DumpRequest(request, true)
		_, err := fmt.Fprintf(client.Logger, "%s REQUEST: %s\n", name, reqDump)
		if err != nil {
			fmt.Printf("error while logging %s request: %v\n",
				strings.ToLower(name), err)
			return
		}
		return
	}
	return
}

// logOut is like log except this is for outgoing client requests:
// http.Request that is supposed to be sent to tigo
func (client *BaseClient) logOut(name string, request *http.Request, response *http.Response) {

	if request != nil {
		reqDump, _ := httputil.DumpRequestOut(request, true)
		_, err := fmt.Fprintf(client.Logger, "%s REQUEST\n%s\n", name, reqDump)
		if err != nil {
			fmt.Printf("error while logging %s request: %v\n",
				strings.ToLower(name), err)
		}
	}

	if response != nil {
		respDump, _ := httputil.DumpResponse(response, true)
		_, err := fmt.Fprintf(client.Logger, "%s RESPONSE\n%s\n", name, respDump)
		if err != nil {
			fmt.Printf("error while logging %s response: %v\n",
				strings.ToLower(name), err)
		}
	}

	return
}

// WithDebugMode set debug mode to true or false
func WithDebugMode(debugMode bool) ClientOption {
	return func(client *BaseClient) {
		client.DebugMode = debugMode

	}
}

// WithLogger set a Logger of user preference but of type io.Writer
// that will be used for debugging use cases. A default value is os.Stderr
// it can be replaced by any io.Writer unless its nil which in that case
// it will be ignored
func WithLogger(out stdio.Writer) ClientOption {
	return func(client *BaseClient) {
		if out == nil {
			return
		}
		client.Logger = out
	}
}

// WithHTTPClient when called unset the present http.Client and replace it
// with c. In case user tries to pass a nil value referencing the pkg
// i.e. WithHTTPClient(nil), it will be ignored and the pkg will not be replaced
// Note: the new pkg Transport will be modified. It will be wrapped by another
// middleware that enables pkg to
func WithHTTPClient(httpClient *http.Client) ClientOption {

	// TODO check if its really necessary to set the default Timeout to 1 minute

	return func(client *BaseClient) {
		if httpClient == nil {
			return
		}

		client.Http = httpClient
	}
}

func WithCACert(caCert []byte) ClientOption {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return func(client *BaseClient) {
		if caCert == nil {
			return
		}

		c := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
			CheckRedirect: client.Http.CheckRedirect,
			Jar:           client.Http.Jar,
			Timeout:       client.Http.Timeout,
		}

		client.Http = c
	}

}
