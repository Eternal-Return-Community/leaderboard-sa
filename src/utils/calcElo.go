package utils

import (
	"fmt"
	"math"
	"strconv"
)

func CalcElo(mmr int) string {
	elo := eloFormatted(mmr)
	division := 4 - int(math.Floor(float64((mmr%1000)/250)))
	rp := mmr % 250

	if elo == "Mytiril" || elo == "Titan" || elo == "Immortal" {
		return fmt.Sprintf("%s - RP: %d", elo, mmr%6000)
	}

	return elo + " " + strconv.Itoa(division) + " - RP: " + strconv.Itoa(rp)
}

func eloFormatted(mmr int) string {

	elo := ""
	if mmr > 0 && mmr < 1000 {
		elo = "Ferro"
	} else if mmr >= 1000 && mmr < 2000 {
		elo = "Bronze"
	} else if mmr >= 2000 && mmr < 3000 {
		elo = "Prata"
	} else if mmr >= 3000 && mmr < 4000 {
		elo = "Ouro"
	} else if mmr >= 4000 && mmr < 5000 {
		elo = "Platina"
	} else if mmr >= 5000 && mmr < 6000 {
		elo = "Diamante"
	} else if mmr >= 6000 && mmr < 7000 {
		elo = "Mytiril"
	} else if mmr >= 7000 && mmr < 8000 {
		elo = "Titan"
	} else if mmr >= 8000 {
		elo = "Immortal"
	} else {
		elo = "Sem Elo"
	}

	return elo
}
