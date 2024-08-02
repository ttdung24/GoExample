package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

}

// RandomInt generates a random integer
func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max - min + 1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte((c))
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(7)
}

// RandomMoney generates a ramdom amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomMoney generates a ramdom currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[r.Intn((n))]
}
