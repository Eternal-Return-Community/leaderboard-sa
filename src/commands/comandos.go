package commands

import (
	"erbs/src/utils"
	"fmt"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix = utils.Env().Prefix
	path   = "./src/commands"
)

func Comandos(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	commandList, err := os.ReadDir(path)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "Ocorreu um erro ao carregar a lista de comandos.", m.Reference())
	}

	var commands []string
	for _, command := range commandList {
		if strings.Replace(command.Name(), ".go", "", 1) != "comandos" {
			commands = append(commands, strings.Replace(prefix+command.Name(), ".go", "", 1))
		}
	}

	if len(commands) == 0 {
		s.ChannelMessageSendReply(m.ChannelID, "Nenhum comando disponível", m.Reference())
		return
	}

	s.ChannelMessageSendReply(m.ChannelID, fmt.Sprintf("Lista de Comandos (%d) \n%s\n", len(commands), strings.Join(commands, ", ")), m.Reference())

}
