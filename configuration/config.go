package configuration

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type (
	configuration struct {
		BotToken string
	}
)

var config *configuration

const (
	filePath = "config.json"
)

func GetConfig() (*configuration, error) {
	if config == nil {
		return parseConfig()
	} else {
		return config, nil
	}
}

func parseConfig() (*configuration, error) {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	cfg := new(configuration)

	fmt.Println("Successfully Opened config.json")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &cfg)

	fmt.Println(cfg.BotToken)
	fmt.Println(byteValue)

	config = cfg

	return cfg, nil

}
