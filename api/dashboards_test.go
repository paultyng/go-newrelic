package api

import (
	"net/http"
	"testing"
)

func TestDashboards_Basic(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
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
		w.Write([]byte(`
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
				  "filter": null,
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
					  }
					}
				  ]
				}
			}
    `))
	}))

	dashboard, err := c.GetDashboard(1234)
	if err != nil {
		t.Log(err)
		t.Fatal("getDashboard error")
	}

	if len(dashboard.Widgets) == 0 {
		t.Fatal("Dashboard widgets found")
	}
}

func TestGetDashboard_WithMetricsWidget(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
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
				  "filter": null,
				  "widgets": [
          {
            "visualization": "metrics_line_chart",
            "account_id": 1,
            "data": [
                {
                  "duration": 1800000,
                  "end_time": null,
                  "entity_ids": [
                      234699717
                  ],
                  "compare_with": null,
                  "metrics": [
                      {
                          "name": "metric_name",
                          "units": null,
                          "scope": "",
                          "values": [
                              "average_value"
                          ]
                      }
                  ],
                  "raw_metric_name": "raw_metric_name",
                  "facet": null,
                  "order_by": "average_value",
                  "limit": 10
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
            }
          }
				  ]
				}
			}
    `))
	}))

	dashboard, err := c.GetDashboard(1234)
	if err != nil {
		t.Log(err)
		t.Fatal("getDashboard error")
	}

	if len(dashboard.Widgets) == 0 {
		t.Fatal("Dashboard widgets found")
	}
}

func TestCreateDashboardCondition_WithMetricsWidget(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
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
				  "filter": null,
				  "widgets": [
          {
            "visualization": "metrics_line_chart",
            "account_id": 1,
            "data": [
                {
                    "duration": 1800000,
                    "end_time": null,
                    "entity_ids": [
                        234699717
                    ],
                    "compare_with": null,
                    "metrics": [
                        {
                            "name": "metric_name",
                            "units": null,
                            "scope": "",
                            "values": [
                                "average_value"
                            ]
                        }
                    ],
                    "raw_metric_name": "raw_metric_name",
                    "facet": null,
                    "order_by": "average_value",
                    "limit": 10
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
            }
          }
				  ]
				}
			}
		`))
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

	dashboardWidgetDataMetrics := []DashboardWidgetDataMetric{
	  {
		  Name:   "metric_name",
		  Values: []string{"average_value"},
	  },
	}

	dashboardWidgetData := []DashboardWidgetData{
	  {
			Duration:      1800000,
			EntityIds:     []int{234699717},
			Metrics:       dashboardWidgetDataMetrics,
			RawMetricName: "raw_metric_name",
			OrderBy:       "average_value",
			Limit:         10,
		},
	}

	dashboardWidgets := []DashboardWidget{
		{
			Visualization: "metrics_line_chart",
			AccountID:     1,
			Data:          dashboardWidgetData,
			Presentation:  dashboardWidgetPresentation,
			Layout:        dashboardWidgetLayout,
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
	if dashboard.Metadata.Version != 1 {
		t.Fatal("CreateDashboard metadata version incorrect")
	}
	if dashboardResp.ID != 1234 {
		t.Fatal("CreateDashboard ID was not parsed correctly")
	}
}

func TestCreateDashboardCondition(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
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
				  "filter": null,
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
					  }
					}
				  ]
				}
			}
		`))
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
	if dashboard.Metadata.Version != 1 {
		t.Fatal("CreateDashboard metadata version incorrect")
	}
	if dashboardResp.ID != 1234 {
		t.Fatal("CreateDashboard ID was not parsed correctly")
	}
}

func TestCreateDashboardWithFilter(t *testing.T) {
	c := newTestAPIClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
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
					  }
					}
				  ]
				}
			}
		`))
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
