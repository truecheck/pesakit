//

package airtel

import (
	"github.com/techcraftlabs/pesakit/internal"
	"time"
)

const (
	PRODUCTION        Environment = "production"
	STAGING           Environment = "staging"
	BaseURLProduction             = "https://openapi.airtel.africa"
	BaseURLStaging                = "https://openapiuat.airtel.africa"
)

type (
	Environment string

	Config struct {
		Endpoints          *Endpoints
		AllowedCountries   map[string][]string
		DisbursePIN        string
		CallbackPrivateKey string
		CallbackAuth       bool
		PublicKey          string
		ClientID           string
		Secret             string
	}

	Client struct {
		baseURL          string
		environment      Environment
		Conf             *Config
		base             *internal.BaseClient
		token            *string
		tokenExpiresAt   time.Time
		pushCallbackFunc PushCallbackHandler
		reqAdapter       RequestAdapter
		resAdapter       ResponseAdapter
	}

	PushCallbackHandler interface {
		Handle(request CallbackRequest) error
	}
	PushCallbackFunc func(request CallbackRequest) error
)

func (pf PushCallbackFunc) Handle(request CallbackRequest) error {
	return pf(request)
}

func (config *Config) SetAllowedCountries(apiName string, countries []string) {
	if config.AllowedCountries == nil {
		m := make(map[string][]string)
		config.AllowedCountries = m
	}

	config.AllowedCountries[apiName] = countries
}

func (c *Client) SetRequestAdapter(adapter RequestAdapter) {
	c.reqAdapter = adapter
}

func (c *Client) SetResponseAdapter(adapter ResponseAdapter) {
	c.resAdapter = adapter
}

func (c *Client) SetPushCallbackHandler(handler PushCallbackHandler) {
	c.pushCallbackFunc = handler
}

func NewClient(config *Config, opts ...ClientOption) *Client {
	client := new(Client)
	if config.AllowedCountries == nil {
		m := make(map[string][]string)
		config.AllowedCountries = m
		config.SetAllowedCountries(CollectionApiGroup, []string{"Tanzania"})
		config.SetAllowedCountries(DisbursementApiGroup, []string{"Tanzania"})
		config.SetAllowedCountries(AccountApiGroup, []string{"Tanzania"})
		config.SetAllowedCountries(KycApiGroup, []string{"Tanzania"})
		config.SetAllowedCountries(TransactionApiGroup, []string{"Tanzania"})

	}
	token := new(string)
	base := internal.NewBaseClient()

	client = &Client{
		environment:    STAGING,
		Conf:           config,
		base:           base,
		token:          token,
		tokenExpiresAt: time.Now(),
		resAdapter:     &adapter{},
		reqAdapter:     &adapter{Conf: config},
	}

	for _, opt := range opts {
		opt(client)
	}
	env := client.environment

	switch env {
	case PRODUCTION:
		client.baseURL = BaseURLProduction
	case STAGING:
		client.baseURL = BaseURLStaging
	}

	return client
}
