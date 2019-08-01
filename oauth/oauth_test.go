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

func TestNewClient_passwordBased(t *testing.T) {
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

	config := oauth.Config{
		TokenURL:     server.URL + "/tokens-endpoint",
		ClientID:     "someclientid",
		ClientSecret: "someclientsecret",
		Username:     "someusername",
		Password:     "somepassword",
	}
	client, err := oauth.NewClient(config)

	assert.Nil(t, err)
	assert.NotNil(t, client.HTTPClient)
	assert.Equal(t, expectedToken.AccessToken, client.Token.AccessToken)
	assert.Equal(t, expectedToken.RefreshToken, client.Token.RefreshToken)
}

func TestNewClient_refreshTokenBased(t *testing.T) {
	expectedToken := oauth2.Token{
		TokenType:    "Bearer",
		AccessToken:  "someaccesstoken",
		RefreshToken: "somerefreshtoken",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		assert.Equal(t, "/tokens-endpoint", r.URL.String())
		assert.Equal(t, "grant_type=refresh_token&refresh_token=sometoken", string(body))

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

	config := oauth.Config{
		TokenURL:     server.URL + "/tokens-endpoint",
		ClientID:     "someclientid",
		ClientSecret: "someclientsecret",
		Username:     "someusername",
		RefreshToken: "sometoken",
	}
	client, err := oauth.NewClient(config)

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

	config := oauth.Config{
		TokenURL:     server.URL + "/tokens-endpoint",
		ClientID:     "someclientid",
		ClientSecret: "someclientsecret",
		Username:     "someusername",
		Password:     "somepassword",
	}
	_, err := oauth.NewClient(config)

	assert.Error(t, err)
}
