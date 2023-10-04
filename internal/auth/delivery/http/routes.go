package http

import (
	"fmt"
	"intro-ai/internal/auth"
	"intro-ai/internal/middleware"
	"net/http"
)

func MapAuthRoutes(prefix string, h auth.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v/register", prefix), mw.Chain(h.Register(), mw.Method("POST")))
	mux.HandleFunc(fmt.Sprintf("/%v/login", prefix), mw.Chain(h.Login(), mw.Method("POST")))
}
