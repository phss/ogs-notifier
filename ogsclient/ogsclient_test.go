package ogsclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/ogs-notifier/ogsclient"
	"github.com/stretchr/testify/assert"
)

func TestNewClient_userDetails(t *testing.T) {
	httpClient := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "/me", r.URL.String())

		response := `{
			"id": 12345,
			"username": "someuser",
			"rating": 948.565,
			"rating_blitz": 1150,
			"rating_live": 850.063,
			"rating_correspondence": 1088.131,
			"ranking": 18,
			"ranking_blitz": 20,
			"ranking_live": 17,
			"ranking_correspondence": 19,
			"about": "",
			"settings": "/api/v1/mesettings",
			"friends": "/api/v1/mefriends",
			"games": "/api/v1/megames",
			"challenges": "/api/v1/mechallenges",
			"groups": "/api/v1/megroups",
			"mail": "/api/v1/memail",
			"tournaments": "/api/v1/metournaments",
			"vacation": "/api/v1/mevacation",
			"notifications": "/api/v1/menotifications"
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	expectedUser := &ogsclient.MeResource{
		ID:        12345,
		Username:  "someuser",
		Rating:    948.565,
		Ranking:   18,
		GamesPath: "/api/v1/megames",
	}
	ogsClient := ogsclient.NewClient(httpClient, server.URL)
	user, err := ogsClient.Me()

	assert.Nil(t, err)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Username, user.Username)
	assert.Equal(t, expectedUser.Rating, user.Rating)
	assert.Equal(t, expectedUser.Ranking, user.Ranking)
	assert.Equal(t, expectedUser.GamesPath, user.GamesPath)
}

func TestNewClient_userDetailsFailure(t *testing.T) {
	httpClient := &http.Client{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := `{
			"detail": "Something went wrong"
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(response))
	}))

	ogsClient := ogsclient.NewClient(httpClient, server.URL)
	_, err := ogsClient.Me()

	assert.EqualError(t, err, "Something went wrong")
}

func TestNewClient_badUrl(t *testing.T) {
	httpClient := &http.Client{}

	ogsClient := ogsclient.NewClient(httpClient, "badurl")
	_, err := ogsClient.Me()

	assert.Error(t, err)
}
