package ogsclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/ogs-notifier/ogsclient"
	"github.com/stretchr/testify/assert"
)

func TestMeService_User(t *testing.T) {
	server := fakeOgsServer(t)
	expectedUser := &ogsclient.User{
		ID:       12345,
		Username: "someuser",
		Rating:   948.565,
		Ranking:  18,
	}
	ogsClient := ogsclient.NewClient(http.DefaultClient, server.URL)
	user, err := ogsClient.Me.User()

	assert.Nil(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestMeService_User_apiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"detail": "Something went wrong"
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
	}))

	ogsClient := ogsclient.NewClient(http.DefaultClient, server.URL)
	_, err := ogsClient.Me.User()

	assert.EqualError(t, err, "Something went wrong")
}

func TestMeService_User_badUrl(t *testing.T) {
	httpClient := &http.Client{}

	ogsClient := ogsclient.NewClient(httpClient, "badurl")
	_, err := ogsClient.Me.User()

	assert.Error(t, err)
}

func TestMeServer_Games(t *testing.T) {
	server := fakeOgsServer(t)
	expectedGames := &[]ogsclient.Game{
		{ID: 101},
		{ID: 102},
		{ID: 201},
		{ID: 202},
		{ID: 203},
		{ID: 214},
		{ID: 215},
	}

	ogsClient := ogsclient.NewClient(http.DefaultClient, server.URL)
	games, err := ogsClient.Me.Games()

	assert.Nil(t, err)
	assert.Equal(t, expectedGames, games)
}

func TestMeServer_Games_apiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"detail": "Something went wrong"
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
	}))

	ogsClient := ogsclient.NewClient(http.DefaultClient, server.URL)
	_, err := ogsClient.Me.Games()

	assert.EqualError(t, err, "Something went wrong")
}
