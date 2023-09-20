package writer

import "net/http"

func WriterJsonHeader(w http.ResponseWriter) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
}
