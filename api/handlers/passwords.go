package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mariuskiessling/password-api/password"
)

var store *password.Store = password.Init()

type generatePasswordBody struct {
	Alternatives int    `json:"alternatives"`
	PublicKey    string `json:"public_key"`
	Options      struct {
		Length            int `json:"length"`
		SpecialCharacters int `json:"special_characters"`
		Numbers           int `json:"numbers"`
	} `json:"options"`
}

// GeneratePassword generates a password, encrypts and stores it.
func GeneratePassword(rw http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// TODO: Add error handling
	rBody, _ := ioutil.ReadAll(request.Body)

	body := &generatePasswordBody{}
	json.Unmarshal(rBody, body)

	gen := password.Generator{
		Length:            body.Options.Length,
		Numbers:           body.Options.Numbers,
		SpecialCharacters: body.Options.SpecialCharacters,
	}

	pw, alternatives := gen.Generate(body.Alternatives)
	fmt.Println(pw)
	fmt.Println(alternatives)

	rw.WriteHeader(201)
}
