package api

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

// Client represents the client state for the API.
type Client struct {
	RestyClient *resty.Client
}

type errorResponse struct {
	Error errorDetail
}

type errorDetail struct {
	Title string
}

// New returns a new Client for the specified apiKey.
func New(apiKey string) Client {
	r := resty.New()

	r.SetHeader("X-Api-Key", apiKey)
	r.SetHostURL("https://api.newrelic.com/v2")

	c := Client{
		RestyClient: r,
	}

	return c
}

// Debug sets the Client in to debug mode which outputs details about the HTTP traffic.
func (c *Client) Debug() {
	c.RestyClient.SetDebug(true)
}

// Do exectes an API request with the specified parameters.
func (c *Client) Do(method string, path string, qs map[string]string, body interface{}, response interface{}) error {
	r := c.RestyClient.R().
		SetError(errorResponse{})

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

	if apiError, ok := rawError.(*errorResponse); ok {
		return fmt.Errorf("Error returned from the API: %v", apiError.Error.Title)
	}

	if rawError != nil {
		return fmt.Errorf("Unexpected error: %v", rawError)
	}

	return fmt.Errorf("Unexpected status %v returned from API", apiResponse.StatusCode())
}
