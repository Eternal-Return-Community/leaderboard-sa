package commands

import (
	"erbs/src/embeds"

	"github.com/bwmarrin/discordgo"
)

func Ranking(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	if len(args) < 2 {
		s.ChannelMessageSendReply(m.ChannelID, "Adicione um nickname.", m.Reference())
		return
	}

	nickname := args[1]
	s.ChannelMessageSendEmbedReply(m.ChannelID, embeds.Ranking(nickname), m.Reference())
}
