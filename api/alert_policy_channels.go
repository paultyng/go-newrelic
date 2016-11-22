package api

import (
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

func (c *Client) DeleteAlertPolicyChannels(policyID int, channelID int) error {
	qs := map[string]string{
		"policy_id":  strconv.Itoa(policyID),
		"channel_id": strconv.Itoa(channelID),
	}

	return c.Do("DELETE", "/alerts_policy_channels.json", qs, nil, nil)
}
