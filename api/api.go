package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mariuskiessling/password-api/api/handlers"
)

// Serve starts a HTTP webserver on the specified port.
func Serve(port int) {
	log.Println("API starting on port " + strconv.Itoa(port) + ".")

	router := httprouter.New()
	router = setupRoutes(router)

	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Fatal("Could not start API on port " + strconv.Itoa(port) + ".")
	}
}

func setupRoutes(router *httprouter.Router) *httprouter.Router {
	router.GET("/ping", handlers.Ping)
	router.POST("/password", handlers.GeneratePassword)

	return router
}
