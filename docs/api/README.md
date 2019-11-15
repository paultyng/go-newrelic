# api
--
    import "github.com/paultyng/go-newrelic/api"


## Usage

```go
const Version string = "4.9.1"
```
Version of this library

```go
var (
	// ErrNotFound is returned when the resource was not found in New Relic.
	ErrNotFound = errors.New("newrelic: Resource not found")
)
```

#### type AlertChannel

```go
type AlertChannel struct {
	ID            int                    `json:"id,omitempty"`
	Name          string                 `json:"name,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	Links         AlertChannelLinks      `json:"links,omitempty"`
}
```

AlertChannel represents a New Relic alert notification channel

#### type AlertChannelLinks

```go
type AlertChannelLinks struct {
	PolicyIDs []int `json:"policy_ids,omitempty"`
}
```

AlertChannelLinks represent the links between policies and alert channels

#### type AlertCondition

```go
type AlertCondition struct {
	PolicyID            int                       `json:"-"`
	ID                  int                       `json:"id,omitempty"`
	Type                string                    `json:"type,omitempty"`
	Name                string                    `json:"name,omitempty"`
	Enabled             bool                      `json:"enabled"`
	Entities            []string                  `json:"entities,omitempty"`
	Metric              string                    `json:"metric,omitempty"`
	RunbookURL          string                    `json:"runbook_url,omitempty"`
	Terms               []AlertConditionTerm      `json:"terms,omitempty"`
	UserDefined         AlertConditionUserDefined `json:"user_defined,omitempty"`
	Scope               string                    `json:"condition_scope,omitempty"`
	GCMetric            string                    `json:"gc_metric,omitempty"`
	ViolationCloseTimer int                       `json:"violation_close_timer,omitempty"`
}
```

AlertCondition represents a New Relic alert condition. TODO: custom unmarshal
entities to ints?

#### type AlertConditionTerm

```go
type AlertConditionTerm struct {
	Duration     int     `json:"duration,string,omitempty"`
	Operator     string  `json:"operator,omitempty"`
	Priority     string  `json:"priority,omitempty"`
	Threshold    float64 `json:"threshold,string"`
	TimeFunction string  `json:"time_function,omitempty"`
}
```

AlertConditionTerm represents the terms of a New Relic alert condition.

#### func (*AlertConditionTerm) UnmarshalJSON

```go
func (t *AlertConditionTerm) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements custom json unmarshalling for the AlertConditionTerm
type

#### type AlertConditionUserDefined

```go
type AlertConditionUserDefined struct {
	Metric        string `json:"metric,omitempty"`
	ValueFunction string `json:"value_function,omitempty"`
}
```

AlertConditionUserDefined represents user defined metrics for the New Relic
alert condition.

#### type AlertInfraCondition

```go
type AlertInfraCondition struct {
	PolicyID            int                  `json:"policy_id,omitempty"`
	ID                  int                  `json:"id,omitempty"`
	Name                string               `json:"name,omitempty"`
	RunbookURL          string               `json:"runbook_url,omitempty"`
	Type                string               `json:"type,omitempty"`
	Comparison          string               `json:"comparison,omitempty"`
	CreatedAt           int                  `json:"created_at_epoch_millis,omitempty"`
	UpdatedAt           int                  `json:"updated_at_epoch_millis,omitempty"`
	Enabled             bool                 `json:"enabled"`
	Event               string               `json:"event_type,omitempty"`
	Select              string               `json:"select_value,omitempty"`
	Where               string               `json:"where_clause,omitempty"`
	ProcessWhere        string               `json:"process_where_clause,omitempty"`
	IntegrationProvider string               `json:"integration_provider,omitempty"`
	ViolationCloseTimer *int                 `json:"violation_close_timer,omitempty"`
	Warning             *AlertInfraThreshold `json:"warning_threshold,omitempty"`
	Critical            *AlertInfraThreshold `json:"critical_threshold,omitempty"`
}
```

AlertInfraCondition represents a New Relic Infra Alert condition.

#### type AlertInfraThreshold

```go
type AlertInfraThreshold struct {
	Value    int    `json:"value,omitempty"`
	Duration int    `json:"duration_minutes,omitempty"`
	Function string `json:"time_function,omitempty"`
}
```

