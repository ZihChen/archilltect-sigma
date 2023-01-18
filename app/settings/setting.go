package settings

import (
	"archilltect-sigma/app/structs"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config *structs.Config

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("viper read config failed")
		return err
	}

	if err = viper.Unmarshal(&Config); err != nil {
		fmt.Println("viper unmarshal err")
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file has been modified")
	})

	return err
}
