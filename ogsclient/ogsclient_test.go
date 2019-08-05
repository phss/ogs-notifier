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

	ogsClient := ogsclient.NewClient(httpClient, server.URL)
	user, err := ogsClient.Me()

	assert.Nil(t, err)
	assert.Equal(t, 12345, user.ID)
}
