package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

func SendRequest(caption string, filePath string) {
	chatId, err := strconv.ParseInt(os.Getenv("TG_CHAT_ID"), 10, 64)
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	photo := tgbotapi.NewPhoto(chatId, tgbotapi.FilePath(filePath))
	photo.Caption = caption
	_, err = bot.Send(photo)
	if err != nil {
		fmt.Println(err)
	}
}
