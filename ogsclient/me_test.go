package ogsclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/ogs-notifier/ogsclient"
	"github.com/stretchr/testify/assert"
)

func TestMeResource_displayRank(t *testing.T) {
	tt := []struct {
		name                   string
		ranking                int
		expectedDisplayRanking string
	}{
		{name: "total beginner", ranking: 0, expectedDisplayRanking: "30 Kyu"},
		{name: "ddk", ranking: 18, expectedDisplayRanking: "12 Kyu"},
		{name: "sdk", ranking: 23, expectedDisplayRanking: "7 Kyu"},
		{name: "almost dan", ranking: 29, expectedDisplayRanking: "1 Kyu"},
		{name: "shodan", ranking: 30, expectedDisplayRanking: "1 Dan"},
		{name: "strong dan", ranking: 36, expectedDisplayRanking: "7 Dan"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			user := &ogsclient.User{
				ID:       12345,
				Username: "someuser",
				Rating:   123.456,
				Ranking:  tc.ranking,
			}

			rank := user.DisplayRanking()

			assert.Equal(t, tc.expectedDisplayRanking, rank)
		})
	}
}

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
