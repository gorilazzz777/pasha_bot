package service

import (
	"fmt"
	"os"
	"pasha_bot"
	"pasha_bot/pkg/img"
	"pasha_bot/pkg/tgbot"
)

type TelegramService struct{}

func NewTelegramService() *TelegramService {
	return &TelegramService{}
}

func (s *TelegramService) Send(webhook pasha_bot.Webhook) {
	webhook.Format()
	imgPath := fmt.Sprintf("%s/%s", os.Getenv("IMG_PATH"), os.Getenv("MERGE_IMG_NAME"))
	img.MergeImages(webhook.Image.Thumbnail, webhook.Image.Original, imgPath)
	mess := fmt.Sprintf("дата: %s\nПерсонаж: %s",
		webhook.Date, webhook.Persona.Name)
	tgbot.SendRequest(mess, imgPath, "Show origin", webhook.Image.FullFrame)
	//mess = fmt.Sprintf("Полноразмерное фото: %s", webhook.Persona.Name)
	//tgbot.SendRequest(mess, webhook.Image.FullFrame)
}
