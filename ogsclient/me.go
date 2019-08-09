package ogsclient

import "fmt"

// MeResource represents an authenticated user resource with links to further actions
type MeResource struct {
	ID        int     `json:"id"`
	Username  string  `json:"username"`
	Rating    float32 `json:"rating"`
	Ranking   int     `json:"ranking"`
	GamesPath string  `json:"games"`
}

// DisplayRanking converts Ranking into the usual Kyu/Dan ranking.
func (user MeResource) DisplayRanking() string {
	switch {
	case user.Ranking >= 30:
		return fmt.Sprintf("%d Dan", user.Ranking-29)
	default:
		return fmt.Sprintf("%d Kyu", 30-user.Ranking)
	}
}
