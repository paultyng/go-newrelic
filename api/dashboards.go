package api

import (
	"fmt"
	"net/url"
)

// ListDashboards allows you to get all dashboards.
func (c *Client) ListDashboards() ([]ListDashboardSummary, error) {
	var response ListDashboardResp

	reqURL, err := url.Parse("/dashboards.json")
	if err != nil {
		return nil, err
	}

	nextPath := reqURL.String()

	nextPath, err = c.Do("GET", nextPath, nil, &response)
	if err != nil {
		return nil, err
	}

	return response.Dashboards, nil
}

// GetDashboard allows you to get a specific dashboard by ID.
func (c *Client) GetDashboard(id int) (GetDashboardDetail, error) {
	var response GetDashboardDetail

	path := fmt.Sprintf("/dashboards/%v.json", id)
	reqURL, err := url.Parse(path)
	if err != nil {
		return GetDashboardDetail{}, err
	}

	nextPath := reqURL.String()

	// u := &url.URL{Path: fmt.Sprintf("/dashboards/%v.json", id)}
	nextPath, err = c.Do("GET", nextPath, nil, &response)
	if err != nil {
		return GetDashboardDetail{}, err
	}

	return response, nil
}

// CreateDashboard allows you to create a dashboard with the specified info and widgets.
func (c *Client) CreateDashboard(dashboardObj GetDashboardResp) (*GetDashboardResp, error) {
	var resp GetDashboardResp

	_, err := c.Do("POST", "/dashboards.json", dashboardObj, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// UpdateDashboard allows you to update an existing dashboard by ID.
func (c *Client) UpdateDashboard(id int, dashboard GetDashboardResp) (*GetDashboardResp, error) {
	var resp GetDashboardResp

	u := &url.URL{Path: fmt.Sprintf("/dashboards/%v.json", id)}
	_, err := c.Do("PUT", u.String(), dashboard, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteDashboard deletes the dashboard with the specified ID.
func (c *Client) DeleteDashboard(id int) error {
	u := &url.URL{Path: fmt.Sprintf("/dashboards/%v.json", id)}
	_, err := c.Do("DELETE", u.String(), nil, nil)
	return err
}
