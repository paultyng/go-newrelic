package api

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestLockResources(t *testing.T) {
	c := New(Config{})
	ids := []int{123, 456}
	c.LockResources("resource", ids)

	for _, id := range ids {
		if _, ok := c.resourceMap.Load(resourceID("resource", id)); !ok {
			t.Log("Failed to lock resources")
			t.Fail()
		}
	}
}

func TestUnLockResources(t *testing.T) {
	c := New(Config{})
	ids := []int{123, 456}

	for _, id := range ids {
		c.resourceMap.Store(resourceID("resource", id), struct{}{})
	}

	c.UnlockResources("resource", ids)

	for _, id := range ids {
		if _, ok := c.resourceMap.Load(resourceID("resource", id)); ok {
			t.Log("Failed to unlock resources")
			t.Fail()
		}
	}
}
