package labels

import "net/http"

type Handlers interface {
	CreateLabel() http.HandlerFunc
	GetLabelsByImageId() http.HandlerFunc
	DeleteLabel() http.HandlerFunc
}
