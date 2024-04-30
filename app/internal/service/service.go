package service

import "pasha_bot"

type Telegram interface {
	Send(token pasha_bot.Webhook)
}

type Service struct {
	Telegram
}

func NewService() *Service {
	return &Service{
		Telegram: NewTelegramService(),
	}
}
