package ogsclient

import "net/http"

// MeResource represents the authenticated user.
type MeResource struct {
	httpClient        *http.Client
	ogsAPIEndpointURL string
}

func newMeResource(httpClient *http.Client, ogsAPIBaseEndpoint string) *MeResource {
	return &MeResource{
		httpClient:        httpClient,
		ogsAPIEndpointURL: ogsAPIBaseEndpoint + "/me",
	}
}
