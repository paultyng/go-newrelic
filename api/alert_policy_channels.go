package api

import (
	"regexp"
	"strconv"
	"strings"
)

func (c *Client) UpdateAlertPolicyChannels(policyID int, channelIDs []int) error {
	channelIDStrings := make([]string, len(channelIDs))

	for i, channelID := range channelIDs {
		channelIDStrings[i] = strconv.Itoa(channelID)
	}

	qs := map[string]string{
		"policy_id":   strconv.Itoa(policyID),
		"channel_ids": strings.Join(channelIDStrings, ","),
	}

	return c.Do("PUT", "/alerts_policy_channels.json", qs, nil, nil)
}

func (c *Client) DeleteAlertPolicyChannel(policyID int, channelID int) error {
	qs := map[string]string{
		"policy_id":  strconv.Itoa(policyID),
		"channel_id": strconv.Itoa(channelID),
	}

	err := c.Do("DELETE", "/alerts_policy_channels.json", qs, nil, nil)
	if err != nil {
		if apiErr, ok := err.(*ErrorResponse); ok {
			matched, err := regexp.MatchString("Alerts policy with ID: \\d+ is not valid.", apiErr.Detail.Title)
			if err != nil {
				return err
			}

			if matched {
				return ErrNotFound
			}
		}

		return err
	}

	return nil
}
