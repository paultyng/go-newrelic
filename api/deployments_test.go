package api

import (
	"net/http"
	"testing"
)

func TestListDeployments(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
      {
        "deployments": [
          {
            "id": 1234567,
            "revision": "master",
            "changelog": null,
            "description": "Not specified",
            "user": "foo",
            "timestamp": "2000-01-01T01:00:00+00:00",
            "links": {
              "application": 123
            }
          }
        ]
      }
    `))
		if err != nil {
			t.Log(err)
		}
	}))

	appID := 123

	deployments, err := c.ListDeployments(appID)
	if err != nil {
		t.Log(err)
		t.Fatal("ListDeployments error")
	}

	if len(deployments) == 0 {
		t.Fatal("No deployments found")
	}
}

func TestCreateDeployment(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
      {
        "deployment": {
          "revision": "master",
          "changelog": "Foo Bar",
          "description": "Not specified",
          "user": "foo"
        }
      }
    `))
		if err != nil {
			t.Log(err)
		}
	}))

	appID := 123

	DeploymentStruct := Deployment{
		Revision:    "master",
		User:        "foo",
		Changelog:   "Foo Bar",
		Description: "12345678-1234-1234-1234-1234567890ab",
	}

	DeploymentResp, err := c.CreateDeployment(appID, DeploymentStruct)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateDeployment error")
	}
	if DeploymentResp == nil {
		t.Log(err)
		t.Fatal("CreateDeployment error")
	}
	if DeploymentResp.Revision != "master" {
		t.Log(DeploymentResp.Revision)
		t.Fatal("Revision was not parsed correctly")
	}
}

func TestDeleteDeployment(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
		{
			"deployment": {
				"id": 456,
				"revision": "master",
				"changelog": "Foo Bar",
				"description": "12345678-1234-1234-1234-1234567890ab",
				"user": "foo",
				"timestamp": "2019-10-31T15:25:38-07:00",
				"links": {
					"application": 123
				}
			},
			"links": {
				"deployment.agent": "/v2/applications/{application_id}"
			}
		}
    `))
		if err != nil {
			t.Log(err)
		}
	}))

	appID := 123
	depID := 456

	err := c.DeleteDeployment(appID, depID)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateDeployment error")
	}
}
