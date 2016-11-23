package api

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

// Client represents the client state for the API.
type Client struct {
	RestyClient *resty.Client
}

type ErrorResponse struct {
	Detail ErrorDetail `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return e.Detail.Title
}

type ErrorDetail struct {
	Title string `json:"title"`
}

// Config contains all the configuration data for the API Client
type Config struct {
	APIKey  string
	BaseURL string
	Debug   bool
}

// New returns a new Client for the specified apiKey.
func New(config Config) Client {
	r := resty.New()

	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://api.newrelic.com/v2"
	}

	r.SetHeader("X-Api-Key", config.APIKey)
	r.SetHostURL(baseURL)

	if config.Debug {
		r.SetDebug(true)
	}

	c := Client{
		RestyClient: r,
	}

	return c
}

// Do exectes an API request with the specified parameters.
func (c *Client) Do(method string, path string, qs map[string]string, body interface{}, response interface{}) error {
	r := c.RestyClient.R().
		SetError(ErrorResponse{})

	if qs != nil {
		r = r.SetQueryParams(qs)
	}

	if body != nil {
		r = r.SetBody(body)
	}

	if response != nil {
		r = r.SetResult(response)
	}

	apiResponse, err := r.Execute(method, path)

	if err != nil {
		return err
	}

	statusClass := apiResponse.StatusCode() / 100 % 10

	if statusClass == 2 {
		return nil
	}

	rawError := apiResponse.Error()

	if apiError, ok := rawError.(*ErrorResponse); ok {
		return apiError
	}

	if rawError != nil {
		return fmt.Errorf("Unexpected error: %v", rawError)
	}

	return fmt.Errorf("Unexpected status %v returned from API", apiResponse.StatusCode())
}
