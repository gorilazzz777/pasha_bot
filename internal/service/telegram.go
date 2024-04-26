package service

import (
	"fmt"
	"pasha_bot"
	"pasha_bot/pkg/requests"
)

const baseUrl = "https://api.telegram.org"
const sendMessageMethod = "sendPhoto"
const token = "7183801317:AAHWE2ww6y3F7rZY5zmNu93xeLuiL-gwRR4"
const chatId = "314529904"

type TelegramService struct{}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

func (s *TelegramService) Send(webhook pasha_bot.Input) (int, error) {
	url := fmt.Sprintf("%s/bot%s/%s?chat_id=%s&caption=%s&photo=%s",
		baseUrl, token, sendMessageMethod, chatId, webhook.Message, webhook.Img)
	_, err := requests.SendRequest(url)
	if err != nil {
		return 0, err
	}
	return 3, nil
	//return 12, nil
}
