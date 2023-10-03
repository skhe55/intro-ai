package http

import (
	"fmt"
	"intro-ai/internal/middleware"
	"intro-ai/internal/projects"
	"net/http"
)

func MapProjectsRoutes(prefix string, h projects.Handlers, mw *middleware.MiddlewareManager, mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/%v/", prefix), mw.AuthJWTMiddleware()(h.GetAllProjects()))
	mux.HandleFunc(fmt.Sprintf("/%v/create", prefix), mw.AuthJWTMiddleware()(h.CreateProject()))
	mux.HandleFunc(fmt.Sprintf("/%v/delete/", prefix), mw.AuthJWTMiddleware()(h.DeleteProject()))
}
