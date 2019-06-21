package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mariuskiessling/password-api/password"
)

var store *password.Store = password.Init()

type generatePasswordBody struct {
	Tag                  string `json:"tag"`
	Alternatives         int    `json:"alternatives"`
	PublicKey            string `json:"public_key"`
	PublicKeyFingerprint string `json:"public_key_fingerprint"`
	Options              struct {
		Length            int `json:"length"`
		SpecialCharacters int `json:"special_characters"`
		Numbers           int `json:"numbers"`
	} `json:"options"`
}

// GeneratePassword generates a password, encrypts and stores it.
func GeneratePassword(rw http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// TODO: Add error handling
	rBody, _ := ioutil.ReadAll(request.Body)
	// TODO: Add check for required arguments

	body := &generatePasswordBody{}
	json.Unmarshal(rBody, body)

	gen := password.Generator{
		Length:            body.Options.Length,
		Numbers:           body.Options.Numbers,
		SpecialCharacters: body.Options.SpecialCharacters,
	}

	pw, alternatives := gen.Generate(body.Alternatives)

	pk, err := password.LoadPublicKey(body.PublicKey)
	if err != nil {
		writeError(err.Error(), 400, rw)
		return
	}
	if pk.Fingerprint != body.PublicKeyFingerprint {
		writeError("Provided fingerprint does not match calculated one for the given public key.", 400, rw)
		return
	}

	// TODO: Add encryption of passwords before storing
	err = store.Add(pk.Fingerprint, body.Tag, pw)
	for _, pwa := range alternatives {
		err = store.Add(pk.Fingerprint, body.Tag, pwa)
	}
	if err != nil {
		writeError("Generated passwords could not be successfully stored.", 500, rw)
	}

	store.Print()

	rw.WriteHeader(201)
}
