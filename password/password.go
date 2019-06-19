package password

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"strconv"
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
		pw[i] = mapping[generateNumber(52)]
	}

	// Step 2: Place gen.numbers random numbers in random locations of the character password
	for i := 0; i < gen.Numbers; i++ {
		index := generateNumber(int64(gen.Length))

		// Generate a new index while the generated index contains a number
		for pw[index] >= 48 && pw[index] <= 57 {
			index = generateNumber(int64(gen.Length))
		}

		// TODO: Rewrite this ugly conversion from int -> string -> []byte -> byte
		r := []byte(strconv.Itoa(generateNumber(10)))
		fmt.Println(index, r)
		pw[index] = r[0]
	}

	return string(pw), nil
}

func generateNumber(max int64) int {
	// Note that this creates the interval [0, max) thus including 0 and
	// excluding max.
	bigR, err := rand.Int(rand.Reader, big.NewInt(max))

	if err != nil {
		log.Fatal("Something went completely wrong while generating random numbers...")
	}

	return int(bigR.Int64())
}
