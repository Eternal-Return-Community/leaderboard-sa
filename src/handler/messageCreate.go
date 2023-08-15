package handler

import (
	"erbs/src/commands"
	"erbs/src/structs"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	channelID = "1129836224681615380"
	prefix    = "!"
)

var (
	Commands = map[string]structs.Commands{
		"leaderboard": commands.Leaderboard,
		"ranking":     commands.Ranking,
	}
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Split(m.Content[len(prefix):], " ")
	command := strings.ToLower(args[0])

	if m.ChannelID != channelID && len(args[0]) >= 1 {
		s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("Execute o comando no canal <#%s>", channelID), m.Reference())
		return
	}

	for commandList, commandRun := range Commands {
		if command == commandList {
			commandRun(s, m, args)
			return
		}
	}

}
