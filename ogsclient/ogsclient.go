package ogsclient

import "net/http"

// Client is an OGS client for making API requests.
type Client struct {
	HttpClient *http.Client
}

// NewClient returns a new OGS client.
func NewClient(httpClient *http.Client, ogsAPIEndpoint string) *Client {
	return &Client{
		HttpClient: httpClient,
	}
}
