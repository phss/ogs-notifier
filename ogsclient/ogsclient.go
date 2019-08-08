package ogsclient

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is an OGS client for making API requests.
type Client struct {
	sling *sling.Sling
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
	apiError := new(ogsAPIError)
	_, err := client.sling.Get("me").Receive(me, apiError)
	return me, handleErrors(err, apiError)
}
