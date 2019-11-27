package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestQueryApplications_Basic(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
      {
        "applications": [
          {
            "id": 123,
            "name": "foo"
          },
          {
            "id": 456,
            "name": "bar"
          }
        ]
      }
    `))
	}))

	apps, err := c.QueryApplications(ApplicationsFilters{})
	if err != nil {
		t.Log(err)
		t.Fatal("queryApplications error")
	}

	if len(apps) == 0 {
		t.Fatal("No applications found")
	}
}

func TestDeleteApplication(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
      {
		"application": {
			"id": 123,
			"name": "test",
			"language": "go",
			"health_status": "gray",
			"reporting": false,
			"settings": {
				"app_apdex_threshold": 0.5,
				"end_user_apdex_threshold": 7,
				"enable_real_user_monitoring": true,
				"use_server_side_config": false
			},
			"links": {
				"application_instances": [],
				"servers": [],
				"application_hosts": []
			}
		}
      }
    `))
	}))

	err := c.DeleteApplication(123)
	if err != nil {
		fmt.Printf("Error: %s", err)
		t.Log(err)
		t.Fatal("DeleteApplication error")
	}
}
