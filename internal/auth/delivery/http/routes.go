package http

import (
	"fmt"
	"intro-ai/internal/auth"
	"intro-ai/internal/middleware"
	"net/http"
)

func MapAuthRoutes(prefix string, h auth.Handlers, mw *middleware.MiddlewareManager) {
	http.HandleFunc(fmt.Sprintf("/%v/register", prefix), h.Register())
	http.HandleFunc(fmt.Sprintf("/%v/login", prefix), mw.AuthJWTMiddleware()(h.Login()))
}
