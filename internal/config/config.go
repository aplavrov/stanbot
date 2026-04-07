package config

import (
	"log"
	"os"
)

type Config struct {
	Bot BotConfig
}

type BotConfig struct {
	Token string
}

func New() *Config {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	botConfig := BotConfig{Token: token}
	return &Config{Bot: botConfig}
}
