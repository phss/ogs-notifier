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
		{name: "finished", game: ogsclient.Game{EndedAt: ""}, expected: false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.game.HasEnded())
		})
	}
}
