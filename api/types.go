package api

type NotFoundError struct {
}

func newNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

func (e *NotFoundError) Error() string {
	return "Resource Not Found"
}

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
	Threshold    float64 `json:"threshold,string,omitempty"`
	TimeFunction string  `json:"time_function,omitempty"`
}

// AlertCondition represents a New Relic alert condition.
type AlertCondition struct {
	PolicyID    int                       `json:"-"`
	ID          int                       `json:"id,omitempty"`
	Type        string                    `json:"type,omitempty"`
	Name        string                    `json:"name,omitempty"`
	Enabled     bool                      `json:"enabled,omitempty"`
	Entities    []string                  `json:"entities,omitempty"`
	Metric      string                    `json:"metric,omitempty"`
	RunbookURL  string                    `json:"runbook_url,omitempty"`
	Terms       []AlertConditionTerm      `json:"terms,omitempty"`
	UserDefined AlertConditionUserDefined `json:"uder_defined,omitempty"`
}

// AlertChannelLinks represent the links between policies and alert channels
type AlertChannelLinks struct {
	PolicyIDs []int `json:"policy_ids,omitempty"`
}

// AlertChannel represents a New Relic alert notification channel
type AlertChannel struct {
	ID            int               `json:"id,omitempty"`
	Name          string            `json:"name,omitempty"`
	Type          string            `json:"type,omitempty"`
	Configuration map[string]string `json:"configuration,omitempty"`
	Links         AlertChannelLinks `json:"links,omitempty"`
}
