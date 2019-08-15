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

// Config describes the required information to perform an OAuth password
// based authentication.
type Config struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	RefreshToken string
}

// PasswordCredentialsClient create a new OAuth Client using password based authentication.
func PasswordCredentialsClient(config Config, username string, password string) (*Client, error) {
	ctx := context.Background()
	conf := oauthConfig(config)

	token, err := conf.PasswordCredentialsToken(ctx, username, password)
	if err != nil {
		return nil, err
	}

	return newClient(ctx, conf, token), nil
}

// RefreshTokenClient create a new OAuth Client using refresh token based authentication.
func RefreshTokenClient(config Config, refreshToken string) (*Client, error) {
	ctx := context.Background()
	conf := oauthConfig(config)

	previousToken := &oauth2.Token{
		RefreshToken: refreshToken,
	}
	token, err := conf.TokenSource(ctx, previousToken).Token()
	if err != nil {
		return nil, err
	}

	return newClient(ctx, conf, token), nil
}

func oauthConfig(config Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read"},
		Endpoint: oauth2.Endpoint{
			TokenURL: config.TokenURL,
		},
	}
}

func newClient(ctx context.Context, conf *oauth2.Config, token *oauth2.Token) *Client {
	httpClient := conf.Client(ctx, token)
	return &Client{
		HTTPClient: httpClient,
		Token:      token,
	}
}
