package embeds

import (
	"erbs/src/services"
	"fmt"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Ranking(nickname string) *discordgo.MessageEmbed {

	info := services.Erbs(nickname)

	if info.Nickname == "" {
		return Error(fmt.Sprintf("**%s** não existe no banco de dados do **Eternal Return**", nickname))
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Ranking (#%s) - %s", rank(info.Rank), info.Nickname),
		Description: fmt.Sprintf("Elo: %s \nMmr: %d", info.Elo, info.Mmr),
		Color:       0xA7C7E7,
	}
}

func rank(top int) string {
	if top == 0 {
		return "N/D"
	}
	return strconv.Itoa(top)
}
