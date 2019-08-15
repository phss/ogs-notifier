package ogsclient_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phss/ogs-notifier/pkg/ogsclient"
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
		{ID: 101, Name: "Friendly Match", Players: ogsclient.Players{Black: ogsclient.User{ID: 23456, Username: "someoneelse", Rating: 0, Ranking: 17}, White: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}}, Outcome: "Resignation", BlackLost: false, WhiteLost: true, StartedAt: "2016-01-02T11:46:43.715348-05:00", EndedAt: "2016-01-02T12:07:38.539597-05:00"},
		{ID: 102, Name: "Friendly Match", Players: ogsclient.Players{Black: ogsclient.User{ID: 101010, Username: "somebot", Rating: 0, Ranking: 16}, White: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}}, Outcome: "33.5 points", BlackLost: false, WhiteLost: true, StartedAt: "2016-01-02T12:14:49.733100-05:00", EndedAt: "2016-01-02T12:15:38.572124-05:00"},
		{ID: 201, Name: "Game with fancy name", Players: ogsclient.Players{Black: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}, White: ogsclient.User{ID: 34512, Username: "another", Rating: 0, Ranking: 14}}, Outcome: "Resignation", BlackLost: false, WhiteLost: true, StartedAt: "2018-12-06T09:50:33.499142-05:00", EndedAt: "2019-02-18T13:29:50.996043-05:00"},
		{ID: 202, Name: "Another game", Players: ogsclient.Players{Black: ogsclient.User{ID: 34512, Username: "another", Rating: 0, Ranking: 14}, White: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}}, Outcome: "Timeout", BlackLost: false, WhiteLost: true, StartedAt: "2019-04-04T18:32:31.029481-04:00", EndedAt: "2019-04-19T00:44:14.896733-04:00"},
		{ID: 203, Name: "Friendly Match", Players: ogsclient.Players{Black: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}, White: ogsclient.User{ID: 101010, Username: "somebot", Rating: 0, Ranking: 16}}, Outcome: "9.5 points", BlackLost: false, WhiteLost: true, StartedAt: "2019-07-11T09:54:19.918904-04:00", EndedAt: "2019-07-11T10:09:11.725619-04:00"},
		{ID: 214, Name: "blah vs. someuser", Players: ogsclient.Players{Black: ogsclient.User{ID: 99999, Username: "blah", Rating: 0, Ranking: 21}, White: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}}, Outcome: "", BlackLost: true, WhiteLost: true, StartedAt: "2019-08-05T17:13:55.375522-04:00", EndedAt: ""},
		{ID: 215, Name: "blah vs. someuser again", Players: ogsclient.Players{Black: ogsclient.User{ID: 12345, Username: "someuser", Rating: 0, Ranking: 19}, White: ogsclient.User{ID: 99999, Username: "blah", Rating: 0, Ranking: 21}}, Outcome: "", BlackLost: true, WhiteLost: true, StartedAt: "2019-08-07T19:51:35.375522-04:00", EndedAt: ""},
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
