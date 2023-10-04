package httpError

import (
	"encoding/json"
	"errors"
	"intro-ai/internal/server/response"
	"net/http"
)

var (
	INTERNAL     = "INTERNAL ERROR"
	NON_INTERNAL = "ERROR"
	WRONG_DTO    = "CHECK DTO FIELDS"
	WRONG_ID     = "CHECK PASSED ID"
)

var (
	InvalidJWTToken     = errors.New("invalid jwt token")
	InvalidJWTClaims    = errors.New("invalid jwt claims")
	InvalidJWTSignature = errors.New("invalid token signature")
)

type HttpError struct{}

func NewHttpError() HttpError {
	return HttpError{}
}

func (e *HttpError) InternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	body, _ := json.Marshal(response.Response{Message: INTERNAL, Status: INTERNAL, Result: nil})
	w.Write(body)
}

func (e *HttpError) NonInternalError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	body, _ := json.Marshal(response.Response{Message: message, Status: NON_INTERNAL, Result: nil})
	w.Write(body)
}
