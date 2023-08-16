package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	american = "https://trello.com/b/VUtjO0n6/er-roadmap"
	korean   = "https://trello.com/b/EjEt8ZPk/영원회귀-로드맵"
)

func Trello(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("Trello Americano **[Clique aqui!](%s)** \nTrello Coreano **(Sempre atualizado) [Clique aqui!](%s)**", american, korean), m.Reference())
}
