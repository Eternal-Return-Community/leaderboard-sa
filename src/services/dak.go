package services

import (
	"encoding/json"
	"erbs/src/structs"
	"erbs/src/utils"
	"fmt"
	"net/http"
)

const (
	url = "https://er.dakgg.io/v1/leaderboard?page=1&seasonKey=PRE_SEASON_10&serverName=Sao+Paulo&teamMode=SQUAD"
)

func Dak() []structs.PlayerInfo {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var leader structs.Leaderboard
	err = json.NewDecoder(resp.Body).Decode(&leader)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	size := min(len(leader.Leaderboards), 10)
	players := make([]structs.PlayerInfo, size)

	for i := 0; i < size; i++ {
		players[i] = structs.PlayerInfo{
			Ranking: i + 1,
			Player:  leader.Leaderboards[i].Nickname,
			Elo:     utils.CalcElo(leader.Leaderboards[i].Mmr),
		}
	}

	return players
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
