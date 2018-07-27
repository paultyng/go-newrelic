package api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestQueryAlertInfraConditions(t *testing.T) {
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
			  "data": [
			    {
					"type": "infra_metric",
					"name": "High CPU",
					"enabled": true,
					"id": 1013339,
					"created_at_epoch_millis": 1521478734169,
					"updated_at_epoch_millis": 1521478734227,
					"policy_id": 210972,
					"event_type": "SystemSample",
					"select_value": "cpuPercent",
					"comparison": "above",
					"critical_threshold": {
						"value": 75,
						"duration_minutes": 2,
						"time_function": "all"
					}
				}
			  ]
			}
			`))
	}))

	policyID := 123

	infraAlertConditions, err := c.queryAlertInfraConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("queryAlertInfraConditions error")
	}

	if len(infraAlertConditions) == 0 {
		t.Fatal("No Infra Alert Conditions found")
	}
}

func TestGetAlertInfraCondition(t *testing.T) {
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"data": [
				  {
					  "type": "infra_metric",
					  "name": "High CPU",
					  "enabled": true,
					  "id": 12345,
					  "created_at_epoch_millis": 1521478734169,
					  "updated_at_epoch_millis": 1521478734227,
					  "policy_id": 210972,
					  "event_type": "SystemSample",
					  "select_value": "cpuPercent",
					  "comparison": "above",
					  "critical_threshold": {
						  "value": 75,
						  "duration_minutes": 2,
						  "time_function": "all"
					  }
				  }
				]
			}
			`))
	}))

	policyID := 123
	conditionID := 12345

	infraAlertCondition, err := c.GetAlertInfraCondition(policyID, conditionID)
	if err != nil {
		t.Log(err)
		t.Fatal("GetAlertInfraCondition error")
	}
	if infraAlertCondition == nil {
		t.Log(err)
		t.Fatal("GetAlertInfraCondition error")
	}
}

func TestListAlertInfraConditions(t *testing.T) {
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"data": [
				  {
					  "type": "infra_metric",
					  "name": "High CPU",
					  "enabled": true,
					  "id": 12345,
					  "created_at_epoch_millis": 1521478734169,
					  "updated_at_epoch_millis": 1521478734227,
					  "policy_id": 210972,
					  "event_type": "SystemSample",
					  "select_value": "cpuPercent",
					  "comparison": "above",
					  "critical_threshold": {
						  "value": 75,
						  "duration_minutes": 2,
						  "time_function": "all"
					  }
				  }
				]
			}
			`))
	}))

	policyID := 123

	infraAlertConditions, err := c.ListAlertInfraConditions(policyID)
	if err != nil {
		t.Log(err)
		t.Fatal("ListAlertInfraConditions error")
	}
	if len(infraAlertConditions) == 0 {
		t.Log(err)
		t.Fatal("ListAlertInfraConditions error")
	}
}

func TestCreateAlertInfraCondition(t *testing.T) {
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"data":{
				   "type":"infra_metric",
				   "name":"Disk Space Condition",
				   "enabled":true,
				   "policy_id":123,
				   "id":12345,
				   "event_type":"StorageSample",
				   "select_value":"diskFreePercent",
				   "comparison":"below",
				   "warning_threshold":{
					  "value":30,
					  "duration_minutes":2,
					  "time_function":"any"
				   }
				}
			 }
			`))
	}))

	infraAlertConditionWarning := &AlertInfraThreshold{
		Value:    30,
		Duration: 100,
		Function: "any",
	}

	infraAlertCondition := AlertInfraCondition{
		PolicyID:   123,
		Name:       "Disk Space Condition",
		Enabled:    true,
		Warning:    infraAlertConditionWarning,
		Comparison: "below",
		Event:      "StorageSample",
		Select:     "diskFreePercent",
	}

	infraAlertConditionResp, err := c.CreateAlertInfraCondition(infraAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertInfraCondition error")
	}
	if infraAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertInfraCondition error")
	}
	if infraAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
}

func TestUpdateAlertInfraCondition(t *testing.T) {
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			{
				"data":{
				   "type":"infra_metric",
				   "name":"Disk Space Condition",
				   "enabled":true,
				   "policy_id":123,
				   "id":12345,
				   "event_type":"StorageSample",
				   "select_value":"diskFreePercent",
				   "comparison":"below",
				   "warning_threshold":{
					  "value":30,
					  "duration_minutes":2,
					  "time_function":"any"
				   }
				}
			 }
			`))
	}))

	infraAlertConditionWarning := &AlertInfraThreshold{
		Value:    30,
		Duration: 100,
		Function: "any",
	}

	infraAlertCondition := AlertInfraCondition{
		PolicyID:   123,
		Name:       "Test Condition",
		Enabled:    true,
		Warning:    infraAlertConditionWarning,
		Comparison: "below",
		Event:      "StorageSample",
		Select:     "diskFreePercent",
	}

	infraAlertConditionResp, err := c.UpdateAlertInfraCondition(infraAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("UpdateAlertInfraCondition error")
	}
	if infraAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("UpdateAlertInfraCondition error")
	}
	if infraAlertConditionResp.ID != 12345 {
		t.Fatal("Condition ID was not parsed correctly")
	}
}

func TestCreateAlertInfraConditionWithIntegrationProvider(t *testing.T) {
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
            {
                "data":{
                   "type":"infra_metric",
                   "name":"ELB Healthy Host Count",
                   "enabled":true,
                   "policy_id":123,
                   "id":12346,
                   "event_type":"LoadBalancerSample",
                   "select_value":"provider.healthyHostCount.Average",
                   "comparison":"below",
                   "warning_threshold":{
                      "value":1,
                      "duration_minutes":5,
                      "time_function":"any"
                   },
                   "integration_provider": "Elb"
                }
             }
            `))
	}))

	infraAlertConditionWarning := &AlertInfraThreshold{
		Value:    1,
		Duration: 5,
		Function: "any",
	}

	infraAlertCondition := AlertInfraCondition{
		PolicyID:            123,
		Name:                "ELB Healthy Host Count",
		Enabled:             true,
		Warning:             infraAlertConditionWarning,
		Comparison:          "below",
		Event:               "LoadBalancerSample",
		Select:              "provider.healthyHostCount.Average",
		IntegrationProvider: "Elb",
	}

	infraAlertConditionResp, err := c.CreateAlertInfraCondition(infraAlertCondition)
	if err != nil {
		t.Log(err)
		t.Fatal("CreateAlertInfraCondition error")
	}
	if infraAlertConditionResp == nil {
		t.Log(err)
		t.Fatal("CreateAlertInfraCondition error")
	}
	if infraAlertConditionResp.ID != 12346 {
		t.Fatal("Condition ID was not parsed correctly")
	}
	if infraAlertConditionResp.IntegrationProvider != "Elb" {
		t.Fatal("Condition IntegrationProvider was not parsed correctly")
	}
}

func TestDeleteAlertInfraCondition(t *testing.T) {
	policyID := 123
	conditionID := 12345
	c := newTestAPIInfraClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if r.Method != "DELETE" {
			t.Fatal("DeleteAlertInfraCondition did not use DELETE method")
		}
		if r.URL.Path != fmt.Sprintf("/alerts/conditions/%v", conditionID) {
			t.Fatal("DeleteAlertInfraCondition did not use the correct URL")
		}
	}))
	err := c.DeleteAlertInfraCondition(policyID, conditionID)
	if err != nil {
		t.Log(err)
		t.Fatal("DeleteAlertInfraCondition error")
	}
}
