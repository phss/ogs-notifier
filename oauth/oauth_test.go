package oauth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/ogs-notifier/oauth"
	"github.com/stretchr/testify/assert"
	"golang.org/x/oauth2"
)

func TestNewClient(t *testing.T) {
	expectedToken := oauth2.Token{
		TokenType:    "Bearer",
		AccessToken:  "someaccesstoken",
		RefreshToken: "somerefreshtoken",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"access_token": "someaccesstoken",
			"expires_in": 36000,
			"token_type": "Bearer",
			"scope": "read",
			"refresh_token": "somerefreshtoken"
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client, err := oauth.NewClient(server.URL+"/tokens-endpoint", "someclientid", "someclientsecret", "someusername", "somepassword")

	assert.Nil(t, err)
	assert.Equal(t, expectedToken.AccessToken, client.Token.AccessToken)
	assert.Equal(t, expectedToken.RefreshToken, client.Token.RefreshToken)
}
