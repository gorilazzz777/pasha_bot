package main

import (
	"github.com/sirupsen/logrus"
	bot "pasha_bot"
	"pasha_bot/internal/handler"
	"pasha_bot/internal/service"
)

// @title PashBot API
// @version 1.0
// @description     Сервис шлет Павлику фотки Наиля в ТГ
func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := new(bot.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error on start app: %s", err.Error())
	}
}
