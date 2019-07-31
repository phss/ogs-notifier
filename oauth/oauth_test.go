package oauth_test

import (
	"io/ioutil"
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
		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "/tokens-endpoint", r.URL.String())
		assert.Equal(t, "grant_type=password&password=somepassword&scope=read&username=someusername", string(body))

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
	assert.NotNil(t, client.HTTPClient)
	assert.Equal(t, expectedToken.AccessToken, client.Token.AccessToken)
	assert.Equal(t, expectedToken.RefreshToken, client.Token.RefreshToken)
}

func TestNewClient_handleFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := oauth.NewClient(server.URL+"/tokens-endpoint", "someclientid", "someclientsecret", "someusername", "somepassword")

	assert.Error(t, err)
}
