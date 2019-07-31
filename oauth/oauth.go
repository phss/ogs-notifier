package oauth

import (
	"context"

	"golang.org/x/oauth2"
)

// Client represents an OAuth Client and also holds a valid Token.
type Client struct {
	Token *oauth2.Token
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

	return &Client{
		Token: token,
	}, nil
}
