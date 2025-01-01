package main

import (
	"botyra/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bot, err := tgbotapi.NewBotAPI("7594228625:AAFEI3IteO3hzic59HIF0W2TN5Q8buFmlK4")
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		services.HandleUpdate(bot, update)
	}
	return nil
}
