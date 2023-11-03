package annotations

import "net/http"

type Handlers interface {
	GetAnnotationsByLabelId() http.HandlerFunc
	CreateAnnotation() http.HandlerFunc
	DeleteAnnotation() http.HandlerFunc
}
