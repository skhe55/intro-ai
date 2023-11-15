package annotations

import "net/http"

type Handlers interface {
	GetAnnotationsByLabelIdOrImageId() http.HandlerFunc
	CreateAnnotation() http.HandlerFunc
	DeleteAnnotation() http.HandlerFunc
}
