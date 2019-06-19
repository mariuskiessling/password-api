package password

import (
	"crypto/rand"
	"log"
	"math/big"
)

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
