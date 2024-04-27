package service

import "pasha_bot"

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Telegram interface {
	Send(token pasha_bot.Input)
}

type Service struct {
	Telegram
}

func NewService() *Service {
	return &Service{
		Telegram: NewTelegramService(),
	}
}
