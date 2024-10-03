package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"test.domain/discordbot/config"
)

var commandMap = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) {
	"hello": helloCommand,
}

func ExecuteCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	cfg := config.LoadConfig()

	if !strings.HasPrefix(m.Content, cfg.Prefix) {
		return
	}

	args := strings.Fields(m.Content[len(cfg.Prefix):])
	if len(args) == 0 {
		return
	}

	command := args[0]
	if cmdFunc, exists := commandMap[command]; exists {
		cmdFunc(s, m)
	}

}

func helloCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := fmt.Sprintf("Hello, %s!", m.Author.Username)
	s.ChannelMessageSend(m.ChannelID, message)
}