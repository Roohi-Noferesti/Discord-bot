package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"test.domain/discordbot/commands"
	"test.domain/discordbot/config"
	"test.domain/discordbot/utils"
)

var BotID string

func Start(cfg *config.Config) {
	dg, err := discordgo.New("Bot " + cfg.Token)
	utils.CheckNilErr(err)

	u, err := dg.User("@me")
	utils.CheckNilErr(err)

	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()
	utils.CheckNilErr(err)

	defer dg.Close()

	log.Println("Bot is now running. Press CTRL+C to exit.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	commands.ExecuteCommand(s, m)
}