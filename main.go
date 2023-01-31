package main

import (
	"archilltect-sigma/app/logger"
	"archilltect-sigma/app/settings"
	_ "archilltect-sigma/docs"
	"archilltect-sigma/kernal/server"
	"embed"
	"fmt"
	"go.uber.org/zap"
)

//go:embed env/*
var f embed.FS

func main() {
	if err := settings.Init(f); err != nil {
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
