package main

import (
	"barberBot/configuration"
	"fmt"
)

func main() {
	config, err := configuration.GetConfig()

	if err != nil {
		fmt.Printf("config error load %#v", err)
	}
	fmt.Printf("config load %v", config.BotToken)
}
