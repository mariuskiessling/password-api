package password

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
)

// Store stores all generated passwords in an encrypted format.
type Store struct {
	passwords map[string][]string
}

// A Generator is used to store all options needed to generate a password.
type Generator struct {
	Length            int
	SpecialCharacters int
	Numbers           int
}

var (
	mapping = []byte{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
		'A',
		'B',
		'C',
		'D',
		'E',
		'F',
		'G',
		'H',
		'I',
		'J',
		'K',
		'L',
		'M',
		'N',
		'O',
		'P',
		'Q',
		'R',
		'S',
		'T',
		'U',
		'V',
		'W',
		'X',
		'Y',
		'Z',
	}
)

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

// Generate generates a password for the given Generator options and *n* alternative passwords.
func (gen *Generator) Generate(numAlternatives int) (password string, alternatives []string) {
	// We will not introduce a recursive function to generate a random string of
	// the needed length due to the 'newly' introduced string builder. This
	// dramatically reduces the number of memory operations needed.

	pw := make([]byte, gen.Length)

	// Step 1: Populate the password array only with random letters (both upper- and lower case)
	for i := 0; i < gen.Length; i++ {
		bigR, err := rand.Int(rand.Reader, big.NewInt(52))
		r := bigR.Int64()
		if err != nil {
			log.Fatal("Something went completely wrong while generating random numbers...")
		}

		pw[i] = mapping[r]
	}

	return string(pw), nil
}
