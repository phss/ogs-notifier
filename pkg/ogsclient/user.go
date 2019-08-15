package ogsclient

import "fmt"

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
