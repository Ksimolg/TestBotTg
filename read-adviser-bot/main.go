package main

import (
	"flag"
	"log"
	"read-adviser-bot/consumer/event-consumer"

	tgClient "read-adviser-bot/clients/telegram"
	"read-adviser-bot/events/telegram"
	"read-adviser-bot/storage/files"
)

//./read-adviser-bot -tg-bot-token '5514456455:AAHryKrnHOnaqA4f-W563b8mBA23W0vXtHM'
const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)
	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stop", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access too telegram bot",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
