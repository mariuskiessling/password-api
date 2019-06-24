package store

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Store stores all generated passwords in an encrypted format.
//
// Passwords are stored under the user's public key fingerprint and tag. Any
// tag can have one or password associated with it. Alternative passwords are
// stored next to the original one inside the tags' array.
type Store struct {
	passwords map[string]map[string][]string
}

// Init initialises a new password store.
func Init() (store *Store) {
	return &Store{
		passwords: make(map[string]map[string][]string),
	}
}

// Add stores a password in the store using the user's public key.
func (store *Store) Add(fingerprint string, tag string, password string) (err error) {
	// Initialize fingerprint map if fingerprint is unknown
	if _, ok := store.passwords[fingerprint]; !ok {
		store.passwords[fingerprint] = make(map[string][]string)
	}

	store.passwords[fingerprint][tag] = append(store.passwords[fingerprint][tag], (password))
	return nil
}

// Remove deletes a user's password from the store using the user's public key.
// Please note that this does not completely removes the user but only deletes one given password.
func (store *Store) Remove(fingerprint string, tag string, password string) {
	// TODO: Implement remove password function
}

// Retrieve retries a user's passwords from the store. If a tag is supplied
// only passwords stored under the tag are returned.
func (store *Store) Retrieve(fingerprint string, tag string) (passwords map[string][]string, err error) {
	pws, ok := store.passwords[fingerprint]
	if !ok {
		return nil, errors.New("No passwords could be found for the supplied public key fingerprint.")
	}

	if tag != "" {
		tagPasswords, ok := pws[tag]
		if !ok {
			return nil, errors.New("No passwords could be found for this tag.")
		}

		tagMap := make(map[string][]string)
		tagMap[tag] = tagPasswords
		return tagMap, nil
	}

	return pws, nil
}

// Print dumps the store's content as a JSON object into stdout.
func (store *Store) Print() {
	dump, _ := json.MarshalIndent(store.passwords, "", "  ")
	fmt.Print(string(dump))
}
