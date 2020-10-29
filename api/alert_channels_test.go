package api

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListAlertChannels(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		{
			"channels": [
			  {
				"id": 1001,
				"name": "webhook",
				"type": "webhook",
				"configuration": {
				  "headers": {
					"key": "value"
				  },
				  "payload_type": "application/json",
				  "payload": {
					"account_id": "$ACCOUNT_ID",
					"account_name": "$ACCOUNT_NAME",
					"closed_violations_count_critical": "$CLOSED_VIOLATIONS_COUNT_CRITICAL",
					"closed_violations_count_warning": "$CLOSED_VIOLATIONS_COUNT_WARNING",
					"condition_family_id": "$CONDITION_FAMILY_ID",
					"condition_name": "$CONDITION_NAME",
					"current_state": "$EVENT_STATE",
					"details": "$EVENT_DETAILS",
					"duration": "$DURATION",
					"event_type": "$EVENT_TYPE",
					"incident_acknowledge_url": "$INCIDENT_ACKNOWLEDGE_URL",
					"incident_id": "$INCIDENT_ID",
					"incident_url": "$INCIDENT_URL",
					"metadata": "$METADATA",
					"open_violations_count_critical": "$OPEN_VIOLATIONS_COUNT_CRITICAL",
					"open_violations_count_warning": "$OPEN_VIOLATIONS_COUNT_WARNING",
					"owner": "$EVENT_OWNER",
					"policy_name": "$POLICY_NAME",
					"policy_url": "$POLICY_URL",
					"runbook_url": "$RUNBOOK_URL",
					"severity": "$SEVERITY",
					"targets": "$TARGETS",
					"timestamp": "$TIMESTAMP",
					"violation_callback_url": "$VIOLATION_CALLBACK_URL",
					"violation_chart_url": "$VIOLATION_CHART_URL"
				  },
				  "auth_username": "auth_username",
				  "base_url": "https://test.com"
				},
				"links": {
					"policy_ids": [ 3001 ]
				}
			  },
			  {
				"id": 1002,
				"name": "xmatters",
				"type": "xmatters",
				"configuration": {
				  "channel": "channel",
				  "url": "https://test.com"
				}
			  },
			  {
				"id": 1003,
				"name": "victorops",
				"type": "victorops",
				"configuration": {
				  "route_key": "route_key"
				}
			  },
			  {
				"id": 1004,
				"name": "slack",
				"type": "slack",
				"configuration": {
				  "channel": "channel"
				}
			  },
			  {
				"id": 1005,
				"name": "opsgenie",
				"type": "opsgenie",
				"configuration": {
				  "teams": "teams",
				  "recipients": "recipients",
				  "region": "US",
				  "tags": "tags"
				}
			  },
			  {
				"id": 1006,
				"name": " <test@test.com>",
				"type": "user",
				"configuration": {
				  "user_id": "2001"
				}
			  },
			  {
				"id": 1007,
				"name": "email",
				"type": "email",
				"configuration": {
				  "include_json_attachment": "1",
				  "recipients": "test@test.com"
				}
			  },
			  {
				"id": 1008,
				"name": "PagerDuty",
				"type": "pagerduty"
			  }
			],
			"links": {
				"channel.policy_ids": "/v2/policies/{policy_id}"
			}
        }
			`))
	}))

	channels := []AlertChannel{
		{
			ID:   1001,
			Name: "webhook",
			Type: "webhook",
			Configuration: &AlertChannelConfiguration{
				Headers: &map[string]string{
					"key": "value",
				},
				PayloadType: "application/json",
				Payload: &map[string]string{
					"account_id":                       "$ACCOUNT_ID",
					"account_name":                     "$ACCOUNT_NAME",
					"closed_violations_count_critical": "$CLOSED_VIOLATIONS_COUNT_CRITICAL",
					"closed_violations_count_warning":  "$CLOSED_VIOLATIONS_COUNT_WARNING",
					"condition_family_id":              "$CONDITION_FAMILY_ID",
					"condition_name":                   "$CONDITION_NAME",
					"current_state":                    "$EVENT_STATE",
					"details":                          "$EVENT_DETAILS",
					"duration":                         "$DURATION",
					"event_type":                       "$EVENT_TYPE",
					"incident_acknowledge_url":         "$INCIDENT_ACKNOWLEDGE_URL",
					"incident_id":                      "$INCIDENT_ID",
					"incident_url":                     "$INCIDENT_URL",
					"metadata":                         "$METADATA",
					"open_violations_count_critical":   "$OPEN_VIOLATIONS_COUNT_CRITICAL",
					"open_violations_count_warning":    "$OPEN_VIOLATIONS_COUNT_WARNING",
					"owner":                            "$EVENT_OWNER",
					"policy_name":                      "$POLICY_NAME",
					"policy_url":                       "$POLICY_URL",
					"runbook_url":                      "$RUNBOOK_URL",
					"severity":                         "$SEVERITY",
					"targets":                          "$TARGETS",
					"timestamp":                        "$TIMESTAMP",
					"violation_callback_url":           "$VIOLATION_CALLBACK_URL",
					"violation_chart_url":              "$VIOLATION_CHART_URL",
				},
				BaseURL:      "https://test.com",
				AuthUsername: "auth_username",
			},
			Links: AlertChannelLinks{
				PolicyIDs: []int{3001},
			},
		},
		{
			ID:   1002,
			Name: "xmatters",
			Type: "xmatters",
			Configuration: &AlertChannelConfiguration{
				URL:     "https://test.com",
				Channel: "channel",
			},
		},
		{
			ID:   1003,
			Name: "victorops",
			Type: "victorops",
			Configuration: &AlertChannelConfiguration{
				RouteKey: "route_key",
			},
		},
		{
			ID:   1004,
			Name: "slack",
			Type: "slack",
			Configuration: &AlertChannelConfiguration{
				Channel: "channel",
			},
		},
		{
			ID:   1005,
			Name: "opsgenie",
			Type: "opsgenie",
			Configuration: &AlertChannelConfiguration{
				Teams:      "teams",
				Recipients: "recipients",
				Region:     "US",
				Tags:       "tags",
			},
		},
		{
			ID:   1006,
			Name: " <test@test.com>",
			Type: "user",
			Configuration: &AlertChannelConfiguration{
				UserID: "2001",
			},
		},
		{
			ID:   1007,
			Name: "email",
			Type: "email",
			Configuration: &AlertChannelConfiguration{
				IncludeJSONAttachment: "1",
				Recipients:            "test@test.com",
			},
		},
		{
			ID:   1008,
			Name: "PagerDuty",
			Type: "pagerduty",
		},
	}

	channelsResp, err := c.ListAlertChannels()

	if err != nil {
		t.Log(err)
		t.Fatal("ListAlertChannels error")
	}
	if channelsResp == nil {
		t.Log(err)
		t.Fatal("ListAlertChannels error")
	}
	if diff := cmp.Diff(channelsResp, channels); diff != "" {
		t.Fatalf("Alert channels not parsed correctly: %s", diff)
	}
}
