package main

import (
	"fmt"
	"intro-ai/config"
	"intro-ai/internal/server"
	"time"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	s := server.NewServer(":3000", nil, 10*time.Second, 10*time.Second, 1<<20)
	if err := s.Run(); err != nil {
		fmt.Println("error: ", err)
		return
	}
}
