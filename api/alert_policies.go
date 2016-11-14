package api

import (
	"fmt"
	"net/url"
)

func (c *Client) queryAlertPolicies(name *string) ([]AlertPolicy, error) {
	qs := map[string]string{}

	if name != nil {
		qs["filter[name]"] = *name
	}

	resp := struct {
		Policies []AlertPolicy `json:"policies,omitempty"`
	}{}

	err := c.Do("GET", "/alerts_policies.json", qs, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Policies, nil
}

// GetAlertPolicy returns a specific alert policy by ID
func (c *Client) GetAlertPolicy(id int) (*AlertPolicy, error) {
	policies, err := c.queryAlertPolicies(nil)
	if err != nil {
		return nil, err
	}

	for _, policy := range policies {
		if policy.ID == id {
			return &policy, nil
		}
	}

	return nil, fmt.Errorf("No matching alert policy found.")
}

// ListAlertPolicies returns all alert policies for the account.
func (c *Client) ListAlertPolicies() ([]AlertPolicy, error) {
	return c.queryAlertPolicies(nil)
}

// CreateAlertPolicy creates a new alert policy for the account.
func (c *Client) CreateAlertPolicy(policy AlertPolicy) (*AlertPolicy, error) {
	req := struct {
		Policy AlertPolicy `json:"policy"`
	}{
		Policy: policy,
	}

	resp := struct {
		Policy AlertPolicy `json:"policy,omitempty"`
	}{}

	err := c.Do("POST", "/alerts_policies.json", nil, req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Policy, nil
}

// DeleteAlertPolicy deletes an existing alert policy from the account.
func (c *Client) DeleteAlertPolicy(id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_policies/%v.json", id)}
	return c.Do("DELETE", u.String(), nil, nil, nil)
}
