package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Ping returns a *pong* indicating a healthy service.
func Ping(rw http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	rw.WriteHeader(200)
	writeResponse("pong", rw)
}
