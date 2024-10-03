package config

import (
	"os"

	"github.com/joho/godotenv"
	"test.domain/discordbot/utils"
)

type Config struct {
	Token string
	Prefix string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	utils.CheckNilErr(err)

	return &Config{
		Token: os.Getenv("DISCORD_TOKEN"),
		Prefix: os.Getenv("COMMAND_PREFIX"),
	}
}