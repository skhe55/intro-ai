package server

import (
	"net/http"
	"time"
)

type Server struct {
	Addr           string
	Handler        http.Handler
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func NewServer(Addr string, Handler http.Handler, ReadTimeout time.Duration, WriteTimeout time.Duration, MaxHeaderBytes int) *Server {
	return &Server{
		Addr:           Addr,
		Handler:        Handler,
		ReadTimeout:    ReadTimeout,
		WriteTimeout:   WriteTimeout,
		MaxHeaderBytes: MaxHeaderBytes,
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

	s.MapHandlers()

	return server.ListenAndServe()
}
