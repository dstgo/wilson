package random

import (
	"encoding/binary"
	"math"
	// #nosec G404 (CWE-338): Use of weak random number generator
	"math/rand/v2"
	"os"
	"runtime"
	"time"
)

var DefaultRand = NewChaCha8()

// NewChaCha8 returns a new ChaCha8 random number generator.
// ChaCha8 is a variant of the ChaCha stream cipher, designed for
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

	s := 1.000001 + pcg.Float64() + pcg.Float64()
	v := 1.0
	imax := uint64(math.MaxUint64)

	return rand.New(rand.NewZipf(pcg, s, v, imax))
}
