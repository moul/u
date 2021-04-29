package u

import "math/rand"

const randomLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomLetters returns a string containing 'n' random letters.
func RandomLetters(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = randomLetters[rand.Intn(len(randomLetters))]
	}
	return string(b)
}
