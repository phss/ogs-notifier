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
	passwordConfig := Config{
		TokenURL:     config.TokenURL,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Username:     username,
		Password:     password,
	}
	return NewClient(passwordConfig)
}

// RefreshTokenClient create a new OAuth Client using refresh token based authentication.
func RefreshTokenClient(config Config, refreshToken string) (*Client, error) {
	passwordConfig := Config{
		TokenURL:     config.TokenURL,
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RefreshToken: refreshToken,
	}
	return NewClient(passwordConfig)
}

// NewClient create a new OAuth Client using password based authentication.
func NewClient(config Config) (*Client, error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       []string{"read"},
		Endpoint: oauth2.Endpoint{
			TokenURL: config.TokenURL,
		},
	}

	var token *oauth2.Token
	var err error

	if config.Password != "" {
		token, err = conf.PasswordCredentialsToken(ctx, config.Username, config.Password)
		if err != nil {
			return nil, err
		}
	} else {
		previousToken := &oauth2.Token{
			RefreshToken: config.RefreshToken,
		}
		token, err = conf.TokenSource(ctx, previousToken).Token()
		if err != nil {
			return nil, err
		}
	}

	httpClient := conf.Client(ctx, token)

	return &Client{
		HTTPClient: httpClient,
		Token:      token,
	}, nil
}
