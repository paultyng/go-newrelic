package api

import (
	"fmt"
	"net/url"
)

// ListAlertPolicies returns all alert policies for the account.
func (c *Client) ListAlertPolicies() ([]AlertPolicy, error) {
	resp := struct {
		Policies []AlertPolicy `json:"policies,omitempty"`
	}{}

	err := c.Do("GET", "/alerts_policies.json", nil, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Policies, nil
}

// CreateAlertPolicy creates a new alert policy for the account.
func (c *Client) CreateAlertPolicy(policy AlertPolicy) error {
	req := struct {
		Policy AlertPolicy `json:"policy"`
	}{
		Policy: policy,
	}

	return c.Do("POST", "/alerts_policies.json", nil, req, nil)
}

// DeleteAlertPolicy deletes an existing alert policy from the account.
func (c *Client) DeleteAlertPolicy(id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_policies/%v.json", id)}
	return c.Do("DELETE", u.String(), nil, nil, nil)
}
