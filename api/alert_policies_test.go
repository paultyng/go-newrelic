package api

import (
	"net/http"
	"testing"
)

func TestUpdateAlertPolicy(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "policy": {
			    "id": 12345,
			    "incident_preference": "PER_POLICY",
			    "name": "New Name",
			    "created_at": 12345678900000,
			    "updated_at": 12345678900000
			  }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	policy := AlertPolicy{
		ID:                 123,
		IncidentPreference: "PER_CONDITION",
		Name:               "Old Name",
	}

	policyResp, err := c.UpdateAlertPolicy(policy)
	if err != nil {
		t.Log(err)
		t.Fatal("UpdateAlertPolicy error")
	}
	if policyResp == nil {
		t.Log(err)
		t.Fatal("UpdateAlertPolicy error")
	}
	if policyResp.Name != "New Name" {
		t.Fatal("Failed to change policy name")
	}
	if policyResp.IncidentPreference != "PER_POLICY" {
		t.Fatal("Failed to change incident preference")
	}

}
