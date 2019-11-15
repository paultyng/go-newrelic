package api

import (
	"net/http"
	"testing"
)

func TestListAlertConditions(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
				"conditions": [
					{
						"id": 1234,
						"type": "browser_metric",
						"name": "End User Apdex (Low)",
						"enabled": false,
						"entities": ["126408", "127809"],
						"metric": "end_user_apdex",
						"condition_scope": "application",
						"terms": [{
							"duration": "120",
							"operator": "below",
							"priority": "critical",
							"threshold": ".9",
							"time_function": "all"
						}]
					}
				]
			}
		`))
		if err != nil {
			t.Log(err)
		}
	}))

	policyID := 123

	alertConditions, err := c.queryAlertConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("GetAlertCondition error")
	}
	if alertConditions == nil {
		t.Log(err)
		t.Fatal("GetAlertCondition error")
	}
}
