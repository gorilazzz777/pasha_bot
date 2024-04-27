package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasha_bot"
)

func (h *Handler) TelegramWebhook(c *gin.Context) {
	var req []pasha_bot.Input
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	h.Services.Send(req[0])
	c.JSON(http.StatusOK, "OK")
}
