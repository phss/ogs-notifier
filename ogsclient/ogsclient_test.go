package ogsclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/ogs-notifier/ogsclient"
	"github.com/stretchr/testify/assert"
)

func TestNewClient_userDetails(t *testing.T) {
	server := fakeOgsServer(t)
	expectedUser := &ogsclient.MeResource{
		ID:        12345,
		Username:  "someuser",
		Rating:    948.565,
		Ranking:   18,
		GamesPath: "/api/v1/megames",
	}
	ogsClient := ogsclient.NewClient(http.DefaultClient, server.URL)
	user, err := ogsClient.Me()

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Rating, user.Rating)
	assert.Equal(t, expectedUser.Ranking, user.Ranking)
	assert.Equal(t, expectedUser.GamesPath, user.GamesPath)
}

func TestNewClient_userDetailsFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"detail": "Something went wrong"
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
	}))

	ogsClient := ogsclient.NewClient(http.DefaultClient, server.URL)
	_, err := ogsClient.Me()

	assert.EqualError(t, err, "Something went wrong")
}

func TestNewClient_badUrl(t *testing.T) {
	httpClient := &http.Client{}

	ogsClient := ogsclient.NewClient(httpClient, "badurl")
	_, err := ogsClient.Me()

	assert.Error(t, err)
}
