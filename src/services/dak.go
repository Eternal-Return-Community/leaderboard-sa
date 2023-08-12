package services

import (
	"encoding/json"
	"erbs/src/structs"
	"erbs/src/utils"
	"fmt"
	"net/http"
)

var playerEloInfo map[int]structs.PlayerTierInfo

const (
	url = "https://er.dakgg.io/v1/leaderboard?page=1&seasonKey=SEASON_10&serverName=Sao+Paulo&teamMode=SQUAD"
)

func Dak() []structs.PlayerInfo {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()
	var leader structs.Data
	err = json.NewDecoder(resp.Body).Decode(&leader)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(leader.Leaderboards)

	size := min(len(leader.Leaderboards), 10)
	players := make([]structs.PlayerInfo, size)

	for i := 0; i < size; i++ {
		userNum := leader.Leaderboards[i].UserNum
		elo, found := leader.PlayerTierByUserNum[userNum]
		if !found {
			fmt.Printf("User with UserNum %d not found\n", userNum)
			continue
		}

		players[i] = structs.PlayerInfo{
			UserNum: userNum,
			Ranking: i + 1,
			Player:  leader.Leaderboards[i].Nickname,
			Elo:     utils.CalcDak(elo.TierType, elo.TierGrade, elo.Lp),
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
