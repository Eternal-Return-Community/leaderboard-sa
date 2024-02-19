package commands

import (
	"erbs/src/embeds"

	"github.com/bwmarrin/discordgo"
)

func Leaderboard(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSendEmbedReply(m.ChannelID, embeds.Leaderboard(), m.Reference())
}
