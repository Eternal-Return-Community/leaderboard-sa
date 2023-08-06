package structs

type ServerRanking struct {
	PageProps PageProps `json:"pageProps"`
}

type PageProps struct {
	Data Data `json:"data"`
}

type Data struct {
	TeamModeSummary []TeamModeSummary `json:"teamModeSummary"`
}

type TeamModeSummary struct {
	TeamMode struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"teamMode"`

	PlayerTier struct {
		MMR       int    `json:"mmr"`
		TierType  int    `json:"tierType"`
		TierGrade int    `json:"tierGrade"`
		LP        int    `json:"lp"`
		Name      string `json:"name"`
		SeasonID  int    `json:"seasonId"`
		ImageURL  string `json:"imageUrl"`
	} `json:"playerTier"`

	GlobalRanking struct {
		Rank int `json:"rank"`
	} `json:"globalRanking"`

	ServerRanking struct {
		Rank       int     `json:"rank"`
		RankRate   float64 `json:"rankRate"`
		ServerName string  `json:"serverName"`
	} `json:"serverRanking"`
}
