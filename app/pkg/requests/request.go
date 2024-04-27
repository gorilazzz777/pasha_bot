package requests

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const token = "7183801317:AAHWE2ww6y3F7rZY5zmNu93xeLuiL-gwRR4"
const chatId = 314529904

func SendRequest(caption string, filePath string) {
	//client := &http.Client{}
	//body := &bytes.Buffer{}
	//io.Copy(body, file)
	//req, _ := http.NewRequest(http.MethodPost, respUrl, map[string]interface{}{"file": body})
	//req.URL.RawQuery = q.Encode()
	//_, err := client.Do(req)
	//if err != nil {
	//	return fmt.Errorf(err.Error())
	//}
	//return nil
	bot, err := tgbotapi.NewBotAPI(token)
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