package util

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func MakeContentButtons(results []string) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	for i, title := range results {
		button := tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("%d. %s", i+1, title), strconv.Itoa(i+1))
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(button))
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func MakeLanguageButtons() tgbotapi.InlineKeyboardMarkup {
	row := tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("🇷🇺 Русский", "rus"),
		tgbotapi.NewInlineKeyboardButtonData("🇬🇧 English", "eng"),
	)
	return tgbotapi.NewInlineKeyboardMarkup(row)
}
