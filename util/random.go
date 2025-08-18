package util

import (
	"math/rand"
	"strings"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Seed the random number generator to ensure different results each time
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max (inclusive).
func RandomInt(min, max int64) int64 {
	if min >= max {
		panic("max must be greater than min")
	}
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n using lowercase letters.
func RandomString(n int) string {
	var sb strings.Builder
	k := len(letters)
	for i := 0; i < n; i++ {
		c := letters[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomOwner generates a random owner name of length 6.
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money between 0 and 1000.
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code from a predefined list.
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "JPY", "GBP", "AUD", "KRW", "CAD", "CHF", "CNY", "SEK"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email address.
func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}
