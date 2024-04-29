package service

import (
	"fmt"
	"pasha_bot"
	"pasha_bot/pkg/img"
	"pasha_bot/pkg/tgbot"
)

type TelegramService struct{}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

func (s *TelegramService) Send(webhook pasha_bot.Input) {
	webhook.Format()
	img.MergeImages(webhook.Image.Thumbnail, webhook.Image.Original, "/img/merge.jpg")
	mess := fmt.Sprintf("дата: %s\nПерсонаж: %s",
		webhook.Date, webhook.Persona.Name)
	tgbot.SendRequest(mess, webhook.Image.FullFrame, "/img/merge.jpg")
	//mess = fmt.Sprintf("Полниразмерное фото: %s", webhook.Persona.Name)
	//tgbot.SendRequest(mess, webhook.Image.FullFrame)
	//img.DeleteImage("./merge.jpg")
}
