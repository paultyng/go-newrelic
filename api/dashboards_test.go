package api

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDashboards_Basic(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
				"dashboards": [
				  {
					"id": 129507,
					"title": "Election!",
					"icon": "bar-chart",
					"created_at": "2016-02-20T01:57:58Z",
					"updated_at": "2016-09-27T22:59:21Z",
					"visibility": "all",
					"editable": "editable_by_all",
					"ui_url": "https://insights.newrelic.com/accounts/1136088/dashboards/129507",
					"api_url": "https://api.newrelic.com/v2/dashboards/129507",
					"owner_email": "csmith+sandbox@newrelic.com",
					"filter": null
				  }
				]
			}
    `))
		if err != nil {
			t.Log(err)
		}
	}))

	apps, err := c.queryDashboards()
	if err != nil {
		t.Log(err)
		t.Fatal("queryDashboards error")
	}

	if len(apps) == 0 {
		t.Fatal("No dashboards found")
	}
}

func TestGetDashboard(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
		{
			"dashboard":{
				"id":1234,
				"title":"Test",
				"icon":"bar-chart",
				"created_at":"2016-02-20T01:57:58Z",
				"updated_at":"2016-09-27T22:59:21Z",
				"visibility":"all",
				"editable":"editable_by_all",
				"ui_url":"https://insights.newrelic.com/accounts/1136088/dashboards/129507",
				"api_url":"https://api.newrelic.com/v2/dashboards/129507",
				"owner_email":"foo@bar.com",
				"metadata":{
					"version":1
				},
				"filter":null,
				"widgets":[
					{
						"visualization":"billboard",
						"account_id":1,
						"data":[
							{
								"nrql":"SELECT percentile(duration, 95) FROM SyntheticCheck FACET monitorName since 7 days ago"
							}
						],
						"presentation":{
							"title":"95th Percentile Load Time (ms)",
							"notes":null,
							"drilldown_dashboard_id":null
						},
						"layout":{
							"width":1,
							"height":1,
							"row":1,
							"column":1
						},
						"threshold":{
							"red":100,
							"yellow":50
						}
					},
					{
						"visualization":"markdown",
						"account_id":1,
						"data":[
							{
								"source":"[test link](https://test.com)"
							}
						],
						"presentation":{
							"title":"Links",
							"notes":null,
							"drilldown_dashboard_id":null
						},
						"layout":{
							"width":1,
							"height":1,
							"row":1,
							"column":2
						}
					},
					{
						"visualization":"metric_line_chart",
						"account_id":1,
						"data":[
							{
								"duration":1800000,
								"end_time":1800000000000,
								"entity_ids":[
									1234
								],
								"compare_with":[
									{
										"offset_duration": "P7D",
										"presentation": {
											"name": "Last week",
											"color": "#b1b6ba"
										}
									},
									{
										"offset_duration": "P1D",
										"presentation": {
											"name": "Yesterday",
											"color": "#77add4"
										}
									}
								  ],
								"metrics":[
									{
										"name":"CPU/System/Utilization",
										"units":null,
										"scope":"",
										"values":[
											"percent"
										]
									}
								],
								"order_by":"score",
								"limit":10,
								"facet":"host",
								"raw_metric_name":"CPU/System/Utilization"
							}
						],
						"presentation":{
							"title":"Links",
							"notes":null
						},
						"layout":{
							"width":1,
							"height":1,
							"row":1,
							"column":3
						}
					}
				]
			}
		}
    `))
		if err != nil {
			t.Log(err)
		}
	}))

	dashboard, err := c.GetDashboard(1234)
	if err != nil {
		t.Log(err)
		t.Fatal("getDashboard error")
	}

	if len(dashboard.Widgets) == 0 {
		t.Fatal("Dashboard widgets not found")
	}
}

func TestCreateDashboardCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
		{
			"dashboard":{
				"id":1234,
				"title":"Test",
				"icon":"bar-chart",
				"created_at":"2016-02-20T01:57:58Z",
				"updated_at":"2016-09-27T22:59:21Z",
				"visibility":"all",
				"editable":"editable_by_all",
				"ui_url":"https://insights.newrelic.com/accounts/1136088/dashboards/129507",
				"api_url":"https://api.newrelic.com/v2/dashboards/129507",
				"owner_email":"foo@bar.com",
				"metadata":{
					"version":1
				},
				"filter":null,
				"widgets":[
					{
						"visualization":"billboard",
						"account_id":1,
						"data":[
							{
								"nrql":"SELECT percentile(duration, 95) FROM SyntheticCheck FACET monitorName since 7 days ago"
							}
						],
						"presentation":{
							"title":"95th Percentile Load Time (ms)",
							"notes":null,
							"drilldown_dashboard_id":null,
							"threshold":{
								"red":100,
								"yellow":50
							}
						},
						"layout":{
							"width":1,
							"height":1,
							"row":1,
							"column":1
						}
					},
					{
						"visualization":"markdown",
						"account_id":1,
						"data":[
							{
								"source":"[test link](https://test.com)"
							}
						],
						"presentation":{
							"title":"Links",
							"notes":null,
							"drilldown_dashboard_id":null
						},
						"layout":{
							"width":1,
							"height":1,
							"row":1,
							"column":2
						}
					},
					{
						"visualization":"metric_line_chart",
						"account_id":1,
						"data":[
							{
								"duration":1800000,
								"end_time":1800000000000,
								"entity_ids":[
									1234
								],
								"compare_with":[
									{
										"offset_duration": "P7D",
										"presentation": {
											"name": "Last week",
											"color": "#b1b6ba"
										}
									},
									{
										"offset_duration": "P1D",
										"presentation": {
											"name": "Yesterday",
											"color": "#77add4"
										}
									}
								  ],
								"metrics":[
									{
										"name":"CPU/System/Utilization",
										"units":null,
										"scope":"",
										"values":[
											"percent"
										]
									}
								],
								"order_by":"score",
								"limit":10,
								"facet":"host",
								"raw_metric_name":"CPU/System/Utilization"
							}
						],
						"presentation":{
							"title":"Links",
							"notes":null
						},
						"layout":{
							"width":1,
							"height":1,
							"row":1,
							"column":3
						}
					}
				]
			}
		}
		`))
		if err != nil {
			t.Log(err)
		}
	}))

	billboardWidgetLayout := DashboardWidgetLayout{
		Width:  1,
		Height: 1,
		Row:    1,
		Column: 1,
	}

	billboardWidgetPresentation := DashboardWidgetPresentation{
		Title: "95th Percentile Load Time (ms)",
		Notes: "",
		Threshold: &DashboardWidgetThreshold{
			Red:    100,
			Yellow: 50,
		},
	}

	billboardWidgetData := []DashboardWidgetData{
		{
			NRQL: "SELECT percentile(duration, 95) FROM SyntheticCheck FACET monitorName since 7 days ago",
		},
	}

	markdownWidgetLayout := DashboardWidgetLayout{
		Width:  1,
		Height: 1,
		Row:    1,
		Column: 2,
	}

	markdownWidgetPresentation := DashboardWidgetPresentation{
		Title: "Links",
		Notes: "",
	}

	markdownWidgetData := []DashboardWidgetData{
		{
			Source: "[test link](https://test.com)",
		},
	}

	metricsWidgetLayout := DashboardWidgetLayout{
		Width:  1,
		Height: 1,
		Row:    1,
		Column: 3,
	}

	metricsWidgetPresentation := DashboardWidgetPresentation{
		Title: "Links",
		Notes: "",
	}

	metricsWidgetData := []DashboardWidgetData{
		{
			Duration: 1800000,
			EndTime:  1800000000000,
			EntityIds: []int{
				1234,
			},
			CompareWith: []DashboardWidgetDataCompareWith{
				{
					OffsetDuration: "P7D",
					Presentation: DashboardWidgetDataCompareWithPresentation{
						Name:  "Last week",
						Color: "#b1b6ba",
					},
				},
				{
					OffsetDuration: "P1D",
					Presentation: DashboardWidgetDataCompareWithPresentation{
						Name:  "Yesterday",
						Color: "#77add4",
					},
				},
			},
			Metrics: []DashboardWidgetDataMetric{
				{
					Name:  "CPU/System/Utilization",
					Units: "",
					Scope: "",
					Values: []string{
						"percent",
					},
				},
			},
			RawMetricName: "CPU/System/Utilization",
			Facet:         "host",
			OrderBy:       "score",
			Limit:         10,
		},
	}

	dashboardWidgets := []DashboardWidget{
		{
			Visualization: "billboard",
			AccountID:     1,
			Data:          billboardWidgetData,
			Presentation:  billboardWidgetPresentation,
			Layout:        billboardWidgetLayout,
		},
		{
			Visualization: "markdown",
			AccountID:     1,
			Data:          markdownWidgetData,
			Presentation:  markdownWidgetPresentation,
			Layout:        markdownWidgetLayout,
		},
		{
			Visualization: "metric_line_chart",
			AccountID:     1,
			Data:          metricsWidgetData,
			Presentation:  metricsWidgetPresentation,
			Layout:        metricsWidgetLayout,
		},
	}

	dashboardMetadata := DashboardMetadata{
		Version: 1,
	}

	dashboardFilter := DashboardFilter{}

	dashboard := Dashboard{
		Title:    "Test",
		Icon:     "bar_chart",
		Widgets:  dashboardWidgets,
		Metadata: dashboardMetadata,
		Filter:   dashboardFilter,
	}

	dashboardResp, err := c.CreateDashboard(dashboard)

	if err != nil {
		t.Log(err)
		t.Fatal("CreateDashboard error")
	}
	if dashboardResp == nil {
		t.Log(err)
		t.Fatal("CreateDashboard error")
	}
	if dashboardResp.Metadata.Version != 1 {
		t.Fatal("CreateDashboard metadata version incorrect")
	}
	if dashboardResp.ID != 1234 {
		t.Fatal("CreateDashboard ID was not parsed correctly")
	}
	if diff := cmp.Diff(dashboardResp.Widgets, dashboard.Widgets); diff != "" {
		t.Fatalf("Widgets not parsed correctly: %s", diff)
	}
}

func TestCreateDashboardWithFilter(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`
			{
				"dashboard": {
				  "id": 1234,
				  "title": "Test",
				  "icon": "bar-chart",
				  "created_at": "2016-02-20T01:57:58Z",
				  "updated_at": "2016-09-27T22:59:21Z",
				  "visibility": "all",
				  "editable": "editable_by_all",
				  "ui_url": "https://insights.newrelic.com/accounts/1136088/dashboards/129507",
				  "api_url": "https://api.newrelic.com/v2/dashboards/129507",
				  "owner_email": "foo@bar.com",
				  "metadata": {
					"version": 1
				  },
				  "filter": {
						"event_types": [
							"test_event"
						],
						"attributes": [
							"test_attribute1",
							"test_attribute2"
						]
					},
				  "widgets": [
					{
					  "visualization": "facet_bar_chart",
					  "account_id": 1,
					  "data": [
						{
						  "nrql": "SELECT percentile(duration, 95) FROM SyntheticCheck FACET monitorName since 7 days ago"
						}
					  ],
					  "presentation": {
						"title": "95th Percentile Load Time (ms)",
						"notes": null,
						"drilldown_dashboard_id": null
					  },
					  "layout": {
						"width": 2,
						"height": 1,
						"row": 1,
						"column": 1
					  },
					  "threshold": {
						"red": 100,
						"yellow": 50
					  }
					}
				  ]
				}
			}
		`))
		if err != nil {
			t.Log(err)
		}
	}))

	dashboardWidgetLayout := DashboardWidgetLayout{
		Width:  2,
		Height: 1,
		Row:    1,
		Column: 1,
	}

	dashboardWidgetPresentation := DashboardWidgetPresentation{
		Title: "95th Percentile Load Time (ms)",
		Notes: "",
	}

	dashboardWidgetData := []DashboardWidgetData{
		{
			NRQL: "SELECT percentile(duration, 95) FROM SyntheticCheck FACET monitorName since 7 days ago",
		},
	}

	dashboardWidgets := []DashboardWidget{
		{
			Visualization: "facet_bar_chart",
			AccountID:     1,
			Data:          dashboardWidgetData,
			Presentation:  dashboardWidgetPresentation,
			Layout:        dashboardWidgetLayout,
		},
	}

	dashboardMetadata := DashboardMetadata{
		Version: 1,
	}

	dashboardFilter := DashboardFilter{
		EventTypes: []string{"test_event"},
		Attributes: []string{"test_attribute1", "test_attribute2"},
	}

	dashboard := Dashboard{
		Title:    "Test",
		Icon:     "bar_chart",
		Widgets:  dashboardWidgets,
		Metadata: dashboardMetadata,
		Filter:   dashboardFilter,
	}

	dashboardResp, err := c.CreateDashboard(dashboard)

	if err != nil {
		t.Log(err)
		t.Fatal("CreateDashboard error")
	}
	if dashboardResp == nil {
		t.Log(err)
		t.Fatal("CreateDashboard error")
	}
	if len(dashboardResp.Filter.Attributes) != 2 {
		t.Fatal("CreateDashboard filter attributes not parsed correctly")
	}
	if len(dashboardResp.Filter.EventTypes) != 1 {
		t.Fatal("CreateDashboard filter event types not parsed correctly")
	}
}
