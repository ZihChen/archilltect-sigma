package linebothandler

import (
	"archilltect-sigma/app/service/gpt3service"
	"archilltect-sigma/app/service/linebotservice"
)

type Handler struct {
	LineBotService linebotservice.Interface
	Gpt3Service    gpt3service.Interface
}

func New() *Handler {
	return &Handler{
		LineBotService: linebotservice.New(),
		Gpt3Service:    gpt3service.New(),
	}
}
