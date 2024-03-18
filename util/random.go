package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n using characters from the alphabet.
//
// n: the length of the random string to generate.
// string: the randomly generated string.
func RandomString(n int) string {
	var builder strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		builder.WriteByte(c)
	}
	return builder.String()
}
