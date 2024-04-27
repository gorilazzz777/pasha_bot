package service

import (
	"fmt"
	"pasha_bot"
	"pasha_bot/pkg/requests"
)

type TelegramService struct{}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

func (s *TelegramService) Send(webhook pasha_bot.Input) {
	webhook.Format()
	mess := fmt.Sprintf("дата: %s\nПерсонаж: %s",
		webhook.Date, webhook.Persona.Name)
	requests.SendRequest(mess, webhook.Image.Thumbnail)
	mess = fmt.Sprintf("Оригинал: %s", webhook.Persona.Name)
	requests.SendRequest(mess, webhook.Image.FullFrame)
}
