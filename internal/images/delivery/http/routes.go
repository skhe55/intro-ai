package http

import (
	"fmt"
	"intro-ai/internal/images"
	"intro-ai/internal/middleware"
	"net/http"
)

func MapImagesRoutes(prefix string, h images.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v/", prefix), mw.Chain(h.GetAllImagesByProjectId(), mw.Method("GET"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/create", prefix), mw.Chain(h.CreateImage(), mw.Method("POST"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/delete/", prefix), mw.Chain(h.DeleteImage(), mw.Method("DELETE"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/upload/", prefix), mw.Chain(h.UploadImage(), mw.Method("POST"), mw.AuthJWTMiddleware()))
}
