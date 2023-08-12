package utils

import (
	"fmt"
)

func CalcDak(elo int, division int, lp int) string {
	newElo := elo_Formatted(elo)

	if elo == 66 || elo == 7 || elo == 8 {
		return fmt.Sprintf("%s - RP: %d", newElo, lp)
	}

	return fmt.Sprintf("%s %d - RP: %d", newElo, division, lp)
}

func elo_Formatted(eloID int) string {
	eloNames := map[int]string{
		0:  "Sem elo",
		1:  "Ferro",
		2:  "Bronze",
		3:  "Prata",
		4:  "Ouro",
		5:  "Platina",
		6:  "Diamante",
		7:  "Titan",
		8:  "Immortal",
		66: "Mythiril",
	}

	return eloNames[eloID]
}
