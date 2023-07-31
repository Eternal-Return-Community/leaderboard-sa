package structs

type Leaderboard struct {
	Leaderboards []struct {
		Nickname string `json:"nickname"`
		Mmr      int    `json:"mmr"`
	}
}

type PlayerInfo struct {
	Ranking int
	Player  string
	Elo     string
}
