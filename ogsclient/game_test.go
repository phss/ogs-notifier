package ogsclient_test

import (
	"testing"

	"github.com/phss/ogs-notifier/ogsclient"
	"github.com/stretchr/testify/assert"
)

func TestGame_HasEnded(t *testing.T) {
	tt := []struct {
		name     string
		game     ogsclient.Game
		expected bool
	}{
		{name: "finished", game: ogsclient.Game{EndedAt: "2019-04-19T00:44:14.896733-04:00"}, expected: true},
		{name: "ongoing", game: ogsclient.Game{EndedAt: ""}, expected: false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.game.HasEnded())
		})
	}
}

func TestGame_HasBlackOrWhiteWon(t *testing.T) {
	tt := []struct {
		name             string
		game             ogsclient.Game
		expectedBlackWon bool
		expectedWhiteWon bool
	}{
		{name: "black won", game: ogsclient.Game{BlackLost: false, WhiteLost: true}, expectedBlackWon: true, expectedWhiteWon: false},
		{name: "white won", game: ogsclient.Game{BlackLost: true, WhiteLost: false}, expectedBlackWon: false, expectedWhiteWon: true},
		{name: "ongoing", game: ogsclient.Game{BlackLost: true, WhiteLost: true}, expectedBlackWon: false, expectedWhiteWon: false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expectedBlackWon, tc.game.HasBlackWon())
			assert.Equal(t, tc.expectedWhiteWon, tc.game.HasWhiteWon())
		})
	}
}
