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
	rBody, err := ioutil.ReadAll(request.Body)
	if err != nil {
		writeError("Something went wrong...", 500, rw)
		return
	}

	body := &generatePasswordBody{}
	json.Unmarshal(rBody, body)

	// Check for required fields
	if body.Tag == "" || body.PublicKey == "" || body.PublicKeyFingerprint == "" {
		writeError("Missing paramters. Please check the API specification.", 400, rw)
		return
	}
	if body.Options.Numbers <= 0 || body.Options.Length >= 2048 {
		writeError("Password length can not be smaller than 0 and larger than 2048.", 400, rw)
		return
	}

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

	// Merge passwords and alternatives into one slice
	pws := append([]string{pw}, alternatives...)

	// Encrypt generated passwords and store them
	for _, pwa := range pws {
		encryptedPw, err := pk.Encrypt(pwa)
		if err != nil {
			writeError("Generated password could not be encoded.", 500, rw)
			return
		}
		err = store.Add(pk.Fingerprint, body.Tag, encryptedPw)
		if err != nil {
			writeError("Generated passwords could not be successfully stored.", 500, rw)
		}

		fmt.Printf("Stored %v => %v\n", pwa, encryptedPw)
	}

	rw.WriteHeader(201)
}
