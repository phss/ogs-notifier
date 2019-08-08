package ogsclient_test

import (
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
			user := &ogsclient.MeResource{
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
