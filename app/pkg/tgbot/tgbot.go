package tgbot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"strconv"
)

func SendRequest(caption, filePath, inlineButtonText, inlineButtonData string) {
	data := fmt.Sprintf("%s;%s;%s", inlineButtonText, inlineButtonData, filePath)
	var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(inlineButtonText, data),
		),
	)
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
	photo.ReplyMarkup = numericKeyboard
	_, err = bot.Send(photo)
	if err != nil {
		fmt.Println(err)
		fmt.Println(data)
	}
}
