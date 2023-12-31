package images

import "net/http"

type Handlers interface {
	GetAllImagesByProjectId() http.HandlerFunc
	GetImageById() http.HandlerFunc
	CreateImage() http.HandlerFunc
	DeleteImage() http.HandlerFunc
	UploadImage() http.HandlerFunc
}
