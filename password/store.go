package password

import (
	"encoding/json"
	"fmt"
)

// Store stores all generated passwords in an encrypted format.
type Store struct {
	passwords map[string][]string
}

// Init initialises a new password store.
func Init() (store *Store) {
	return &Store{
		passwords: make(map[string][]string),
	}
}

// Add stores a password in the store using the user's public key.
func (store *Store) Add(publicKey string, password string) {
	store.passwords[publicKey] = append(store.passwords[publicKey], (password))
}

// Remove deletes a user's password from the store using the user's public key.
// Please note that this does not completely removes the user but only deletes one given password.
func (store *Store) Remove(publicKey string, password string) {
	// TODO: Implement remove password function
}

// Print dumps the store's content as a JSON object into stdout.
func (store *Store) Print() {
	dump, _ := json.MarshalIndent(store.passwords, "", "  ")
	fmt.Print(string(dump))
}
