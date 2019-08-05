package ogsclient

import "net/http"

// Client is an OGS client for making API requests.
type Client struct {
	httpClient *http.Client
	Me         *MeResource
}

// NewClient returns a new OGS client.
func NewClient(httpClient *http.Client, ogsAPIBaseEndpoint string) *Client {
	return &Client{
		httpClient: httpClient,
		Me:         newMeResource(httpClient, ogsAPIBaseEndpoint),
	}
}
