package api

// LabelLinks represents external references on the Label.
type LabelLinks struct {
	Applications []int `json:"applications"`
	Servers      []int `json:"servers"`
}

// Label represents a New Relic label.
type Label struct {
	Key      string     `json:"key,omitempty"`
	Category string     `json:"category,omitempty"`
	Name     string     `json:"name,omitempty"`
	Links    LabelLinks `json:"links,omitempty"`
}

// AlertPolicy represents a New Relic alert policy.
type AlertPolicy struct {
	ID                 int    `json:"id,omitempty"`
	IncidentPreference string `json:"incident_preference,omitempty"`
	Name               string `json:"name,omitempty"`
	CreatedAt          int    `json:"created_at,omitempty"`
	UpdatedAt          int    `json:"updated_at,omitempty"`
}

// AlertConditionUserDefined represents user defined metrics for the New Relic alert condition.
type AlertConditionUserDefined struct {
	Metric        string `json:"metric,omitempty"`
	ValueFunction string `json:"value_function,omitempty"`
}

// AlertConditionTerm represents the terms of a New Relic alert condition.
type AlertConditionTerm struct {
	Duration     int     `json:"duration,string,omitempty"`
	Operator     string  `json:"operator,omitempty"`
	Priority     string  `json:"priority,omitempty"`
	Threshold    float32 `json:"threshold,string,omitempty"`
	TimeFunction string  `json:"time_function,omitempty"`
}

// AlertCondition represents a New Relic alert condition.
type AlertCondition struct {
	Type        string                    `json:"type,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Enabled     bool                      `json:"enabled,omitempty"`
	Entities    []string                  `json:"entities,omitempty"`
	Metric      string                    `json:"metric,omitempty"`
	RunbookURL  string                    `json:"runbook_url,omitempty"`
	Terms       []AlertConditionTerm      `json:"terms,omitempty"`
	UserDefined AlertConditionUserDefined `json:"uder_defined,omitempty"`
}
