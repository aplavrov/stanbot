package main

import (
	"log"
	"stanBot/internal/ads"
	telegramClient "stanBot/internal/client/telegram"
	"stanBot/internal/config"
	event_consumer "stanBot/internal/consumer/event-consumer"
	"stanBot/internal/event/telegram"
	"stanBot/internal/storage/files"

	"github.com/joho/godotenv"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatalf("failed to load env: %v", err)
	}
	cfg := config.New()
	s := files.New(storagePath)
	f := ads.New()

	tgClient := telegramClient.New(tgBotHost, cfg.Bot.Token)
	eventsProcessor := telegram.New(
		tgClient,
		s,
		f,
	)
	eventsProcessor.StartGetter()
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize, f)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
