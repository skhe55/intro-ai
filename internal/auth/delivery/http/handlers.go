package http

import (
	"encoding/json"
	"intro-ai/internal/server/response"
	"intro-ai/internal/server/writer"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	writer.WriterJsonHeader(w)

	if r.Method != http.MethodGet {
		res, err := json.Marshal(response.Error("UNKNOWN HTTP METHOD."))
		if err != nil {
			log.Fatalf("ERROR (Hello): %s", err)
		}
		w.Write(res)
		return
	}

	res, err := json.Marshal(response.OK())
	if err != nil {
		log.Fatalf("ERROR (HELLO): %s", err)
	}

	w.Write(res)
}
