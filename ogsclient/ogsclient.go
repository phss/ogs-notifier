package ogsclient

import (
	"errors"
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
	return nil, errors.New("not implemented")
}
