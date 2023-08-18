package handler

import (
	"erbs/src/commands"
	"erbs/src/structs"
	"erbs/src/utils"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	channelID = "1129836224681615380"
	prefix    = utils.Env().Prefix
)

var (
	Commands = map[string]structs.Commands{
		"leaderboard": commands.Leaderboard,
		"ranking":     commands.Ranking,
		"comandos":    commands.Comandos,
		"patchnotes":  commands.Patchnotes,
		"trello":      commands.Trello,
		"dodge":       commands.Dodge,
	}
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID || !strings.HasPrefix(m.Content, prefix) {
		return
	}

	args := strings.Split(m.Content[len(prefix):], " ")
	command := strings.ToLower(args[0])

	for commandList, commandRun := range Commands {
		if command == commandList {
			if m.ChannelID != channelID {
				s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("Execute o comando no canal <#%s>", channelID), m.Reference())
				return
			}
			commandRun(s, m, args)
			return
		}
	}

}
