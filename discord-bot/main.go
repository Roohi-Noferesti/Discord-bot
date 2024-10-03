package main

import (
	"log"

	"test.domain/discordbot/bot"
	"test.domain/discordbot/config"
)

func main() {
	cfg := config.LoadConfig()
	if cfg.Token == "" {
		log.Fatal("Bot token not provided!")
	}

	bot.Start(cfg)
}