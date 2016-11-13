package api

import "strconv"

// ListAlertConditions returns alert conditions for the specified policy.
func (c *Client) ListAlertConditions(policyID int) ([]AlertCondition, error) {
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

	return resp.Conditions, nil
}
