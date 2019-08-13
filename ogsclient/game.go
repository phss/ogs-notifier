package ogsclient

type ogsGameResult struct {
	Next    string `json:"next"`
	Results []Game `json:"results"`
}

// Game represents basic game information.
type Game struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Players   Players `json:"players"`
	Outcome   string  `json:"outcome"`
	BlackLost bool    `json:"black_lost"`
	WhiteLost bool    `json:"white_lost"`
	StartedAt string  `json:"started"`
	EndedAt   string  `json:"ended"`
}

// Players holds both black and white users.
type Players struct {
	Black User `json:"black"`
	White User `json:"white"`
}

// HasEnded return true if the game has finished.
func (game *Game) HasEnded() bool {
	return game.EndedAt != ""
}
