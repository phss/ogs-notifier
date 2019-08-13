package ogsclient

import (
	"github.com/dghubble/sling"
)

type ogsGameResult struct {
	Next    string `json:"next"`
	Results []Game `json:"results"`
}

// Game represents basic game information.
type Game struct {
	ID int `json:"id"`
}

// MeService provides method for accessing resources under /me endpoint
// of the OGS API
type MeService struct {
	sling *sling.Sling
}

// User fetches the /me resource endpoint, which represents the authenticated user.
func (s *MeService) User() (*User, error) {
	user := new(User)
	apiError := new(ogsAPIError)
	_, err := s.sling.Get("me").Receive(user, apiError)
	return user, handleErrors(err, apiError)
}

// Games fetches all the authenticated user's games.
func (s *MeService) Games() (*[]Game, error) {
	var allGames []Game
	sling := s.sling.New().Path("me/games")

	for {
		games, err := fetchGames(sling)
		if err != nil {
			return nil, err
		}

		allGames = append(allGames, games.Results...)
		sling = sling.New().Base(games.Next)

		if games.Next == "" {
			break
		}
	}
	return &allGames, nil
}

func fetchGames(sling *sling.Sling) (*ogsGameResult, error) {
	games := new(ogsGameResult)
	apiError := new(ogsAPIError)
	_, err := sling.Receive(games, apiError)
	return games, handleErrors(err, apiError)
}
