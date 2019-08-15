package ogsclient

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is an OGS client for making OGS API requests.
type Client struct {
	sling *sling.Sling
	Me    *MeService
}

// NewClient returns a new OGS client.
func NewClient(httpClient *http.Client, ogsAPIBaseEndpoint string) *Client {
	base := sling.New().Client(httpClient).Base(ogsAPIBaseEndpoint)
	return &Client{
		sling: base,
		Me:    &MeService{base.New()},
	}
}
