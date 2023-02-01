package linebothandler

import (
	"archilltect-sigma/app/structs"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
	"strings"
)

var bot *linebot.Client

func (h *Handler) Callback(c *gin.Context) {
	bot = h.LineBotService.GetClient()
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		zap.L().Error("[LINE-bot parse request error]:", zap.Any("error", err))
	}
	for _, event := range events {
		userID := event.Source.UserID
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				h.pushMessageToUser(userID, linebot.NewTextMessage("🤖：執行中，請稍候..."))
				var res structs.CompletionResponse
				res, err = h.Gpt3Service.Completion(c, structs.CompletionRequest{
					Prompt:    message.Text,
					MaxTokens: 1000,
					N:         1,
				})
				if err != nil {
					zap.L().Error("[Gpt3 completion error]:", zap.Any("error", err.Error()))
					return
				}

				h.pushMessageToUser(userID, linebot.NewTextMessage(strings.TrimLeft(res.Choices[0].Text, "\n\n")))
				zap.L().Debug("[Model respond]:",
					zap.String("prompt", message.Text),
					zap.String("resp", res.Choices[0].Text))
			}
		}
	}
}

func (h *Handler) pushMessageToUser(id string, messages ...linebot.SendingMessage) {
	_, err := bot.PushMessage(id, messages...).Do()
	if err != nil {
		zap.L().Error("[LINE-bot push message error]:", zap.Any("error", err.Error()))
	}
}
