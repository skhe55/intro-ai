package labels

import "net/http"

type Handlers interface {
	CreateLabel() http.HandlerFunc
	GetLabelsByProjectId() http.HandlerFunc
	DeleteLabel() http.HandlerFunc
}
