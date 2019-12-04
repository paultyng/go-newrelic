package api

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListAlertConditions(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
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
	}))

	terms := []AlertConditionTerm{
		{
			Duration:     120,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    0.9,
			TimeFunction: "all",
		},
	}

	expected := []AlertCondition{
		{
			ID:       1234,
			Type:     "browser_metric",
			Name:     "End User Apdex (Low)",
			Enabled:  false,
			Entities: []string{"126408", "127809"},
			Metric:   "end_user_apdex",
			Scope:    "application",
			Terms:    terms,
		},
	}

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
	if diff := cmp.Diff(alertConditions, expected); diff != "" {
		t.Fatalf("Alert conditions not parsed correctly: %s", diff)
	}
}
