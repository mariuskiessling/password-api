package main

import (
	"log"

	"github.com/mariuskiessling/password-api/api"
	"github.com/mariuskiessling/password-api/config"
)

func main() {
	config, err := config.Load("config.json")
	if err != nil {
		log.Fatal(err)
	}

	api.Serve(config.Port)
}
