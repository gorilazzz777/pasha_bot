package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"pasha_bot/pkg/tgbot"
	"strings"
)

func main() {
	//if err := godotenv.Load(); err != nil {
	//	logrus.Fatalf("error loading env file: %s", err.Error())
	//}

	tgBot, err := tgbotapi.NewBotAPI(os.Getenv("TG_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	inlineButtonText := ""
	tgBot.Debug = true

	log.Printf("Authorized on account %s", tgBot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.GetUpdatesChan(u)

	// Loop through each update.
	for update := range updates {
		// Check if we've gotten a message update.
		// Check if we've gotten a message update.
		if update.CallbackQuery != nil {
			// Respond to the callback query, telling Telegram to show the user
			// a message with the data received.

			// And finally, send a message containing the data received.
			text := strings.Split(update.CallbackQuery.Data, ";")[0]
			data := strings.Split(update.CallbackQuery.Data, ";")[1]
			filePath := strings.Split(update.CallbackQuery.Data, ";")[2]
			fmt.Println("------------------------------------------")
			fmt.Println(update.CallbackQuery.Data)
			fmt.Println("------------------------------------------")
			switch text {
			case "Show origin":
				inlineButtonText = "Show thumb"
			case "Show thumb":
				inlineButtonText = "Show origin"

			}
			tgbot.SendRequest("", data, inlineButtonText, filePath)
		}
	}
}
