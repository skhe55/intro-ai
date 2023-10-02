package http

import (
	"fmt"
	"intro-ai/internal/auth"
	"intro-ai/internal/middleware"
	"net/http"
)

func MapAuthRoutes(prefix string, h auth.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v/register", prefix), h.Register())
	mux.HandleFunc(fmt.Sprintf("/%v/login", prefix), (h.Login()))
}
