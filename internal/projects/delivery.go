package projects

import "net/http"

type Handlers interface {
	GetAllProjects() http.HandlerFunc
	CreateProject() http.HandlerFunc
	DeleteProject() http.HandlerFunc
}
