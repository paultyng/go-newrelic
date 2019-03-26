package api

import (
	"crypto/tls"
	"fmt"

	"github.com/tomnomnom/linkheader"
	resty "gopkg.in/resty.v1"
)

// Client represents the client state for the API.
type Client struct {
	RestyClient *resty.Client

	// For locking resources to prevent concurrent mutations
	resourceMap map[string](chan struct{})
	// Indicates whether or not to sequentialize policy_channel updates
	seqPolicyChannelUpdates bool
}

// InfraClient represents the client state for the Infrastructure API
type InfraClient struct {
	Client
}

// NewInfraClient returns a new InfraClient for the specified apiKey.
func NewInfraClient(config Config) InfraClient {
	if config.BaseURL == "" {
		config.BaseURL = "https://infra-api.newrelic.com/v2"
	}

	return InfraClient{New(config)}
}

// ErrorResponse represents an error response from New Relic.
type ErrorResponse struct {
	Detail *ErrorDetail `json:"error,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e != nil && e.Detail != nil {
		return e.Detail.Title
	}
	return "Unknown error"
}

// ErrorDetail represents the details of an ErrorResponse from New Relic.
type ErrorDetail struct {
	Title string `json:"title,omitempty"`
}

// Config contains all the configuration data for the API Client plus
// a SequentializePolicyChannelUpdates option. When set to true, this will
// force the client to sequentialize policy_channel updates that include some
// overlapping subset of channel_ids.
type Config struct {
	APIKey    string
	BaseURL   string
	ProxyURL  string
	Debug     bool
	TLSConfig *tls.Config

	// Indicates whether to force sequential execution of policy_channel updates
	SequentializePolicyChannelUpdates bool
}

// New returns a new Client for the specified apiKey.
func New(config Config) Client {
	r := resty.New()

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.newrelic.com/v2"
	}

	proxyURL := config.ProxyURL
	if proxyURL != "" {
		r.SetProxy(proxyURL)
	}

	r.SetHeader("X-Api-Key", config.APIKey)
	r.SetHostURL(baseURL)

	if config.TLSConfig != nil {
		r.SetTLSClientConfig(config.TLSConfig)
	}
	if config.Debug {
		r.SetDebug(true)
	}

	c := Client{
		RestyClient: r,

		seqPolicyChannelUpdates: config.SequentializePolicyChannelUpdates,
		resourceMap:             make(map[string]chan struct{}),
	}

	return c
}

// Do executes an API request with the specified parameters.
func (c *Client) Do(method string, path string, body interface{}, response interface{}) (string, error) {
	r := c.RestyClient.R().
		SetError(&ErrorResponse{}).
		SetHeader("Content-Type", "application/json")

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	apiResponse, err := r.Execute(method, path)

	if err != nil {
		return "", err
	}

	nextPath := ""
	header := apiResponse.Header().Get("Link")
	if header != "" {
		links := linkheader.Parse(header)

		for _, link := range links.FilterByRel("next") {
			nextPath = link.URL
			break
		}
	}

	statusClass := apiResponse.StatusCode() / 100 % 10

	if statusClass == 2 {
		return nextPath, nil
	}

	rawError := apiResponse.Error()

	if rawError != nil {
		apiError := rawError.(*ErrorResponse)

		if apiError.Detail != nil {
			return "", apiError
		}
	}

	return "", fmt.Errorf("Unexpected status %v returned from API", apiResponse.StatusCode())
}

// LockResources "locks" a set of resources from being modified by
// blocking further operations until the matching channels send struct.
// If the resourceID has not been seen before, the map value is initialized
// to a struct channel of buffer 1.
func (c *Client) LockResources(resourceType string, ids []int) {
	for _, id := range ids {
		resID := resourceID(resourceType, id)
		if _, ok := c.resourceMap[resID]; ok {
			_ = <-c.resourceMap[resID]
		} else {
			c.resourceMap[resID] = make(chan struct{}, 1)
		}
	}
}

// UnlockResources "unlocks" a set of resources, opening them up for
// modification, by sending a struct on the matching channels.
// This will unblock any processes waiting for the signals. If the resource
// has not been previously locked, no operation is performed.
func (c *Client) UnlockResources(resourceType string, ids []int) {
	for _, id := range ids {
		resID := resourceID(resourceType, id)
		if _, ok := c.resourceMap[resID]; ok {
			c.resourceMap[resID] <- struct{}{}
		}
	}
}

// resourceID generates unique identifier for storing and accessing
// the resource in the resource map
func resourceID(resourceType string, id int) string {
	return fmt.Sprintf("%s-%d", resourceType, id)
}
