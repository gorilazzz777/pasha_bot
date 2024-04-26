package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasha_bot"
)

func (h *Handler) TelegramWebhook(c *gin.Context) {
	var req pasha_bot.Input
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Fields validation failed. Make sure that the required fields are specified")
		return
	}
	trackNumber, err := h.Services.Send(req)
	if err != nil {
		c.JSON(http.StatusOK, map[string]interface{}{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{"trackNumber": trackNumber})
	}
}
