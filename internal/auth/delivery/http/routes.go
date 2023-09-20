package http

import (
	"fmt"
	"net/http"
)

func MapAuthRoutes(prefix string) {
	http.HandleFunc(fmt.Sprintf("/%v/hello", prefix), Hello)
}
