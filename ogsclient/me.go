package ogsclient

// MeResource represents an authenticated user resource representation
type MeResource struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Rating   float32 `json:"rating"`
	Ranking  int     `json:"ranking"`
}
