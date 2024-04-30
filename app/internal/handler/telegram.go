package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"pasha_bot"
)

func (h *Handler) TelegramWebhook(c *gin.Context) {
	var req []pasha_bot.Webhook
	forLog, _ := httputil.DumpRequest(c.Request, true)
	fmt.Println(string(forLog))
	if err := c.ShouldBindJSON(&req); err != nil {
		return
	}
	h.Services.Send(req[0])
	c.JSON(http.StatusOK, "OK")
}
