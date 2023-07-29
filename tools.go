package toolkit

import (
	"math/rand"
	"time"
)

const randomStringSource = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_+"

// Tools is the type used to instanciate the toolkit module. Any variable of this type will have access
// to all the methods with the receiver of *Tools
type Tools struct{}

// RandomStringL returns a string of random characters with the length n, using randomStringSource
// as the source of characters
func (t *Tools) RandomStringL(n int) string {

	s, r := make([]rune, n), []rune(randomStringSource)
	for i := range s {
		s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
		s[i] = r[s1.Intn(len(r))]
	}
	return string(s)
}
