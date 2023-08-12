package structs

type Data struct {
	PlayerTierByUserNum PlayerTierData    `json:"playerTierByUserNum"`
	Leaderboards        []LeaderboardInfo `json:"leaderboards"`
}

type LeaderboardInfo struct {
	UserNum  int    `json:"usernum"`
	Nickname string `json:"nickname"`
	Mmr      int    `json:"mmr"`
}

type PlayerTierData map[int]PlayerTierInfo
type PlayerTierInfo struct {
	Mmr       int    `json:"mmr"`
	TierType  int    `json:"tierType"`
	TierGrade int    `json:"tierGrade"`
	Lp        int    `json:"lp"`
	Name      string `json:"name"`
	SeasonId  int    `json:"seasonId"`
	ImageUrl  string `json:"imageUrl"`
}

type PlayerInfo struct {
	UserNum int
	Ranking int
	Player  string
	Elo     string
}
