package structs

import "github.com/bwmarrin/discordgo"

type Commands = func(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
