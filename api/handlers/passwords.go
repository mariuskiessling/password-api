package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mariuskiessling/password-api/password"
)

var store *password.Store = password.Init()

// GeneratePassword generates a password, encrypts and stores it.
func GeneratePassword(rw http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	//store.Add("test", "123")
	//log.Println("A new password was generated.")
	//store.Print()

	gen := password.Generator{
		Length:            8,
		Numbers:           2,
		SpecialCharacters: 2,
	}

	pw, _ := gen.Generate(0)
	fmt.Println(pw)

	rw.WriteHeader(201)
}
