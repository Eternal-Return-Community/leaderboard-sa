package structs

type ER_PlayerInfo struct {
	User struct {
		UserNum int `json:"userNum"`
	}
}

type Ranked struct {
	Code     int `json:"code"`
	UserRank struct {
		UserNum  int    `json:"userNum"`
		Mmr      int    `json:"mmr"`
		Nickname string `json:"nickname"`
		Rank     int    `json:"rank"`
	}
}

type RankedInfo struct {
	Code     int    `json:"code"`
	Mmr      int    `json:"mmr"`
	Elo      string `json:"elo"`
	Rank     int    `json:"rank"`
	Nickname string `json:"nickname"`
}
