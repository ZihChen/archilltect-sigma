package main

import (
	"archilltect-sigma/app/settings"
	"fmt"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed:%s \n", err)
		return
	}
}
