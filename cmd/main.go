package main

import (
	"fmt"
	"intro-ai/config"
	"intro-ai/internal/server"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	db, err := sqlx.Connect("postgres", cfg.PsqlDSN)
	if err != nil {
		fmt.Printf("failed connect to database: %s", err)
		return
	}
	defer db.Close()

	s := server.NewServer(":3000", nil, 10*time.Second, 10*time.Second, 1<<20, cfg, db)
	if err := s.Run(); err != nil {
		fmt.Printf("EROROR OCCURED WHILE SHUTDOWN SERVER: %s", err)
		return
	}
}
