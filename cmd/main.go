package main

import (
	"fmt"
	"intro-ai/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(config.PsqlDSN)
}
