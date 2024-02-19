package commands

import (
	"erbs/src/embeds"
	"erbs/src/services"

	"github.com/bwmarrin/discordgo"
)

func Patchnotes(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	patchnotes := services.Pathnotes()
	s.ChannelMessageSendEmbedReply(m.ChannelID, embeds.Patchnotes(patchnotes), m.Reference())

}
