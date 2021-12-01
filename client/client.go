package client

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	Config     Config
	HttpClient http.Client
}

func (c *Client) Exec(req http.Request, v interface{}) error {
	req.Header.Set("Authorization", c.Config.ApiToken)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	resp, err := c.HttpClient.Do(&req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return err
	}

	return nil
}
