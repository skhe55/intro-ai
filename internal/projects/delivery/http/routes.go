package http

import (
	"fmt"
	"intro-ai/internal/middleware"
	"intro-ai/internal/projects"
	"net/http"
)

func MapProjectsRoutes(prefix string, h projects.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v", prefix), mw.Chain(h.GetAllProjects(), mw.Method("GET"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/create", prefix), mw.Chain(h.CreateProject(), mw.Method("POST"), mw.AuthJWTMiddleware()))
	mux.HandleFunc(fmt.Sprintf("/%v/delete/", prefix), mw.Chain(h.DeleteProject(), mw.Method("DELETE"), mw.AuthJWTMiddleware()))
}
