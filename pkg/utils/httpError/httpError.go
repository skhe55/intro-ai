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
)

var (
	InvalidJWTToken  = errors.New("INVALID JWT TOKEN")
	InvalidJWTClaims = errors.New("INVALID JWT CLAIMS")
)

type HttpError struct {
	Writer http.ResponseWriter
}

func NewHttpError(w http.ResponseWriter) *HttpError {
	return &HttpError{Writer: w}
}

func (e *HttpError) InternalError() {
	e.Writer.WriteHeader(http.StatusInternalServerError)
	body, _ := json.Marshal(response.Response{Message: INTERNAL, Status: INTERNAL, Result: nil})
	e.Writer.Write(body)
}

func (e *HttpError) NonInternalError(status int, message string) {
	e.Writer.WriteHeader(status)
	body, _ := json.Marshal(response.Response{Message: message, Status: NON_INTERNAL, Result: nil})
	e.Writer.Write(body)
}
