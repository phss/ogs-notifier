package ogsclient

import (
	"fmt"

	"github.com/dghubble/sling"
)

// User represents an authenticated user resource
type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Rating   float32 `json:"rating"`
	Ranking  int     `json:"ranking"`
}

// DisplayRanking converts Ranking into the usual Kyu/Dan ranking.
func (user User) DisplayRanking() string {
	switch {
	case user.Ranking >= 30:
		return fmt.Sprintf("%d Dan", user.Ranking-29)
	default:
		return fmt.Sprintf("%d Kyu", 30-user.Ranking)
	}
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
