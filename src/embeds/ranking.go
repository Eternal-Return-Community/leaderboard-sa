package embeds

import (
	"erbs/src/services"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ranking(nickname string) *discordgo.MessageEmbed {

	player := services.ShowPlayers(nickname)
	if len(player) < 3 {
		return Error(fmt.Sprintf("**%s** não existe no banco de dados do **Eternal Return**.", nickname))
	}

	info := player[2]
	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Ranking (#%d) - %s", info.ServerRanking.Rank, nickname),
		Description: fmt.Sprintf("Elo: %s - RP: %d \nMmr: %d", info.PlayerTier.Name, info.PlayerTier.LP, info.PlayerTier.MMR),
		Color:       0xA7C7E7,
	}
}
