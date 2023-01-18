package main

import (
	"archilltect-sigma/app/logger"
	"archilltect-sigma/app/settings"
	"archilltect-sigma/kernal/server"
	"fmt"
	"go.uber.org/zap"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("[Init settings failed]:%s \n", err)
		return
	}

	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed:%s \n", err)
		return
	}
	zap.L().Debug("[Logger init]: success!")

	server.Run()
}
