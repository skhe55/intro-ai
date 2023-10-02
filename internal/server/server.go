package server

import (
	"context"
	"intro-ai/config"
	"intro-ai/pkg/logger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

type Server struct {
	Addr           string
	Handler        http.Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
	cfg            *config.Config
	db             *sqlx.DB
	logger         logger.Logger
}

func NewServer(
	Addr string,
	Handler http.Handler,
	ReadTimeout time.Duration,
	WriteTimeout time.Duration,
	MaxHeaderBytes int,
	cfg *config.Config,
	db *sqlx.DB,
	logger logger.Logger,
) *Server {
	return &Server{
		Addr:           Addr,
		Handler:        Handler,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
		cfg:            cfg,
		db:             db,
		logger:         logger,
	}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.Addr,
		Handler:        s.Handler,
		ReadTimeout:    s.ReadTimeout,
		WriteTimeout:   s.WriteTimeout,
		MaxHeaderBytes: s.MaxHeaderBytes,
	}

	mux := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{s.cfg.OriginRemote},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodPut,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	s.MapHandlers(mux)

	server.Handler = cors.Handler(mux)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("ERROR OCCURED WHILE STARTING SERVER: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 1*time.Second)
	defer shutdown()

	return server.Shutdown(ctx)
}
