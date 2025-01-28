package randx

import (
	"crypto/rand"
	"math/big"
)

func Int64(max int64) int64 {
	n, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		panic(err)
	}
	return n.Int64()
}
