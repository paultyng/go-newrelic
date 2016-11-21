package api

import (
	"fmt"
	"net/url"
)

func (c *Client) queryAlertChannels() ([]AlertChannel, error) {
	resp := struct {
		Channels []AlertChannel `json:"channels,omitempty"`
	}{}

	err := c.Do("GET", "/alerts_channels.json", nil, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Channels, nil
}

// GetAlertChannel returns a specific alert channel by ID
func (c *Client) GetAlertChannel(id int) (*AlertChannel, error) {
	channels, err := c.queryAlertChannels()
	if err != nil {
		return nil, err
	}

	for _, channel := range channels {
		if channel.ID == id {
			return &channel, nil
		}
	}

	return nil, newNotFoundError()
}

// ListAlertChannels returns all alert policies for the account.
func (c *Client) ListAlertChannels() ([]AlertChannel, error) {
	return c.queryAlertChannels()
}

func (c *Client) CreateAlertChannel(channel AlertChannel) (*AlertChannel, error) {
	// TODO: support attaching policy ID's here?
	// qs := map[string]string{
	// 	"policy_ids[]": channel.Links.PolicyIDs,
	// }

	if len(channel.Links.PolicyIDs) > 0 {
		return nil, fmt.Errorf("You cannot create an alert channel with policy IDs, you must attach polidy IDs after creation.")
	}

	req := struct {
		Channel AlertChannel `json:"channel"`
	}{
		Channel: channel,
	}

	resp := struct {
		Channel AlertChannel `json:"channel,omitempty"`
	}{}

	err := c.Do("POST", "/alerts_channels.json", nil, req, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Channel, nil
}

func (c *Client) DeleteAlertChannel(id int) error {
	u := &url.URL{Path: fmt.Sprintf("/alerts_channels/%v.json", id)}
	return c.Do("DELETE", u.String(), nil, nil, nil)
}
