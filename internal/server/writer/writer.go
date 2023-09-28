package writer

import "net/http"

func WriterJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
