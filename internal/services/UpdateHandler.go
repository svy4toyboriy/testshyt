package services

import (
	"botyra/util"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const (
	resultsAmount = 5
	audioFormat   = "m4a"
	videoFormat   = "mp4"
)

var (
	phrases = [][]string{
		{"Отлично! Выберите формат /audio или /video. Изменить его можно будет в любой момент.",
			"Выберите аудио stupid:", "Скачиваю...",
			"Загружаю...", "А вот и оно!", "Формат изменён на аудио. Введите свой запрос!",
			"Формат изменён на видео. Введите свой запрос!", "Недоступно. Выберите другую кнопку.", "Выберите видео:"},
		{"Alright! Choose format /audio or /video. You'll be able to change it anytime.", "Choose audio:", "Downloading...",
			"Uploading...", "Here it is!", "Format changed to audio. Search for anything!",
			"Format changed to video. Search for anything!", "Unavailable. Try another one.", "Choose video:"},
	}
	language = 0
	format   = audioFormat
	addition = " audio"
	chatID   int64
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil {
		chatID = update.Message.Chat.ID
		message := update.Message.Text

		switch message {
		case "/start":
			keyboard := util.MakeLanguageButtons()
			msg := tgbotapi.NewMessage(chatID, "Добро пожаловать! Выберите язык.\nWelcome! Choose language.")
			msg.ReplyMarkup = keyboard
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending start message: %v", err)
			}
		case "/audio":
			format = audioFormat
			addition = " audio"
			util.SendMessage(bot, chatID, phrases[language][5])
		case "/video":
			format = videoFormat
			addition = ""
			util.SendMessage(bot, chatID, phrases[language][6])
		default:
			err := Search(message+addition, resultsAmount)
			if err != nil {
				err = fmt.Errorf("failed to make a search: %w", err)
				fmt.Println(err)
				return
			}

			keyboard := util.MakeContentButtons(ContentTitle)
			msg := tgbotapi.NewMessage(chatID, phrases[language][1])
			msg.ReplyMarkup = keyboard
			if _, err := bot.Send(msg); err != nil {
				log.Printf("Error sending content buttons: %v", err)
			}
		}
	} else if update.CallbackQuery != nil {
		data := update.CallbackQuery.Data
		if data == "rus" {
			language = 0
			util.SendMessage(bot, chatID, phrases[language][0])
		} else if data == "eng" {
			language = 1
			util.SendMessage(bot, chatID, phrases[language][0])
		} else {
			util.SendMessage(bot, chatID, phrases[language][2])
			buttonNumber, err := strconv.Atoi(data)
			if err != nil {
				log.Printf("Error parsing button number: %v", err)
				return
			}
			buttonNumber--

			re := regexp.MustCompile("[^\\da-zA-Zа-яёА-ЯЁ]")
			fileName := re.ReplaceAllString(ContentTitle[buttonNumber], "")

			songUrl := ContentURL[buttonNumber]
			fmt.Println(fileName)

			audioPath := filepath.Join("D:/JavaContent/IdeaProjects/MusicDealerWin/resources/Audio/downloads",
				fmt.Sprintf("%s.%s", fileName, format))

			if _, err := os.Stat(audioPath); os.IsNotExist(err) {
				log.Printf("File does not exist, downloading: %s", audioPath)
				if err := Call(fileName, songUrl, format); err != nil {
					log.Printf("Error downloading file: %v", err)
					util.SendMessage(bot, chatID, phrases[language][7])
					return
				}
			}

			if _, err := os.Stat(audioPath); os.IsNotExist(err) {
				log.Printf("File still does not exist after download: %s", audioPath)
				util.SendMessage(bot, chatID, phrases[language][7])
				return
			}

			util.SendMessage(bot, chatID, phrases[language][3])

			if format == "m4a" {
				audio := tgbotapi.NewAudio(chatID, tgbotapi.FilePath(audioPath))
				audio.Title = ContentTitle[buttonNumber]
				audio.Caption = phrases[language][4]
				if _, err := bot.Send(audio); err != nil {
					log.Printf("Error sending audio: %v", err)
				}
			} else if format == "mp4" {
				video := tgbotapi.NewVideo(chatID, tgbotapi.FilePath(audioPath))
				video.Caption = phrases[language][4]
				if _, err := bot.Send(video); err != nil {
					log.Printf("Error sending video: %v", err)
				}
			}
		}
	}
}
