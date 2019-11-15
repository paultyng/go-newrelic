package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestQueryAlertPluginsConditions(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "plugins_conditions": [
			    {
			      "id": 12345,
			      "name": "Plugins Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
			      "entities": [
			        "987654321"
			      ],
			      "metric_description": "Requests",
			      "metric": "Component/Requests[Requests/Minute]",
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "average",
			      "plugin": {
			        "query": "11223",
			        "guid": "com.railsware.haproxy"
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

	pluginsAlertConditions, err := c.queryAlertPluginsConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("queryAlertPluginsConditions error")
	}

	if len(pluginsAlertConditions) == 0 {
		t.Fatal("No Plugins Alert Conditions found")
	}
}

func TestGetAlertPluginsCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "plugins_conditions": [
			    {
			      "id": 12345,
			      "name": "Plugins Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
			      "entities": [
			        "987654321"
			      ],
			      "metric_description": "Requests",
			      "metric": "Component/Requests[Requests/Minute]",
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "average",
			      "plugin": {
			        "query": "11223",
			        "guid": "com.railsware.haproxy"
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

	pluginsAlertCondition, err := c.GetAlertPluginsCondition(policyID, conditionID)
	if err != nil {
		t.Log(err)
		t.Fatal("GetAlertPluginsCondition error")
	}
	if pluginsAlertCondition == nil {
		t.Log(err)
		t.Fatal("GetAlertPluginsCondition error")
	}
}

func TestListAlertPluginsConditions(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "plugins_conditions": [
			    {
			      "id": 12345,
			      "name": "Plugins Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
			      "entities": [
			        "987654321"
			      ],
			      "metric_description": "Requests",
			      "metric": "Component/Requests[Requests/Minute]",
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "average",
			      "plugin": {
			        "query": "11223",
			        "guid": "com.railsware.haproxy"
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

	pluginsAlertConditions, err := c.ListAlertPluginsConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("ListAlertPluginsConditions error")
	}
	if len(pluginsAlertConditions) == 0 {
		t.Log(err)
		t.Fatal("ListAlertNrqlConditions error")
	}
}

func TestCreateAlertPluginsCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "plugins_condition":
			    {
			      "id": 12345,
			      "name": "Plugins Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
			      "entities": [
			        "987654321"
			      ],
			      "metric_description": "Queued Sessions",
			      "metric": "Component/Sessions/Queued[Sessions]",
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "average",
			      "plugin": {
			        "query": "11223",
			        "guid": "com.railsware.haproxy"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	pluginsAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	pluginsAlert := AlertPlugin{
		ID:   "11223",
		GUID: "com.railsware.haproxy",
	}

	pluginsEntities := []string{"111222333"}

	pluginAlertCondition := AlertPluginsCondition{
		PolicyID:          123,
		Name:              "Test Condition",
		Enabled:           true,
		Entities:          pluginsEntities,
		MetricDescription: "Queued Sessions",
		Metric:            "Component/Sessions/Queued[Sessions]",
		RunbookURL:        "https://example.com/runbook.md",
		Terms:             pluginsAlertConditionTerms,
		ValueFunction:     "all",
		Plugin:            pluginsAlert,
	}

	pluginAlertConditionResp, err := c.CreateAlertPluginsCondition(pluginAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertPluginsCondition error")
	}
	if pluginAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertPluginsCondition error")
	}
	if pluginAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
	if pluginAlertConditionResp.Metric != "Component/Sessions/Queued[Sessions]" {
		t.Fatal("Metric was not parsed correctly")
	}
}

func TestUpdateAlertPluginsCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
			  "plugins_condition":
			    {
			      "id": 12345,
			      "name": "Plugins Condition",
			      "runbook_url": "https://example.com/runbook.md",
			      "enabled": true,
			      "entities": [
			        "987654321"
			      ],
			      "metric_description": "Queued Sessions",
			      "metric": "Component/Sessions/Queued[Sessions]",
			      "terms": [
			        {
			          "duration": "10",
			          "operator": "below",
			          "priority": "critical",
			          "threshold": "2",
			          "time_function": "all"
			         }
			      ],
			      "value_function": "average",
			      "plugin": {
			        "query": "11223",
			        "guid": "com.railsware.haproxy"
			      }
			    }
			}
			`))
		if err != nil {
			t.Log(err)
		}
	}))

	pluginsAlertConditionTerms := []AlertConditionTerm{
		{
			Duration:     10,
			Operator:     "below",
			Priority:     "critical",
			Threshold:    2.0,
			TimeFunction: "all",
		},
	}

	pluginsAlert := AlertPlugin{
		ID:   "11223",
		GUID: "com.railsware.haproxy",
	}

	pluginsEntities := []string{"111222333"}

	pluginAlertCondition := AlertPluginsCondition{
		PolicyID:          123,
		Name:              "Test Condition",
		Enabled:           true,
		Entities:          pluginsEntities,
		MetricDescription: "Queued Sessions",
		Metric:            "Component/Sessions/Queued[Sessions]",
		RunbookURL:        "https://example.com/runbook.md",
		Terms:             pluginsAlertConditionTerms,
		ValueFunction:     "all",
		Plugin:            pluginsAlert,
	}

	pluginAlertConditionResp, err := c.UpdateAlertPluginsCondition(pluginAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("UpdateAlertPluginsCondition error")
	}
	if pluginAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("UpdateAlertPluginsCondition error")
	}
	if pluginAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
	if pluginAlertConditionResp.Metric != "Component/Sessions/Queued[Sessions]" {
		t.Fatal("Metric was not parsed correctly")
	}
}

func TestDeleteAlertPluginsCondition(t *testing.T) {
	policyID := 123
	conditionID := 12345
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.Method != "DELETE" {
			t.Fatal("DeleteAlertPluginsCondition did not use DELETE method")
		}
		if r.URL.Path != fmt.Sprintf("/alerts_plugins_conditions/%v.json", conditionID) {
			t.Fatal("DeleteAlertPluginsCondition did not use the correct URL")
		}
	}))
	err := c.DeleteAlertPluginsCondition(policyID, conditionID)
	if err != nil {
		t.Log(err)
		t.Fatal("DeleteAlertPluginsCondition error")
	}
}
