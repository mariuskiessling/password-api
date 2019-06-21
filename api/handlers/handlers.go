package handlers

import (
	"encoding/json"
	"net/http"
)

type errorMessage struct {
	Error string `json:"error"`
}

func writeResponse(content string, rw http.ResponseWriter) {
	rw.Write([]byte(content))
}

func writeJSON(content interface{}, code int, rw http.ResponseWriter) {
	json, _ := json.Marshal(content)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(json)
}

func writeError(message string, code int, rw http.ResponseWriter) {
	writeJSON(&errorMessage{Error: message}, code, rw)
}
