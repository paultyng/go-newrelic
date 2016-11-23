package api

import (
	"fmt"
	"net/url"
)

func (c *Client) queryLabels() ([]Label, error) {
	resp := struct {
		Labels []Label `json:"labels,omitempty"`
	}{}

	err := c.Do("GET", "/labels.json", nil, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Labels, nil
}

func (c *Client) GetLabel(key string) (*Label, error) {
	labels, err := c.queryLabels()
	if err != nil {
		return nil, err
	}

	for _, label := range labels {
		if label.Key == key {
			return &label, nil
		}
	}

	return nil, ErrNotFound
}

// ListLabels returns the labels for the account.
func (c *Client) ListLabels() ([]Label, error) {
	return c.queryLabels()
}

// CreateLabel creates a new label for the account.
func (c *Client) CreateLabel(label Label) error {
	if label.Links.Applications == nil {
		label.Links.Applications = make([]int, 0)
	}

	if label.Links.Servers == nil {
		label.Links.Servers = make([]int, 0)
	}

	req := struct {
		Label Label `json:"label,omitempty"`
	}{
		Label: label,
	}

	return c.Do("PUT", "/labels.json", nil, req, nil)
}

// DeleteLabel deletes a label on the account specified by key.
func (c *Client) DeleteLabel(key string) error {
	u := &url.URL{Path: fmt.Sprintf("/labels/%v.json", key)}
	return c.Do("DELETE", u.String(), nil, nil, nil)
}
