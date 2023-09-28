package main

import (
	"fmt"
	"intro-ai/config"
	"intro-ai/internal/server"
	"intro-ai/pkg/logger"
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

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	db, err := sqlx.Connect("postgres", cfg.PsqlDSN)
	if err != nil {
		logger.Error("FAILED CONNECT TO DATABASE", err)
		return
	}
	defer db.Close()

	s := server.NewServer(":3000", nil, 10*time.Second, 10*time.Second, 1<<20, cfg, db, logger)
	if err := s.Run(); err != nil {
		logger.Error("FAILED START SERVER", err)
		return
	}
}
