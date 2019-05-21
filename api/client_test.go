package api

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientDoPaging(t *testing.T) {
	for i, c := range []struct {
		expectedNext string
		linkHeader   string
		body         string
	}{
		{"", "", ""},
		{"", "", "{}"},
		{"", `<https://api.github.com/user/58276/repos?page=2>; rel="last"`, "{}"},
		{"", "", `{"links":null}`},
		{"", "", `{"links":{}}`},
		{"", "", `{"links":{"last":"foo"}}`},

		{"https://api.github.com/user/58276/repos?page=2", `<https://api.github.com/user/58276/repos?page=2>; rel="next"`, "{}"},
		{"https://api.github.com/user/58276/repos?page=2", "", `{"links":{"next":"https://api.github.com/user/58276/repos?page=2"}}`},
		{"https://api.github.com/user/58276/repos?page=2", "", `{"links":{"next":"https://api.github.com/user/58276/repos?page=2"}}`},
		{"https://api.github.com/user/58276/repos?page=2", `<https://api.github.com/user/58276/repos?page=2>; rel="next"`, `{"links":{"next":"https://should-not-match"}}`},
	} {
		t.Run(fmt.Sprintf("%d %s", i, c.expectedNext), func(t *testing.T) {
			cli := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if c.linkHeader != "" {
					w.Header().Set("Link", c.linkHeader)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(c.body))
			}))
			actualNext, err := cli.Do("GET", "/path", nil, nil)
			if err != nil {
				t.Fatal(err)
			}
			if actualNext != c.expectedNext {
				t.Fatalf("expected %q but got %q", c.expectedNext, actualNext)
			}
		})
	}

}

func newTestAPIClient(handler http.Handler) *Client {
	ts := httptest.NewServer(handler)

	c := New(Config{
		APIKey:  "123456",
		BaseURL: ts.URL,
		Debug:   false,
	})

	return &c
}

func newTestAPIInfraClient(handler http.Handler) *InfraClient {
	ts := httptest.NewServer(handler)

	c := NewInfraClient(Config{
		APIKey:  "123456",
		BaseURL: ts.URL,
		Debug:   false,
	})

	return &c
}

func newTestAPIClientTLSConfig(handler http.Handler) *Client {
	ts := httptest.NewServer(handler)

	tlsCfg := &tls.Config{}
	tlsCfg.InsecureSkipVerify = true

	c := New(Config{
		APIKey:    "123456",
		BaseURL:   ts.URL,
		Debug:     false,
		TLSConfig: tlsCfg,
	})

	return &c
}
