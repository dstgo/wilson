package randx

import (
	"crypto/rand"
	"math/big"
)

// SecInt64 returns a secure random 64-bit integer in the range [0, max).
func SecInt64(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