AlertInfraThreshold represents an Infra alerting condition

#### type AlertNrqlCondition

```go
type AlertNrqlCondition struct {
	Nrql                AlertNrqlQuery       `json:"nrql,omitempty"`
	Terms               []AlertConditionTerm `json:"terms,omitempty"`
	Type                string               `json:"type,omitempty"`
	Name                string               `json:"name,omitempty"`
	RunbookURL          string               `json:"runbook_url,omitempty"`
	ValueFunction       string               `json:"value_function,omitempty"`
	PolicyID            int                  `json:"-"`
	ID                  int                  `json:"id,omitempty"`
	ExpectedGroups      int                  `json:"expected_groups,omitempty"`
	ViolationCloseTimer int                  `json:"violation_time_limit_seconds,omitempty"`
	Enabled             bool                 `json:"enabled"`
	IgnoreOverlap       bool                 `json:"ignore_overlap,omitempty"`
}
```

AlertNrqlCondition represents a New Relic NRQL Alert condition.

#### type AlertNrqlQuery

```go
type AlertNrqlQuery struct {
	Query      string `json:"query,omitempty"`
	SinceValue string `json:"since_value,omitempty"`
}
```

AlertNrqlQuery represents a NRQL query to use with a NRQL alert condition

#### type AlertPlugin

```go
type AlertPlugin struct {
	ID   string `json:"id,omitempty"`
	GUID string `json:"guid,omitempty"`
}
```

AlertPlugin represents a plugin to use with a Plugin alert condition.

#### type AlertPluginsCondition

```go
type AlertPluginsCondition struct {
	PolicyID          int                  `json:"-"`
	ID                int                  `json:"id,omitempty"`
	Name              string               `json:"name,omitempty"`
	Enabled           bool                 `json:"enabled"`
	Entities          []string             `json:"entities,omitempty"`
	Metric            string               `json:"metric,omitempty"`
	MetricDescription string               `json:"metric_description,omitempty"`
	RunbookURL        string               `json:"runbook_url,omitempty"`
	Terms             []AlertConditionTerm `json:"terms,omitempty"`
	ValueFunction     string               `json:"value_function,omitempty"`
	Plugin            AlertPlugin          `json:"plugin,omitempty"`
}
```

AlertPluginsCondition represents a New Relic Plugin Alert condition.

#### type AlertPolicy

```go
type AlertPolicy struct {
	ID                 int    `json:"id,omitempty"`
	IncidentPreference string `json:"incident_preference,omitempty"`
	Name               string `json:"name,omitempty"`
	CreatedAt          int64  `json:"created_at,omitempty"`
	UpdatedAt          int64  `json:"updated_at,omitempty"`
}
```

AlertPolicy represents a New Relic alert policy.

#### type AlertSyntheticsCondition

```go
type AlertSyntheticsCondition struct {
	PolicyID   int    `json:"-"`
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Enabled    bool   `json:"enabled"`
	RunbookURL string `json:"runbook_url,omitempty"`
	MonitorID  string `json:"monitor_id,omitempty"`
}
```

AlertSyntheticsCondition represents a New Relic NRQL Alert condition.

#### type Application

