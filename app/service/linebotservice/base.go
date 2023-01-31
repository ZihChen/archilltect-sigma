package linebotservice

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"go.uber.org/zap"
	"net/http"
	"os"
	"sync"
)

type Interface interface {
	// GetClient 取得 Line-bot 實例
	GetClient() *linebot.Client
}

type service struct{}

var bot *linebot.Client
var singleton *service
var once sync.Once

func New() Interface {
	once.Do(func() {
		singleton = &service{}
	})
	return singleton
}

func (s *service) GetClient() *linebot.Client {
	var err error
	bot, err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
		linebot.WithHTTPClient(&http.Client{}))
	if err != nil {
		zap.L().Error("[LINE-bot get client error]:", zap.Any("error", err))
	}
	return bot
}
