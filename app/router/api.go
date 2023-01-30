package router

import (
	"archilltect-sigma/app/handler/linebothandler"
	"archilltect-sigma/app/logger"
	"github.com/gin-gonic/gin"
	"os"
)

func Setup(r *gin.Engine) {
	if os.Getenv("ENV") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	lineBotHandler := linebothandler.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.POST("/callback", lineBotHandler.Callback)
}
