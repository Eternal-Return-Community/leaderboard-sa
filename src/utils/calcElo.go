package utils

import (
	"fmt"
	"math"
)

func CalcElo(mmr int, rank int) string {
	elo := eloFormatted(mmr, rank)
	division := 4 - int(math.Floor(float64((mmr%1000)/250)))
	rp := mmr % 250

	if elo == "Mythril" || elo == "Titan" || elo == "Immortal" {
		return fmt.Sprintf("%s - RP: %d", elo, mmr%6000)
	}

	return fmt.Sprintf("%s %d - RP: %d", elo, division, rp)
}

func eloFormatted(mmr int, rank int) string {

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
	} else if mmr >= 6000 && rank > 700 {
		elo = "Mythril"
	} else if rank >= 201 && rank <= 700 {
		elo = "Titan"
	} else if rank <= 200 {
		elo = "Immortal"
	} else {
		elo = "Sem Elo"
	}

	return elo
}
