package api

import (
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) queryAlertConditions(policyID int) ([]AlertCondition, error) {
	qs := map[string]string{
		"policy_id": strconv.Itoa(policyID),
	}

	resp := struct {
		Conditions []AlertCondition `json:"conditions,omitempty"`
	}{}

	err := c.Do("GET", "/alerts_conditions.json", qs, nil, &resp)
	if err != nil {
		return nil, err
	}

	for _, c := range resp.Conditions {
		c.PolicyID = policyID
	}

	return resp.Conditions, nil
}

func (c *Client) GetAlertCondition(policyID int, id int) (*AlertCondition, error) {
	conditions, err := c.queryAlertConditions(policyID)
	if err != nil {
		return nil, err
	}

	for _, condition := range conditions {
		if condition.ID == id {
			return &condition, nil
		}
	}

	return nil, newNotFoundError()
}

// ListAlertConditions returns alert conditions for the specified policy.
func (c *Client) ListAlertConditions(policyID int) ([]AlertCondition, error) {
	return c.queryAlertConditions(policyID)
}

func (c *Client) CreateAlertCondition(condition AlertCondition) (*AlertCondition, error) {
	policyID := condition.PolicyID

	req := struct {
		Condition AlertCondition `json:"condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertCondition `json:"condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_conditions/policies/%v.json", policyID)}
	err := c.Do("POST", u.String(), nil, req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

func (c *Client) UpdateAlertCondition(condition AlertCondition) (*AlertCondition, error) {
	policyID := condition.PolicyID
	id := condition.ID

	req := struct {
		Condition AlertCondition `json:"condition"`
	}{
		Condition: condition,
	}

	resp := struct {
		Condition AlertCondition `json:"condition,omitempty"`
	}{}

	u := &url.URL{Path: fmt.Sprintf("/alerts_conditions/%v.json", id)}
	err := c.Do("PUT", u.String(), nil, req, &resp)
	if err != nil {
		return nil, err
	}

	resp.Condition.PolicyID = policyID

	return &resp.Condition, nil
}

func (c *Client) DeleteAlertCondition(policyID int, id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_conditions/%v.json", id)}
	return c.Do("DELETE", u.String(), nil, nil, nil)
}
