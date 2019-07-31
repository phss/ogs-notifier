package oauth

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

// Client represents an OAuth client, holding an authenticated HTTP client
// and a valid Token.
type Client struct {
	HTTPClient *http.Client
	Token      *oauth2.Token
}

// NewClient create a new OAuth Client using password based authentication.
func NewClient(tokenURL string, clientID string, clientSecret string, username string, password string) (*Client, error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"read"},
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
		},
	}

	token, err := conf.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return nil, err
	}
	httpClient := conf.Client(ctx, token)

	return &Client{
		HTTPClient: httpClient,
		Token:      token,
	}, nil
}
