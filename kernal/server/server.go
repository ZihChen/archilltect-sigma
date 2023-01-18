package server

import (
	"archilltect-sigma/app/router"
	"archilltect-sigma/app/settings"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var r *gin.Engine

func Run() {
	r = gin.New()
	// 註冊路由
	router.Setup(r)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Config.Port),
		Handler: r,
	}

	zap.L().Debug("[Server listen]:", zap.String("port", strconv.Itoa(settings.Config.Port)))

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("[Listen failed]:", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("[Shutdown server]: prepare shutdown server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("[Server error]: ", zap.Error(err))
	}
	zap.L().Info("[Server exiting]: server already shutdown")
}
