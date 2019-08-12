package ogsclient

import (
	"github.com/dghubble/sling"
)

type gameResult struct {
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
	allGames := make([]Game, 0)
	sling := s.sling.New().Path("me/games")

	for {
		games := new(gameResult)
		apiError := new(ogsAPIError)
		_, err := sling.Get("").Receive(games, apiError)

		if err != nil || apiError.Detail != "" {
			return nil, handleErrors(err, apiError)
		}

		for _, game := range games.Results {
			allGames = append(allGames, game)
		}

		if games.Next == "" {
			break
		}
		sling = sling.New().Base(games.Next)
	}
	return &allGames, nil
}
