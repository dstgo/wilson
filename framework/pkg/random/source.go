package random

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"math"
	// #nosec G404 (CWE-338): Use of weak random number generator
	"math/rand/v2"
	"os"
	"runtime"
	"time"
)

var (
	DefaultRand = ChaCha8Rand

	// CryptoRand low speed and high secure cryptographic randomness source
	CryptoRand = NewCrypto()
	// ChaCha8Rand middle speed and middle secure cryptographic randomness source
	ChaCha8Rand = NewChaCha8()
	// PCGRand high speed and low secure cryptographic randomness source
	PCGRand = NewPCG()
	// ZipfRand is zipf distribution random number generator
	ZipfRand = NewZipf()
)

// NewChaCha8 returns a new ChaCha8Rand random number generator.
// ChaCha8Rand is a variant of the ChaCha stream cipher, designed for
// high-speed encryption and random number generation. It offers:
// - Strong cryptographic properties
// - Excellent performance on a wide range of platforms
// - Resistance to timing attacks
// - Simple and efficient implementation
// This implementation uses PCG to generate the initial seed.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewChaCha8() *rand.Rand {
	pcg := NewPCG()

	var seed [32]byte
	for i := 0; i < 32; i += 8 {
		binary.BigEndian.PutUint64(seed[i:], pcg.Uint64()%math.MaxUint8)
	}

	return rand.New(rand.NewChaCha8(seed))
}

// NewPCG returns a new PCG random number generator.
// PCG (Permuted Congruential Generator) is a family of simple,
// fast, and statistically excellent random number generators.
// It offers better statistical quality, smaller state size,
// longer periods, and greater performance than many alternatives.
// PCG combines a linear congruential generator with output permutation,
// providing high-quality randomness with minimal computational overhead.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewPCG() *rand.Rand {
	ng := runtime.NumGoroutine()
	pid := os.Getpid()
	micro := time.Now().UnixMicro()
	nano := time.Now().UnixNano()

	seed1 := uint64(micro ^ int64(ng)&int64(pid))
	seed2 := uint64(nano&int64(ng) ^ int64(pid))

	return rand.New(rand.NewPCG(seed1, seed2))
}

// NewZipf returns a new Zipf-distributed random number generator.
// It uses PCG as the underlying random source and generates a dynamic
// 's' parameter in the range [1.000001, 3.000001). The 'v' parameter
// is fixed at 1.0, and the maximum value is set to math.MaxUint64.
// This generator is suitable for modeling power-law distributions in
// various applications, such as word frequencies or city populations.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewZipf() *rand.Rand {
	pcg := NewPCG()

	s := 1.0 + 1e-6 + pcg.Float64() + pcg.Float64()
	v := 1.0
	imax := uint64(math.MaxUint64)

	return rand.New(rand.NewZipf(pcg, s, v, imax))
}

// NewCrypto returns a concurrent-safe new cryptographically secure random number generator.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewCrypto() *rand.Rand {
	return rand.New(&_Crypto{})
}

// NewCryptoBackoff returns a new concurrent-safe cryptographically secure random number generator
// with a specified backoff generator.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewCryptoBackoff(backoff *rand.Rand) *rand.Rand {
	return rand.New(&_Crypto{backoff: backoff})
}

// _Crypto is a concurrent-safe cryptographically secure random number generator that with a backoff generator.
type _Crypto struct {
	// #nosec G404 (CWE-338): Use of weak random number generator
	backoff *rand.Rand
}

func (c *_Crypto) Uint64() uint64 {
	if c.backoff == nil {
		// default to ChaCha8 if no backoff generator is provided.
		c.backoff = NewChaCha8()
	}

	var buf [binary.MaxVarintLen64]byte

	// if crypto/rand.Read fails, fall back to the backoff generator.
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		return c.backoff.Uint64()
	}

	return binary.BigEndian.Uint64(buf[:])
}
