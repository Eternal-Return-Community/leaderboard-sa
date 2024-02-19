package services

import (
	"encoding/json"
	"erbs/src/structs"
	"erbs/src/utils"
	"fmt"
	"net/http"
)

const url = "https://er-node.dakgg.io/api/v0/leaderboard?page=1&seasonKey=SEASON_11&serverName=saopaulo&teamMode=SQUAD&hl=en"

func Dak() ([]structs.PlayerInfo, structs.HighLP) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, structs.HighLP{}
	}

	defer resp.Body.Close()
	var leader structs.Data
	err = json.NewDecoder(resp.Body).Decode(&leader)
	if err != nil {
		fmt.Println(err)
		return nil, structs.HighLP{}
	}

	size := len(leader.Leaderboards[0:10])
	players := make([]structs.PlayerInfo, size)
	lp := highLP(leader)

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

	return players, lp
}

func highLP(leader structs.Data) structs.HighLP {
	titan := leader.Cutofss[len(leader.Cutofss)-2]
	immortal := leader.Cutofss[len(leader.Cutofss)-1]

	return structs.HighLP{
		Titan:    titan.Mmr % 250,
		Immortal: immortal.Mmr % 250,
	}

}
