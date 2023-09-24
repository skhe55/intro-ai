package http

import (
	"fmt"
	"intro-ai/internal/auth"
	"net/http"
)

func MapAuthRoutes(prefix string, h auth.Handlers) {
	http.HandleFunc(fmt.Sprintf("/%v/register", prefix), h.Register())
	http.HandleFunc(fmt.Sprintf("/%v/login", prefix), h.Login())
}
