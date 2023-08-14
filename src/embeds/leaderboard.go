package embeds

import (
	"erbs/src/services"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Leaderboard() *discordgo.MessageEmbed {

	info, lp := services.Dak()

	//Retornar a Embed de Error quando não possui nenhum jogador.
	if len(info) == 0 {
		return Error("Atualmente não possui nenhum jogador ranqueado, nessa nova temporada. \nVerifique a tabela mais tarde.")
	}

	var fields []*discordgo.MessageEmbedField
	for _, player := range info {
		field := &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%d. %s", player.Ranking, player.Player),
			Value:  player.Elo,
			Inline: false,
		}
		fields = append(fields, field)
	}

	return &discordgo.MessageEmbed{
		Title:       "Leaderboard (SQUAD) - Super Server SA",
		Description: fmt.Sprintf("* **LP** necessário para subir pro **Titan**: `%d` **|** **Immortal**: `%d`", lp.Titan, lp.Immortal),
		Fields:      fields,
		Color:       0xA7C7E7,
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Você não está no TOP 10? use: !ranking <nickname>",
		},
	}
}
