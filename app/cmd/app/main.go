package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
	bot "pasha_bot"
	"pasha_bot/internal/handler"
	"pasha_bot/internal/service"
)

// @title PashBot API
// @version 1.0
// @description     Сервис шлет Павлику фотки Наиля в ТГ
func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env file: %s", err.Error())
	}

	services := service.NewService()
	handlers := handler.NewHandler(services)

	server := new(bot.Server)
	if err := server.Run(os.Getenv("APP_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error on start app: %s", err.Error())
	}
}