```go
type Application struct {
	ID             int                       `json:"id,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Language       string                    `json:"language,omitempty"`
	HealthStatus   string                    `json:"health_status,omitempty"`
	Reporting      bool                      `json:"reporting"`
	LastReportedAt string                    `json:"last_reported_at,omitempty"`
	Summary        ApplicationSummary        `json:"application_summary,omitempty"`
	EndUserSummary ApplicationEndUserSummary `json:"end_user_summary,omitempty"`
	Settings       ApplicationSettings       `json:"settings,omitempty"`
	Links          ApplicationLinks          `json:"links,omitempty"`
}
```

Application represents information about a New Relic application.

#### type ApplicationEndUserSummary

```go
type ApplicationEndUserSummary struct {
	ResponseTime float64 `json:"response_time"`
	Throughput   float64 `json:"throughput"`
	ApdexTarget  float64 `json:"apdex_target"`
	ApdexScore   float64 `json:"apdex_score"`
}
```

ApplicationEndUserSummary represents performance information about a New Relic
application.

#### type ApplicationLinks

```go
type ApplicationLinks struct {
	ServerIDs     []int `json:"servers,omitempty"`
	HostIDs       []int `json:"application_hosts,omitempty"`
	InstanceIDs   []int `json:"application_instances,omitempty"`
	AlertPolicyID int   `json:"alert_policy"`
}
```

ApplicationLinks represents all the links for a New Relic application.

#### type ApplicationSettings

```go
type ApplicationSettings struct {
	AppApdexThreshold        float64 `json:"app_apdex_threshold,omitempty"`
	EndUserApdexThreshold    float64 `json:"end_user_apdex_threshold,omitempty"`
	EnableRealUserMonitoring bool    `json:"enable_real_user_monitoring"`
	UseServerSideConfig      bool    `json:"use_server_side_config"`
}
```

ApplicationSettings represents some of the settings of a New Relic application.

#### type ApplicationSummary

```go
type ApplicationSummary struct {
	ResponseTime            float64 `json:"response_time"`
	Throughput              float64 `json:"throughput"`
	ErrorRate               float64 `json:"error_rate"`
	ApdexTarget             float64 `json:"apdex_target"`
	ApdexScore              float64 `json:"apdex_score"`
	HostCount               int     `json:"host_count"`
	InstanceCount           int     `json:"instance_count"`
	ConcurrentInstanceCount int     `json:"concurrent_instance_count"`
}
```

ApplicationSummary represents performance information about a New Relic
application.

#### type Client

```go
type Client struct {
	RestyClient *resty.Client
}
```

Client represents the client state for the API.

#### func  New

```go
func New(config Config) Client
```
New returns a new Client for the specified apiKey.

#### func (*Client) CreateAlertChannel

```go
func (c *Client) CreateAlertChannel(channel AlertChannel) (*AlertChannel, error)
```
CreateAlertChannel allows you to create an alert channel with the specified data
and links.

#### func (*Client) CreateAlertCondition

```go
func (c *Client) CreateAlertCondition(condition AlertCondition) (*AlertCondition, error)
```
CreateAlertCondition creates an alert condition given the passed configuration.

#### func (*Client) CreateAlertNrqlCondition

```go
func (c *Client) CreateAlertNrqlCondition(condition AlertNrqlCondition) (*AlertNrqlCondition, error)
```
CreateAlertNrqlCondition creates an NRQL alert condition given the passed
configuration.

#### func (*Client) CreateAlertPluginsCondition

```go
func (c *Client) CreateAlertPluginsCondition(condition AlertPluginsCondition) (*AlertPluginsCondition, error)
```
CreateAlertPluginsCondition creates an Plugin alert condition given the passed
configuration.

#### func (*Client) CreateAlertPolicy

```go
func (c *Client) CreateAlertPolicy(policy AlertPolicy) (*AlertPolicy, error)
```
CreateAlertPolicy creates a new alert policy for the account.

#### func (*Client) CreateAlertSyntheticsCondition

```go
func (c *Client) CreateAlertSyntheticsCondition(condition AlertSyntheticsCondition) (*AlertSyntheticsCondition, error)
```
CreateAlertSyntheticsCondition creates an Synthetics alert condition given the
passed configuration.

#### func (*Client) CreateDashboard

```go
func (c *Client) CreateDashboard(dashboard Dashboard) (*Dashboard, error)
```
CreateDashboard creates dashboard given the passed configuration.

#### func (*Client) CreateDeployment

```go
func (c *Client) CreateDeployment(applicationID int, deployment Deployment) (*Deployment, error)
```
CreateDeployment creates a deployment for an application.

#### func (*Client) CreateLabel

```go
func (c *Client) CreateLabel(label Label) error
```
CreateLabel creates a new label for the account.

#### func (*Client) DeleteAlertChannel

```go
func (c *Client) DeleteAlertChannel(id int) error
```
DeleteAlertChannel deletes the alert channel with the specified ID.

#### func (*Client) DeleteAlertCondition

```go
func (c *Client) DeleteAlertCondition(policyID int, id int) error
```
DeleteAlertCondition removes the alert condition given the specified ID and
policy ID.

#### func (*Client) DeleteAlertNrqlCondition

```go
func (c *Client) DeleteAlertNrqlCondition(policyID int, id int) error
```
DeleteAlertNrqlCondition removes the NRQL alert condition given the specified ID
and policy ID.

#### func (*Client) DeleteAlertPluginsCondition

```go
func (c *Client) DeleteAlertPluginsCondition(policyID, id int) error
```
DeleteAlertPluginsCondition removes the Plugin alert condition given the
specified ID and policy ID.

#### func (*Client) DeleteAlertPolicy

```go
func (c *Client) DeleteAlertPolicy(id int) error
```
DeleteAlertPolicy deletes an existing alert policy from the account.

#### func (*Client) DeleteAlertPolicyChannel

```go
func (c *Client) DeleteAlertPolicyChannel(policyID int, channelID int) error
```
DeleteAlertPolicyChannel deletes a notification channel from an alert policy.

#### func (*Client) DeleteAlertSyntheticsCondition

```go
func (c *Client) DeleteAlertSyntheticsCondition(policyID int, id int) error
```
DeleteAlertSyntheticsCondition removes the Synthetics alert condition given the
specified ID and policy ID.

#### func (*Client) DeleteApplication

```go
func (c *Client) DeleteApplication(id int) error
```
DeleteApplication deletes a non-reporting application from your account.

#### func (*Client) DeleteDashboard

```go
func (c *Client) DeleteDashboard(id int) error
```
DeleteDashboard deletes an existing dashboard given the passed configuration

#### func (*Client) DeleteDeployment

```go
func (c *Client) DeleteDeployment(applicationID, deploymentID int) error
```
DeleteDeployment deletes an application deployment from an application.

#### func (*Client) DeleteLabel

```go
func (c *Client) DeleteLabel(key string) error
```
DeleteLabel deletes a label on the account specified by key.

#### func (*Client) Do

```go
func (c *Client) Do(method string, path string, body interface{}, response interface{}) (string, error)
```
Do exectes an API request with the specified parameters.

#### func (*Client) GetAlertChannel

```go
func (c *Client) GetAlertChannel(id int) (*AlertChannel, error)
```
GetAlertChannel returns a specific alert channel by ID

#### func (*Client) GetAlertCondition

```go
func (c *Client) GetAlertCondition(policyID int, id int) (*AlertCondition, error)
```
GetAlertCondition gets information about an alert condition given an ID and
policy ID.

#### func (*Client) GetAlertNrqlCondition

```go
func (c *Client) GetAlertNrqlCondition(policyID int, id int) (*AlertNrqlCondition, error)
```
GetAlertNrqlCondition gets information about a NRQL alert condition given an ID
and policy ID.

#### func (*Client) GetAlertPluginsCondition

```go
func (c *Client) GetAlertPluginsCondition(policyID, id int) (*AlertPluginsCondition, error)
```
GetAlertPluginsCondition gets information about a plugin alert condition given
an ID and policy ID.

#### func (*Client) GetAlertPolicy

```go
func (c *Client) GetAlertPolicy(id int) (*AlertPolicy, error)
```
GetAlertPolicy returns a specific alert policy by ID

#### func (*Client) GetAlertSyntheticsCondition

```go
func (c *Client) GetAlertSyntheticsCondition(policyID int, id int) (*AlertSyntheticsCondition, error)
```
GetAlertSyntheticsCondition gets information about a Synthetics alert condition
given an ID and policy ID.

#### func (*Client) GetDashboard

```go
func (c *Client) GetDashboard(id int) (*Dashboard, error)
```
GetDashboard returns a specific dashboard for the account.

#### func (*Client) GetKeyTransaction

```go
func (c *Client) GetKeyTransaction(id int) (*KeyTransaction, error)
```
GetKeyTransaction returns a specific key transaction by ID.

#### func (*Client) GetLabel

```go
func (c *Client) GetLabel(key string) (*Label, error)
```
GetLabel gets the label for the specified key.

#### func (*Client) ListAlertChannels

```go
func (c *Client) ListAlertChannels() ([]AlertChannel, error)
```
ListAlertChannels returns all alert policies for the account.

#### func (*Client) ListAlertConditions

```go
func (c *Client) ListAlertConditions(policyID int) ([]AlertCondition, error)
```
ListAlertConditions returns alert conditions for the specified policy.

#### func (*Client) ListAlertNrqlConditions

```go
func (c *Client) ListAlertNrqlConditions(policyID int) ([]AlertNrqlCondition, error)
```
ListAlertNrqlConditions returns NRQL alert conditions for the specified policy.

#### func (*Client) ListAlertPluginsConditions

```go
func (c *Client) ListAlertPluginsConditions(policyID int) ([]AlertPluginsCondition, error)
```
ListAlertPluginsConditions returns Plugin alert conditions for the specified
policy.

#### func (*Client) ListAlertPolicies

```go
func (c *Client) ListAlertPolicies() ([]AlertPolicy, error)
```
ListAlertPolicies returns all alert policies for the account.

#### func (*Client) ListAlertSyntheticsConditions

```go
func (c *Client) ListAlertSyntheticsConditions(policyID int) ([]AlertSyntheticsCondition, error)
```
ListAlertSyntheticsConditions returns Synthetics alert conditions for the
specified policy.

#### func (*Client) ListApplications

```go
func (c *Client) ListApplications() ([]Application, error)
```
ListApplications lists all the applications you have access to.

#### func (*Client) ListComponentMetricData

```go
func (c *Client) ListComponentMetricData(componentID int, names []string) ([]Metric, error)
```
ListComponentMetricData lists all the metric data for the specified component ID
and metric names.

#### func (*Client) ListComponentMetrics

```go
func (c *Client) ListComponentMetrics(componentID int) ([]ComponentMetric, error)
```
ListComponentMetrics lists all the component metrics for the specificed
component ID.

#### func (*Client) ListComponents

```go
func (c *Client) ListComponents(pluginID int) ([]Component, error)
```
ListComponents lists all the components for the specified plugin ID.

#### func (*Client) ListDashboards

```go
func (c *Client) ListDashboards() ([]Dashboard, error)
```
ListDashboards returns all dashboards for the account.

#### func (*Client) ListDeployments

```go
func (c *Client) ListDeployments(id int) ([]Deployment, error)
```
ListDeployments returns deployments by newrelic applicationID.

#### func (*Client) ListKeyTransactions

```go
func (c *Client) ListKeyTransactions() ([]KeyTransaction, error)
```
ListKeyTransactions returns all key transactions for the account.

#### func (*Client) ListLabels

```go
func (c *Client) ListLabels() ([]Label, error)
```
ListLabels returns the labels for the account.

#### func (*Client) ListPlugins

```go
func (c *Client) ListPlugins() ([]Plugin, error)
```
ListPlugins lists all the plugins you have access to.

#### func (*Client) UpdateAlertCondition

```go
func (c *Client) UpdateAlertCondition(condition AlertCondition) (*AlertCondition, error)
```
UpdateAlertCondition updates an alert condition with the specified changes.

#### func (*Client) UpdateAlertNrqlCondition

```go
func (c *Client) UpdateAlertNrqlCondition(condition AlertNrqlCondition) (*AlertNrqlCondition, error)
```
UpdateAlertNrqlCondition updates a NRQL alert condition with the specified
changes.

#### func (*Client) UpdateAlertPluginsCondition

```go
func (c *Client) UpdateAlertPluginsCondition(condition AlertPluginsCondition) (*AlertPluginsCondition, error)
```
UpdateAlertPluginsCondition updates a Plugin alert condition with the specified
changes.

#### func (*Client) UpdateAlertPolicy

```go
func (c *Client) UpdateAlertPolicy(policy AlertPolicy) (*AlertPolicy, error)
```
UpdateAlertPolicy updates an alert policy with the specified changes

#### func (*Client) UpdateAlertPolicyChannels

```go
func (c *Client) UpdateAlertPolicyChannels(policyID int, channelIDs []int) error
```
UpdateAlertPolicyChannels updates a policy by adding the specified notification
channels.

#### func (*Client) UpdateAlertSyntheticsCondition

```go
func (c *Client) UpdateAlertSyntheticsCondition(condition AlertSyntheticsCondition) (*AlertSyntheticsCondition, error)
```
UpdateAlertSyntheticsCondition updates a Synthetics alert condition with the
specified changes.

#### func (*Client) UpdateDashboard

```go
func (c *Client) UpdateDashboard(dashboard Dashboard) (*Dashboard, error)
```
UpdateDashboard updates a dashboard given the passed configuration

#### type Component

```go
type Component struct {
	ID             int             `json:"id"`
	Name           string          `json:"name,omitempty"`
	HealthStatus   string          `json:"health_status,omitempty"`
	SummaryMetrics []SummaryMetric `json:"summary_metrics"`
}
```

Component represnets information about a New Relic component.

#### type ComponentMetric

```go
type ComponentMetric struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values"`
}
```

ComponentMetric represents metric information for a specific component.

#### type Config

```go
type Config struct {
	// APIKey is the Admin API Key for your New Relic account.
	// This parameter is required.
	APIKey string

	// BaseURL is the base API URL for the client.
	// `Client` defaults to `https://api.newrelic.com/v2`.
	// Use `https://api.eu.newrelic.com/v2` for EU-based accounts.
	// `InfraClient` defaults to `https://infra-api.newrelic.com/v2`.
	// Use `https://intra-api.eu.newrelic.com/v2` for EU-based accounts.
	BaseURL string

	// ProxyURL sets the Resty client's proxy URL (optional).
	ProxyURL string

	// Debug sets the Resty client's debug mode.
	// Defaults to `false`.
	Debug bool

	// TLSConfig is passed to the Resty client's SetTLSClientConfig method (optional).
	// Used to set a custom root certificate or disable security.
	TLSConfig *tls.Config

	// UserAgent is passed to the Resty client's SetHeaders to allow overriding
	// the default user-agent header (go-newrelic/$version)
	UserAgent string

	// HttpTransport is passed to the Resty client's SetTransport method (optional).
	HTTPTransport http.RoundTripper
}
```

Config contains all the configuration data for the API Client.

#### type Dashboard

```go
type Dashboard struct {
	ID         int               `json:"id"`
	Title      string            `json:"title,omitempty"`
	Icon       string            `json:"icon,omitempty"`
	CreatedAt  string            `json:"created_at,omitempty"`
	UpdatedAt  string            `json:"updated_at,omitempty"`
	Visibility string            `json:"visibility,omitempty"`
	Editable   string            `json:"editable,omitempty"`
	UIURL      string            `json:"ui_url,omitempty"`
	APIRL      string            `json:"api_url,omitempty"`
	OwnerEmail string            `json:"owner_email,omitempty"`
	Metadata   DashboardMetadata `json:"metadata"`
	Filter     DashboardFilter   `json:"filter,omitempty"`
	Widgets    []DashboardWidget `json:"widgets,omitempty"`
}
```

Dashboard represents information about a New Relic dashboard.

#### type DashboardFilter

```go
type DashboardFilter struct {
	EventTypes []string `json:"event_types,omitempty"`
	Attributes []string `json:"attributes,omitempty"`
}
```

DashboardFilter represents the filter in a dashboard.

#### type DashboardMetadata

```go
type DashboardMetadata struct {
	Version int `json:"version"`
}
```

DashboardMetadata represents metadata about the dashboard (like version)

#### type DashboardWidget

```go
type DashboardWidget struct {
	Visualization string                      `json:"visualization,omitempty"`
	ID            int                         `json:"widget_id,omitempty"`
	AccountID     int                         `json:"account_id,omitempty"`
	Data          []DashboardWidgetData       `json:"data,omitempty"`
	Presentation  DashboardWidgetPresentation `json:"presentation,omitempty"`
	Layout        DashboardWidgetLayout       `json:"layout,omitempty"`
}
```

DashboardWidget represents a widget in a dashboard.

#### type DashboardWidgetData

```go
type DashboardWidgetData struct {
	NRQL          string                           `json:"nrql,omitempty"`
	Source        string                           `json:"source,omitempty"`
	Duration      int                              `json:"duration,omitempty"`
	EndTime       int                              `json:"end_time,omitempty"`
	EntityIds     []int                            `json:"entity_ids,omitempty"`
	CompareWith   []DashboardWidgetDataCompareWith `json:"compare_with,omitempty"`
	Metrics       []DashboardWidgetDataMetric      `json:"metrics,omitempty"`
	RawMetricName string                           `json:"raw_metric_name,omitempty"`
	Facet         string                           `json:"facet,omitempty"`
	OrderBy       string                           `json:"order_by,omitempty"`
	Limit         int                              `json:"limit,omitempty"`
}
```

DashboardWidgetData represents the data backing a dashboard widget.

#### type DashboardWidgetDataCompareWith

```go
type DashboardWidgetDataCompareWith struct {
	OffsetDuration string                                     `json:"offset_duration,omitempty"`
	Presentation   DashboardWidgetDataCompareWithPresentation `json:"presentation,omitempty"`
}
```

DashboardWidgetDataCompareWith represents the compare with configuration of the
widget.

#### type DashboardWidgetDataCompareWithPresentation

```go
type DashboardWidgetDataCompareWithPresentation struct {
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}
```

DashboardWidgetDataCompareWithPresentation represents the compare with
presentation configuration of the widget.

#### type DashboardWidgetDataMetric

```go
type DashboardWidgetDataMetric struct {
	Name   string   `json:"name,omitempty"`
	Units  string   `json:"units,omitempty"`
	Scope  string   `json:"scope,omitempty"`
	Values []string `json:"values,omitempty"`
}
```

DashboardWidgetDataMetric represents the metrics data of the widget.

#### type DashboardWidgetLayout

```go
type DashboardWidgetLayout struct {
	Width  int `json:"width"`
	Height int `json:"height"`
	Row    int `json:"row"`
	Column int `json:"column"`
}
```

DashboardWidgetLayout represents the layout of a widget in a dashboard.

#### type DashboardWidgetPresentation

```go
type DashboardWidgetPresentation struct {
	Title                string                    `json:"title,omitempty"`
	Notes                string                    `json:"notes,omitempty"`
	DrilldownDashboardID int                       `json:"drilldown_dashboard_id,omitempty"`
	Threshold            *DashboardWidgetThreshold `json:"threshold,omitempty"`
}
```

DashboardWidgetPresentation represents the visual presentation of a dashboard
widget.

#### type DashboardWidgetThreshold

```go
type DashboardWidgetThreshold struct {
	Red    float64 `json:"red,omitempty"`
	Yellow float64 `json:"yellow,omitempty"`
}
```

DashboardWidgetThreshold represents the threshold configuration of a dashboard
widget.

#### type Deployment

```go
type Deployment struct {
	ID          int    `json:"id,omitempty"`
	Revision    string `json:"revision"`
	Changelog   string `json:"changelog,omitempty"`
	Description string `json:"description,omitempty"`
	User        string `json:"user,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
}
```

Deployment represents information about a New Relic application deployment.

#### type ErrorDetail

```go
type ErrorDetail struct {
	Title string `json:"title,omitempty"`
}
```

ErrorDetail represents the details of an ErrorResponse from New Relic.

#### type ErrorResponse

```go
type ErrorResponse struct {
	Detail *ErrorDetail `json:"error,omitempty"`
}
```

ErrorResponse represents an error response from New Relic.

#### func (*ErrorResponse) Error

```go
func (e *ErrorResponse) Error() string
```

#### type InfraClient

```go
type InfraClient struct {
	Client
}
```

InfraClient represents the client state for the Infrastructure API

#### func  NewInfraClient

```go
func NewInfraClient(config Config) InfraClient
```
NewInfraClient returns a new InfraClient for the specified apiKey.

#### func (*InfraClient) CreateAlertInfraCondition

```go
func (c *InfraClient) CreateAlertInfraCondition(condition AlertInfraCondition) (*AlertInfraCondition, error)
```
CreateAlertInfraCondition creates an Infra alert condition given the passed
configuration.

#### func (*InfraClient) DeleteAlertInfraCondition

```go
func (c *InfraClient) DeleteAlertInfraCondition(policyID int, id int) error
```
DeleteAlertInfraCondition removes the Infra alert condition given the specified
ID and policy ID.

#### func (*InfraClient) GetAlertInfraCondition

```go
func (c *InfraClient) GetAlertInfraCondition(policyID int, id int) (*AlertInfraCondition, error)
```
GetAlertInfraCondition gets information about a Infra alert condition given an
ID and policy ID.

#### func (*InfraClient) ListAlertInfraConditions

```go
func (c *InfraClient) ListAlertInfraConditions(policyID int) ([]AlertInfraCondition, error)
```
ListAlertInfraConditions returns Infra alert conditions for the specified
policy.

#### func (*InfraClient) UpdateAlertInfraCondition

```go
func (c *InfraClient) UpdateAlertInfraCondition(condition AlertInfraCondition) (*AlertInfraCondition, error)
```
UpdateAlertInfraCondition updates an Infra alert condition with the specified
changes.

#### type KeyTransaction

```go
type KeyTransaction struct {
	ID              int                       `json:"id,omitempty"`
	Name            string                    `json:"name,omitempty"`
	TransactionName string                    `json:"transaction_name,omitempty"`
	HealthStatus    string                    `json:"health_status,omitempty"`
	Reporting       bool                      `json:"reporting"`
	LastReportedAt  string                    `json:"last_reported_at,omitempty"`
	Summary         ApplicationSummary        `json:"application_summary,omitempty"`
	EndUserSummary  ApplicationEndUserSummary `json:"end_user_summary,omitempty"`
	Links           ApplicationLinks          `json:"links,omitempty"`
}
```

KeyTransaction represents information about a New Relic key transaction.

#### type Label

```go
type Label struct {
	Key      string     `json:"key,omitempty"`
	Category string     `json:"category,omitempty"`
	Name     string     `json:"name,omitempty"`
	Links    LabelLinks `json:"links,omitempty"`
}
```

Label represents a New Relic label.

#### type LabelLinks

```go
type LabelLinks struct {
	Applications []int `json:"applications"`
	Servers      []int `json:"servers"`
}
```

LabelLinks represents external references on the Label.

#### type Metric

```go
type Metric struct {
	Name       string            `json:"name"`
	Timeslices []MetricTimeslice `json:"timeslices"`
}
```

Metric represents data for a specific metric.

#### type MetricThreshold

```go
type MetricThreshold struct {
	Caution  float64 `json:"caution"`
	Critical float64 `json:"critical"`
}
```

MetricThreshold represents the different thresholds for a metric in an alert.

#### type MetricTimeslice

```go
type MetricTimeslice struct {
	From   string                 `json:"from,omitempty"`
	To     string                 `json:"to,omitempty"`
	Values map[string]interface{} `json:"values,omitempty"`
}
```

MetricTimeslice represents the values of a metric over a given time.

#### type MetricValue

```go
type MetricValue struct {
	Raw       float64 `json:"raw"`
	Formatted string  `json:"formatted"`
}
```

MetricValue represents the observed value of a metric.

#### type Plugin

```go
type Plugin struct {
	ID                  int             `json:"id"`
	Name                string          `json:"name,omitempty"`
	GUID                string          `json:"guid,omitempty"`
	Publisher           string          `json:"publisher,omitempty"`
	ComponentAgentCount int             `json:"component_agent_count"`
	Details             PluginDetails   `json:"details"`
	SummaryMetrics      []SummaryMetric `json:"summary_metrics"`
}
```

Plugin represents information about a New Relic plugin.

#### type PluginDetails

```go
type PluginDetails struct {
	Description           int    `json:"description"`
	IsPublic              string `json:"is_public"`
	CreatedAt             string `json:"created_at,omitempty"`
	UpdatedAt             string `json:"updated_at,omitempty"`
	LastPublishedAt       string `json:"last_published_at,omitempty"`
	HasUnpublishedChanges bool   `json:"has_unpublished_changes"`
	BrandingImageURL      string `json:"branding_image_url"`
	UpgradedAt            string `json:"upgraded_at,omitempty"`
	ShortName             string `json:"short_name"`
	PublisherAboutURL     string `json:"publisher_about_url"`
	PublisherSupportURL   string `json:"publisher_support_url"`
	DownloadURL           string `json:"download_url"`
	FirstEditedAt         string `json:"first_edited_at,omitempty"`
	LastEditedAt          string `json:"last_edited_at,omitempty"`
	FirstPublishedAt      string `json:"first_published_at,omitempty"`
	PublishedVersion      string `json:"published_version"`
}
```

PluginDetails represents information about a New Relic plugin.

#### type SummaryMetric

```go
type SummaryMetric struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Metric        string          `json:"metric"`
	ValueFunction string          `json:"value_function"`
	Thresholds    MetricThreshold `json:"thresholds"`
	Values        MetricValue     `json:"values"`
}
```

SummaryMetric represents summary information for a specific metric.
