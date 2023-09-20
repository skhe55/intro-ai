package auth

import "net/http"

type Handlers interface {
	Hello() http.HandlerFunc
}
