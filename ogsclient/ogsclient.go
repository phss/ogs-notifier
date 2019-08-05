package ogsclient

import (
	"encoding/json"
	"net/http"
)

// Client is an OGS client for making API requests.
type Client struct {
	httpClient         *http.Client
	ogsAPIBaseEndpoint string
}

// NewClient returns a new OGS client.
func NewClient(httpClient *http.Client, ogsAPIBaseEndpoint string) *Client {
	return &Client{
		httpClient:         httpClient,
		ogsAPIBaseEndpoint: ogsAPIBaseEndpoint,
	}
}

// Me makes a request to the /me resource endpoint.
func (client *Client) Me() (*MeResource, error) {
	resp, err := client.httpClient.Get(client.ogsAPIBaseEndpoint + "/me")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	me := new(MeResource)
	err = json.NewDecoder(resp.Body).Decode(me)
	if err != nil {
		return nil, err
	}

	return me, nil
}
