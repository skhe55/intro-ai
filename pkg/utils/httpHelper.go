package utils

import (
	"errors"
	"intro-ai/internal/server/response"
	"net/http"
)

func CheckHttpMethod(w http.ResponseWriter, r *http.Request, validMethod string) error {
	if r.Method != validMethod {
		res, _ := ToJSON[response.Response](response.Error("NOT ALLOWED HTTP METHOD."))
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write(res)
		return errors.New("unknown method")
	}
	return nil
}
