package main

import (
	"botyra/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	// Получаем значение переменной BOT_API_KEY
	botAPIKey := os.Getenv("BOT_API_KEY")
	if botAPIKey == "" {
		log.Fatal("BOT_API_KEY не задан")
	}
	bot, err := tgbotapi.NewBotAPI(botAPIKey)
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
