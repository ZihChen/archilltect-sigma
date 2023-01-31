package settings

import (
	"archilltect-sigma/app/structs"
	"embed"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

var Config *structs.Config

func Init(f embed.FS) error {
	env := os.Getenv("ENV")

	pathList := []string{
		"env/" + env + "/config.yaml",
	}

	for k := range pathList {
		configFile, err := f.ReadFile(pathList[k])
		if err != nil {
			fmt.Println(fmt.Sprintf("Read File Error: %v", err.Error()))
			return err
		}

		if err = yaml.Unmarshal(configFile, &Config); err != nil {
			fmt.Println(fmt.Sprintf("Yaml Unmarshal Error: %v", err.Error()))
			return err
		}
	}

	return nil
}
