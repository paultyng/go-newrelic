package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestQueryAlertNrqlConditions(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "nrql_conditions": [
			    {
			      "id": 12345,
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
                  "violation_time_limit_seconds": 3600,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			  ]
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	policyID := 123

	nrqlAlertConditions, err := c.queryAlertNrqlConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("queryAlertNrqlConditions error")
	}

	if len(nrqlAlertConditions) == 0 {
		t.Fatal("No NRQL Alert Conditions found")
	}
}

func TestGetAlertNrqlCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "nrql_conditions": [
			    {
			      "id": 12345,
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
                  "violation_time_limit_seconds": 3600,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			  ]
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	policyID := 123
	conditionID := 12345

	nrqlAlertCondition, err := c.GetAlertNrqlCondition(policyID, conditionID)
	if err != nil {
		t.Log(err)
		t.Fatal("GetAlertNrqlCondition error")
	}
	if nrqlAlertCondition == nil {
		t.Log(err)
		t.Fatal("GetAlertNrqlCondition error")
	}
}

func TestListAlertNrqlConditions(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "nrql_conditions": [
			    {
			      "id": 12345,
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
                  "violation_time_limit_seconds": 3600,
			      "enabled": true,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			  ]
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	policyID := 123

	nrqlAlertConditions, err := c.ListAlertNrqlConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("ListAlertNrqlConditions error")
	}
	if len(nrqlAlertConditions) == 0 {
		t.Log(err)
		t.Fatal("ListAlertNrqlConditions error")
	}
}

func TestCreateAlertNrqlCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "nrql_condition":
			    {
			      "id": 12345,
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
                  "violation_time_limit_seconds": 3600,
			      "enabled": true,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	nrqlAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	nrqlAlertQuery := AlertNrqlQuery{
		Query:      "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
		SinceValue: "5",
	}

	nrqlAlertCondition := AlertNrqlCondition{
		PolicyID:      123,
		Name:          "Test Condition",
		Enabled:       true,
		RunbookURL:    "https://example.com/runbook.md",
		Terms:         nrqlAlertConditionTerms,
		ValueFunction: "all",
		Nrql:          nrqlAlertQuery,
	}

	nrqlAlertConditionResp, err := c.CreateAlertNrqlCondition(nrqlAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
}

func TestCreateAlertNrqlStaticCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		condition, err := extractAlertNrqlConditionFromRequest(r)
		if err != nil {
			t.Fatalf("Failed to parse request: %s", err)
		}

		if condition.Type != "static" {
			t.Fatalf("Type different from expected value: expected 'static', got '%s'", condition.Type)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`
			{
			  "nrql_condition":
			    {
			      "id": 12345,
                  "type": "static",
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
                  "violation_time_limit_seconds": 3600,
			      "enabled": true,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	nrqlAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	nrqlAlertQuery := AlertNrqlQuery{
		Query:      "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
		SinceValue: "5",
	}

	nrqlAlertCondition := AlertNrqlCondition{
		PolicyID:      123,
		Name:          "Test Condition",
		Type:          "static",
		Enabled:       true,
		RunbookURL:    "https://example.com/runbook.md",
		Terms:         nrqlAlertConditionTerms,
		ValueFunction: "all",
		Nrql:          nrqlAlertQuery,
	}

	nrqlAlertConditionResp, err := c.CreateAlertNrqlCondition(nrqlAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
	if nrqlAlertConditionResp.Type != "static" {
		t.Fatal("Type was not parsed correctly")
	}
}

func TestCreateAlertNrqlBaselineCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		condition, err := extractAlertNrqlConditionFromRequest(r)
		if err != nil {
			t.Fatalf("Failed to parse request: %s", err)
		}

		if condition.Type != "baseline" {
			t.Fatalf("Type different from expected value: expected 'baseline', got '%s'", condition.Type)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`
			{
			  "nrql_condition":
			    {
			      "id": 12345,
                  "type": "baseline",
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
                  "violation_time_limit_seconds": 3600,
			      "enabled": true,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	nrqlAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	nrqlAlertQuery := AlertNrqlQuery{
		Query:      "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
		SinceValue: "5",
	}

	nrqlAlertCondition := AlertNrqlCondition{
		PolicyID:      123,
		Type:          "baseline",
		Name:          "Test Condition",
		Enabled:       true,
		RunbookURL:    "https://example.com/runbook.md",
		Terms:         nrqlAlertConditionTerms,
		ValueFunction: "all",
		Nrql:          nrqlAlertQuery,
	}

	nrqlAlertConditionResp, err := c.CreateAlertNrqlCondition(nrqlAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
	if nrqlAlertConditionResp.Type != "baseline" {
		t.Fatal("Type was not parsed correctly")
	}
}

func TestCreateAlertNrqlOutlierCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		condition, err := extractAlertNrqlConditionFromRequest(r)
		if err != nil {
			t.Fatalf("Failed to parse request: %s", err)
		}

		if condition.Type != "outlier" {
			t.Fatalf("Type different from expected value: expected 'outlier', got '%s'", condition.Type)
		}
		if condition.ExpectedGroups != 2 {
			t.Fatalf("ExpectedGroups different from expected value: expected '2', got '%d'", condition.ExpectedGroups)
		}
		if condition.IgnoreOverlap != false {
			t.Fatalf("IgnoreOverlap different from expected value: expected 'false', got '%t'", condition.IgnoreOverlap)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(`
			{
			  "nrql_condition":
			    {
			      "id": 12345,
                  "type": "outlier",
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
                  "violation_time_limit_seconds": 3600,
			      "enabled": true,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
                  "expected_groups": 2,
                  "ignore_overlap": false,
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	nrqlAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	nrqlAlertQuery := AlertNrqlQuery{
		Query:      "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
		SinceValue: "5",
	}

	nrqlAlertCondition := AlertNrqlCondition{
		PolicyID:       123,
		Type:           "outlier",
		Name:           "Test Condition",
		Enabled:        true,
		RunbookURL:     "https://example.com/runbook.md",
		Terms:          nrqlAlertConditionTerms,
		ValueFunction:  "all",
		ExpectedGroups: 2,
		IgnoreOverlap:  false,
		Nrql:           nrqlAlertQuery,
	}

	nrqlAlertConditionResp, err := c.CreateAlertNrqlCondition(nrqlAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
	if nrqlAlertConditionResp.Type != "outlier" {
		t.Fatal("Type was not parsed correctly")
	}
	if nrqlAlertConditionResp.ExpectedGroups != 2 {
		t.Fatal("Expected Groups was not parsed correctly")
	}
	if nrqlAlertConditionResp.IgnoreOverlap != false {
		t.Fatal("Ignore Overlap was not parsed correctly")
	}
}

func TestUpdateAlertNrqlCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "nrql_condition":
			    {
			      "id": 12345,
			      "name": "NRQL Condition",
			      "runbook_url": "https://example.com/runbook.md",
                  "violation_time_limit_seconds": 3600,
			      "enabled": true,
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "single_value",
			      "nrql": {
			        "query": "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
			        "since_value": "5"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	nrqlAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	nrqlAlertQuery := AlertNrqlQuery{
		Query:      "SELECT uniqueCount(fieldname) FROM indexname WHERE fieldname2 = 'somevaluetofilterby'",
		SinceValue: "5",
	}

	nrqlAlertCondition := AlertNrqlCondition{
		PolicyID:      123,
		Name:          "Test Condition",
		Enabled:       true,
		RunbookURL:    "https://example.com/runbook.md",
		Terms:         nrqlAlertConditionTerms,
		ValueFunction: "all",
		Nrql:          nrqlAlertQuery,
	}

	nrqlAlertConditionResp, err := c.UpdateAlertNrqlCondition(nrqlAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("UpdateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("UpdateAlertNrqlCondition error")
	}
	if nrqlAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
}

func TestDeleteAlertNrqlCondition(t *testing.T) {
	policyID := 123
	conditionID := 12345
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.Method != "DELETE" {
			t.Fatal("DeleteAlertNrqlCondition did not use DELETE method")
		}
		if r.URL.Path != fmt.Sprintf("/alerts_nrql_conditions/%v.json", conditionID) {
			t.Fatal("DeleteAlertNrqlCondtion did not use the correct URL")
		}
	}))
	err := c.DeleteAlertNrqlCondition(policyID, conditionID)
	if err != nil {
		t.Log(err)
		t.Fatal("DeleteAlertNrqlCondition error")
	}
}

func extractAlertNrqlConditionFromRequest(r *http.Request) (*AlertNrqlCondition, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var query *struct {
		Condition AlertNrqlCondition `json:"nrql_condition"`
	}
	err = json.Unmarshal(body, &query)
	if err != nil {
		return nil, err
	}

	return &query.Condition, nil
}
