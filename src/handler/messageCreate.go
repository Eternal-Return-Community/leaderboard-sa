package handler

import (
	"erbs/src/embeds"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	channelID = "1131439928631361617"
	prefix    = "!"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	if m.ChannelID != channelID {
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("Execute o comando no canal <#%s>", channelID), m.Reference())
		return
	}

	if m.Content == prefix+"leaderboard" {
		s.ChannelMessageSendEmbedReply(m.ChannelID, embeds.Leaderboard(), m.Reference())
		return
	}

	if strings.HasPrefix(m.Content, prefix+"ranking") {
		arg := strings.TrimPrefix(m.Content, prefix+"ranking")
		nickname := strings.TrimSpace(arg)

		if len(nickname) == 0 {
			s.ChannelMessageSendReply(m.ChannelID, "Adicione um nickname.", m.Reference())
			return
		}

		s.ChannelMessageSendEmbedReply(m.ChannelID, embeds.Ranking(nickname), m.Reference())

	}

}
