package api

import (
	"net/http"
	"net/http/httptest"
)

func newTestAPIClient(handler http.Handler) *Client {
	ts := httptest.NewServer(handler)

	c := New(Config{
		APIKey:  "123456",
		BaseURL: ts.URL,
		Debug:   false,
	})

	return &c
}
