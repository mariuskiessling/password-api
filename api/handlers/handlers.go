package handlers

import (
	"fmt"
	"net/http"
)

func writeResponse(content string, rw http.ResponseWriter) {
	fmt.Fprintln(rw, content)
}
