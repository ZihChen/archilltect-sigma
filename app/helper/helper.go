package helper

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func PrintLog(param interface{}) {
	byteData, _ := jsoniter.Marshal(param)
	show := string(byteData)
	fmt.Println("PrintLog:" + show)
	panic(show)
}
