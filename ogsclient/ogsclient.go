package ogsclient

import (
	"errors"
	"net/http"

	"github.com/dghubble/sling"
)

// Client is an OGS client for making API requests.
type Client struct {
	sling *sling.Sling
}

type OgsAPIError struct {
	Detail string `json:"detail"`
}

// NewClient returns a new OGS client.
func NewClient(httpClient *http.Client, ogsAPIBaseEndpoint string) *Client {
	base := sling.New().Client(httpClient).Base(ogsAPIBaseEndpoint)
	return &Client{
		sling: base,
	}
}

// Me makes a request to the /me resource endpoint.
func (client *Client) Me() (*MeResource, error) {
	me := new(MeResource)
	apiError := new(OgsAPIError)

	_, err := client.sling.Get("me").Receive(me, apiError)
	if err != nil {
		return nil, err
	}

	if apiError.Detail != "" {
		return nil, errors.New(apiError.Detail)
	}

	return me, nil
}
