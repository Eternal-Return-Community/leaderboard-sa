package services

import (
	"encoding/json"
	"erbs/src/structs"
	"erbs/src/utils"
	"fmt"
	"net/http"
)

var (
	seasonID = "10"
	url      = fmt.Sprintf("https://er.dakgg.io/v1/leaderboard?page=1&seasonKey=SEASON_%s&serverName=Sao+Paulo&teamMode=SQUAD&hl=en", seasonID)
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

func ShowPlayers(nickname string) []structs.TeamModeSummary {
	refresh(nickname)
	resp, err := http.Get(fmt.Sprintf("https://dak.gg/er/_next/data/MgiH6rT0nezahHCKG_Q4y/players/%s.json?teamMode=ALL&season=SEASON_%s&name=%s&hl=en", nickname, seasonID, nickname))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	var info structs.ServerRanking
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		fmt.Println(err)
	}

	return info.PageProps.Data.TeamModeSummary

}

func refresh(nickname string) {

	resp, err := http.Get(fmt.Sprintf("https://er.dakgg.io/v1/players/%s/renew", nickname))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	//ignore this code
	/*
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
	*/

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
