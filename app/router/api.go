package router

import (
	"archilltect-sigma/app/logger"
	"github.com/gin-gonic/gin"
	"os"
)

func Setup(r *gin.Engine) {
	if os.Getenv("ENV") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(logger.GinLogger(), logger.GinRecovery(true))
}
