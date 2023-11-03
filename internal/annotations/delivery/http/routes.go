package http

import (
	"fmt"
	"intro-ai/internal/annotations"
	"intro-ai/internal/middleware"
	"net/http"
)

func MapAnnotationsRoutes(prefix string, h annotations.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v/", prefix), mw.Chain(h.GetAnnotationsByLabelId(), mw.Method("GET"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/create", prefix), mw.Chain(h.CreateAnnotation(), mw.Method("POST"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/delete/", prefix), mw.Chain(h.DeleteAnnotation(), mw.Method("DELETE"), mw.AuthJWTMiddleware()))
}
