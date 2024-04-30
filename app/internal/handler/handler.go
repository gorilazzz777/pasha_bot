package handler

import (
	"github.com/gin-gonic/gin"
	"pasha_bot/internal/service"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	telegram := router.Group("/")
	{
		telegram.POST("/webhook", h.TelegramWebhook)
	}
	return router
}
