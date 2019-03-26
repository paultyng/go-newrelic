package api

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
		if _, ok := c.resourceMap[(resourceID("resource", id))]; !ok {
			t.Log("Failed to lock resources")
			t.Fail()
		}
	}

}

func TestUnLockResources(t *testing.T) {
	c := New(Config{})
	ids := []int{123, 456}

	for _, id := range ids {
		c.resourceMap[resourceID("resource", id)] = make(chan struct{}, 1)
	}

	c.UnlockResources("resource", ids)

	for _, id := range ids {
		select {
		case <-c.resourceMap[resourceID("resource", id)]:
			continue
		default:
			t.Log("Failed to unlock resources")
			t.Fail()
		}
	}
}

func TestLockingOfResources(t *testing.T) {
	var c = New(Config{})
	var res123Locked bool
	var res456Locked bool
	var done = make(chan struct{})

	f := func(id int, t *testing.T) {
		c.LockResources("resource", []int{id})
		defer func() {
			// before this func f returns, set locked bool to false
			// and actually unlock the resources
			if id == 123 {
				res123Locked = false
			}
			if id == 456 {
				res456Locked = false
			}
			c.UnlockResources("resource", []int{id})
			done <- struct{}{}
		}()
		if id == 123 {
			if res123Locked {
				// if res123 is locked, fail because we shouldn't be able
				// to access it
				t.Log("Resource accessed while locked")
				t.Fail()
			}
			// now set the locked bool to true
			res123Locked = true
		}
		if id == 456 {
			if res456Locked {
				// if res456 is locked, fail because we shouldn't be able
				// to access it
				t.Log("Resource accessed while locked")
				t.Fail()
			}
			// now set the locked bool to true
			res456Locked = true
		}
		// wait a little bit to ensure the goroutines overlap
		time.Sleep(time.Second / 10)

	}

	for i := 0; i < 4; i++ {
		go f(123, t)
		go f(456, t)
	}

	// make sure all goroutines have finished before continuing
	for i := 0; i < 8; i++ {
		<-done
	}
}
