package embeds

import (
	"erbs/src/structs"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Patchnotes(patchnotes structs.Notes) *discordgo.MessageEmbed {

	notes := patchnotes

	return &discordgo.MessageEmbed{
		Title:       notes.I18ns.En_Us.Title,
		Description: fmt.Sprintf("%s \n\n**>> [Clique aqui para ir até o Site](%s) <<**", notes.I18ns.En_Us.Description, notes.I18ns.En_Us.Link),
		Image: &discordgo.MessageEmbedImage{
			URL: notes.Thumbnail_url,
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Ultima atualização: " + translateTimeDescription(notes.I18ns.En_Us.Time),
		},
		Color: 0xA7C7E7,
	}
}

func translateTimeDescription(time string) string {
	text := strings.Split(time, " ")
	formattedTime := text[1] + " " + text[2]

	var result string
	switch formattedTime {
	case "hours ago":
		result = text[0] + " hora(s) atrás"
		break
	case "week ago":
		result = text[0] + " semana(s) atrás"
		break
	case "month ago":
		result = text[0] + " mês atrás"
		break
	}

	return result

}
