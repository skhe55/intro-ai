package http

import (
	"fmt"
	"intro-ai/internal/labels"
	"intro-ai/internal/middleware"
	"net/http"
)

func MapLabelsRoutes(prefix string, h labels.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v/create", prefix), mw.Chain(h.CreateLabel(), mw.Method("POST"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v", prefix), mw.Chain(h.GetLabelsByImageId(), mw.Method("GET"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/delete/", prefix), mw.Chain(h.DeleteLabel(), mw.Method("DELETE"), mw.AuthJWTMiddleware()))
}
