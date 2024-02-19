package embeds

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Error(text string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:       "Ocorreu um erro",
		Description: fmt.Sprintf(text),
		Color:       0xff6961,
	}
}
