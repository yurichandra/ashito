package helper

import "math/rand"

func StringNumber(n int) string {
	numerals := []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = numerals[rand.Intn(len(numerals))]
	}
	return string(b)
}
