package mpesa

import (
	"context"
	"fmt"
	"github.com/pesakit/pesakit/internal"
	"github.com/pesakit/pesakit/pkg/crypto"
	"net/http"
	"time"
)

var (
	_ Service             = (*Client)(nil)
	_ PushCallbackHandler = (*PushCallbackFunc)(nil)
)

type (
	Service interface {
		SessionID(ctx context.Context) (response SessionResponse, err error)
		PushAsync(ctx context.Context, request Request) (PushAsyncResponse, error)
		Disburse(ctx context.Context, request Request) (DisburseResponse, error)
		CallbackServeHTTP(w http.ResponseWriter, r *http.Request)
	}

	PushCallbackHandler interface {
		Handle(request PushCallbackRequest) (PushCallbackResponse, error)
	}
	PushCallbackFunc func(request PushCallbackRequest) (PushCallbackResponse, error)

	// Config contains details initialize in mpesa portal
	// Applications require the following details:
	//•	Application Name – human-readable name of the application
	//•	Version – version number of the application, allowing changes in API products to be managed in different versions
	//•	Description – Free text field to describe the use of the application
	//•	APIKey – Unique authorisation key used to authenticate the application on the first call. API Keys need to be encrypted in the first “Generate Session API Call” to create a valid session key to be used as an access token for future calls. Encrypting the API Key is documented in the GENERATE SESSION API page
	//•	SessionLifetime – The session key has a finite lifetime of availability that can be configured. Once a session key has expired, the session is no longer usable, and the caller will need to authenticate again.
	//•	TrustedSources – the originating caller can be limited to specific IP address(es) as an additional security measure.
	//•	Products / Scope / Limits – the required API products for the application can be enabled and limits defined for each call.
	Config struct {
		Endpoints              *Endpoints
		Name                   string
		Version                string
		Description            string
		BasePath               string
		Market                 Market
		Platform               Platform
		APIKey                 string
		PublicKey              string
		SessionLifetimeMinutes int64
		ServiceProvideCode     string
		TrustedSources         []string
	}

	Endpoints struct {
		AuthEndpoint     string
		PushEndpoint     string
		DisburseEndpoint string
	}

	Client struct {
		Conf *Config
		base *internal.BaseClient
		//Market            Market
		//Platform          Platform
		encryptedApiKey   *string
		sessionID         *string
		sessionExpiration time.Time
		pushCallbackFunc  PushCallbackHandler
		requestAdapter    *requestAdapter
		rp                internal.Replier
		rv                internal.Receiver
	}
)

func NewClient(conf *Config, opts ...ClientOption) *Client {
	enc := new(string)
	ses := new(string)

	client := new(Client)

	basePath := conf.BasePath

	client = &Client{
		Conf: conf,
		base: internal.NewBaseClient(),
		//Market:            TanzaniaMarket,
		//Platform:          SANDBOX,
		encryptedApiKey:   enc,
		sessionID:         ses,
		sessionExpiration: time.Now(),
	}

	for _, opt := range opts {
		opt(client)
	}

	platform := client.Conf.Platform
	market := client.Conf.Market

	platformStr, marketStr := platform.String(), market.URLContextValue()
	p := fmt.Sprintf("https://%s/%s/ipg/v2/%s/", basePath, platformStr, marketStr)
	client.Conf.BasePath = p
	client.requestAdapter = &requestAdapter{
		platform:            platform,
		market:              market,
		serviceProviderCode: conf.ServiceProvideCode,
	}

	rp := internal.NewReplier(client.base.Logger, client.base.DebugMode)
	rv := internal.NewReceiver(client.base.Logger, client.base.DebugMode)
	client.rp = rp
	client.rv = rv
	return client
}

func (p PushCallbackFunc) Handle(request PushCallbackRequest) (PushCallbackResponse, error) {
	return p(request)
}

func (c *Client) SessionID(ctx context.Context) (response SessionResponse, err error) {

	token, err := c.getEncryptionKey()
	if err != nil {
		return response, err
	}
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	var opts []internal.RequestOption
	headersOpt := internal.WithRequestHeaders(headers)
	opts = append(opts, headersOpt)
	re := c.makeInternalRequest(SessionID, nil, opts...)
	res, err := c.base.Do(ctx, SessionID.String(), re, &response)
	if err != nil {
		return response, err
	}

	resErr := res.Error
	if resErr != nil {
		return SessionResponse{}, fmt.Errorf("could not fetch session id: %w", resErr)
	}

	//save the session id
	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not fetch session id: %s", response.OutputErr)
		return response, err1
	}

	sessLifeTimeMin := c.Conf.SessionLifetimeMinutes
	sessID := response.ID
	up := time.Duration(sessLifeTimeMin) * time.Minute
	expiration := time.Now().Add(up)
	c.sessionExpiration = expiration
	c.sessionID = &sessID

	return response, nil
}

func (c *Client) PushAsync(ctx context.Context, request Request) (response PushAsyncResponse, err error) {
	sess, err := c.checkSessionID()
	if err != nil {
		return response, err
	}
	token, err := crypto.EncryptKey(sess, c.Conf.PublicKey)
	if err != nil {
		return response, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	payload, err := c.requestAdapter.Adapt(PushPay, request)
	if err != nil {
		return PushAsyncResponse{}, err
	}

	var opts []internal.RequestOption
	headersOpt := internal.WithRequestHeaders(headers)
	opts = append(opts, headersOpt)
	re := c.makeInternalRequest(PushPay, payload, opts...)
	res, err := c.base.Do(ctx, PushPay.String(), re, &response)

	if err != nil {
		return response, err
	}
	fmt.Printf("pushpay response: %s: %v\n", PushPay.String(), res)

	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not perform c2b single stage request: %s", response.OutputErr)
		return response, err1
	}

	return response, nil
}

func (c *Client) Disburse(ctx context.Context, request Request) (response DisburseResponse, err error) {
	sess, err := c.checkSessionID()
	if err != nil {
		return response, err
	}
	token, err := crypto.EncryptKey(sess, c.Conf.PublicKey)
	if err != nil {
		return response, err
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Origin":        "*",
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}

	payload, err := c.requestAdapter.Adapt(Disburse, request)
	if err != nil {
		return DisburseResponse{}, err
	}

	var opts []internal.RequestOption
	headersOpt := internal.WithRequestHeaders(headers)
	opts = append(opts, headersOpt)
	re := c.makeInternalRequest(Disburse, payload, opts...)
	res, err := c.base.Do(ctx, Disburse.String(), re, &response)

	if err != nil {
		return response, err
	}
	fmt.Printf("disburse response: %s: %v\n", Disburse.String(), res)

	if response.OutputErr != "" {
		err1 := fmt.Errorf("could not perform disburse request: %s", response.OutputErr)
		return response, err1
	}

	return response, nil
}

func (c *Client) CallbackServeHTTP(writer http.ResponseWriter, request *http.Request) {
	body := new(PushCallbackRequest)
	_, err := c.rv.Receive(context.TODO(), "mpesa push callback", request, body)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	reqBody := *body

	resp, err := c.pushCallbackFunc.Handle(reqBody)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	response := internal.NewResponse(200, resp)
	c.rp.Reply(writer, response)
}
