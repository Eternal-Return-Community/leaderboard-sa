package services

import (
	"encoding/json"
	"erbs/src/structs"
	"erbs/src/utils"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	Key = utils.Env().Key
)

const (
	baseURL    = "https://open-api.bser.io/v1"
	getUserNum = "/user/nickname?query="
)

func Client(endpoint string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		fmt.Println(err)
	}

	//Header
	req.Header.Add("x-api-key", Key)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(body)
	}

	return string(body)

}

func Erbs(nickname string) structs.RankedInfo {

	//Pega o ID da conta através do nickname.
	var userNum structs.ER_PlayerInfo
	err := json.NewDecoder(strings.NewReader(Client(getUserNum + nickname))).Decode(&userNum)
	if err != nil {
		fmt.Println(err)
	}

	//Pega as informações da ranked do usuário.
	var ranked structs.Ranked
	err = json.NewDecoder(strings.NewReader(Client(fmt.Sprintf("/rank/%d/19/3", userNum.User.UserNum)))).Decode(&ranked)
	if err != nil {
		fmt.Println(err)
	}

	return structs.RankedInfo{
		Mmr:      ranked.UserRank.Mmr,
		Elo:      utils.CalcElo(ranked.UserRank.Mmr, ranked.UserRank.Rank),
		Rank:     ranked.UserRank.Rank,
		Nickname: ranked.UserRank.Nickname,
	}

}
