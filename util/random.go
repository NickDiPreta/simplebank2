package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// RandomInt generates a random int64 between min and max
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

// func RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < len(alphabet); i++ {
		c := alphabet[r.Intn(len(alphabet))]
		sb.WriteByte(c)
	}

	return sb.String()
}

// func RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// func RandomMoney generates a random amount of money between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// func RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[r.Intn(n)]
}
