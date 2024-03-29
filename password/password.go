package password

import (
	"bytes"
	"crypto/rand"
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
	letters = []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o',
		'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D',
		'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S',
		'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}

	specialCharacters = []byte{
		'!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/',
		':', ';', '<', '=', '>', '?', '@', '[', ']',
	}

	mappingVowelToNumber = map[byte]byte{
		'A': '4',
		'E': '3',
		'I': '1',
		'O': '0',
		'U': '8',
	}
)

// Generate generates a password for the given Generator options and *n* alternative passwords.
func (gen *Generator) Generate(numAlternatives int) (password string, alternatives []string) {
	password = gen.GeneratePassword()

	for i := 0; numAlternatives != 0 && i < numAlternatives; i++ {
		alternatives = append(alternatives, gen.GeneratePassword())
	}

	return
}

// GeneratePassword generates *one* password for the given generator.
func (gen *Generator) GeneratePassword() (password string) {
	// Step 1: Populate the password array only with random letters (both upper- and lower case)
	pw := generateLetterBytes(gen.Length)

	// Step 2: Randomly replace vowels with their mapped numbers
	pw, replacedVowels := generateVowelReplacements(pw, gen.Numbers)

	// Step 3: Place gen.numbers random numbers in random locations of the password
	pw = generateNumberReplacements(pw, gen.Numbers-replacedVowels)

	// Step 4: Place gen.specialCharacters random chars in random locations of the password
	pw = generateSpecialCharsReplacements(pw, gen.SpecialCharacters)

	return string(pw)
}

func generateSpecialCharsReplacements(b []byte, chars int) (pw []byte) {
	for i := 0; i < chars; i++ {
		index := generateNumber(int64(len(b)))

		// Generate a new index while the generated index contains a number or
		// already contains a special character
		// -> TODO: An easier solution would be to check if the given byte is a
		// letter
		for byteIsNumber(b[index]) || byteIsSpecialCharacter(b[index]) {
			index = generateNumber(int64(len(b)))
		}

		b[index] = specialCharacters[generateNumber(int64(len(specialCharacters)))]
	}

	return b
}

func generateNumberReplacements(b []byte, maxNumbers int) (pw []byte) {
	for i := 0; i < maxNumbers; i++ {
		index := generateNumber(int64(len(b)))

		// Generate a new index while the generated index contains a number or is
		// already replaced with a special character
		for byteIsNumber(b[index]) || bytes.Contains(specialCharacters, []byte{b[index]}) {
			index = generateNumber(int64(len(b)))
		}

		r := []byte(strconv.Itoa(generateNumber(10)))
		b[index] = r[0]
	}

	return b
}

func generateVowelReplacements(b []byte, maxNumbers int) (pw []byte, replacedVowels int) {
	replacedVowels = 0

	for i := 0; i < len(b); i++ {
		if _, ok := mappingVowelToNumber[b[i]]; ok && replacedVowels < maxNumbers {
			// Generate random binary number that will determine wether the vowel
			// will be replaced with a number
			replaceVowel := generateNumber(2)
			if replaceVowel == 1 {
				oldPw := append(make([]byte, 0, len(b)), b...)

				b[i] = mappingVowelToNumber[b[i]]

				log.Printf("Vowel replacement during generation: %v -> %v\n", string(oldPw), string(b))

				replacedVowels++
			}
		}
	}

	return b, replacedVowels
}

func generateLetterBytes(length int) (b []byte) {
	b = make([]byte, length)

	for i := 0; i < length; i++ {
		b[i] = letters[generateNumber(int64(len(letters)))]
	}

	return
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

func byteIsNumber(b byte) bool {
	return b >= 48 && b <= 57
}

func byteIsSpecialCharacter(b byte) bool {
	for _, specialCharacter := range specialCharacters {
		if b == specialCharacter {
			return true
		}
	}

	return false
}
