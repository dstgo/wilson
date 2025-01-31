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
	_ "unsafe"
)

var (
	// DefaultRng is the default concurrent-safe random number generator.
	DefaultRng = RuntimeRng
	CryptoRng  = NewCrypto()
	RuntimeRng = NewRuntime()
)

// NewCrypto returns a new concurrent-safe cryptographically secure random number generator.
// It utilizes the operating system's cryptographically secure random number generator:
// - Windows: Uses CryptGenRandom, which leverages the CryptoAPI
// - Linux: Uses getrandom() system call, falling back to /dev/urandom if unavailable
// - macOS: Uses SecRandomCopyBytes, which is backed by the Yarrow algorithm
// This generator offers:
// - High-quality, unpredictable random numbers suitable for cryptographic use
// - Protection against common cryptographic attacks
// - Automatic seeding from system entropy sources
// - Consistent interface across different operating systems
// Note: While highly secure, this generator may be slower than non-cryptographic alternatives.
func NewCrypto() *Rand {
	return New(&_Crypto{})
}

// NewCryptoBackoff returns a new concurrent-safe cryptographically secure random number generator
// with a specified backoff generator.
func NewCryptoBackoff(backoff *Rand) *Rand {
	return New(&_Crypto{backoff: backoff})
}

// _Crypto is a concurrent-safe cryptographically secure random number generator
// that with a backoff generator.
type _Crypto struct {
	backoff *Rand
}

func (c *_Crypto) Uint64() uint64 {
	if c.backoff == nil {
		// default to ChaCha8 if no backoff generator is provided.
		c.backoff = RuntimeRng
	}

	var buf [binary.MaxVarintLen64]byte

	// if crypto/rand.Read fails, fall back to the backoff generator.
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		return c.backoff.Uint64()
	}

	return binary.BigEndian.Uint64(buf[:])
}

// NewRuntime returns a new concurrent-safe, high-speed random number generator.
// This generator offers:
// - Concurrent safety through goroutine-local random number generators
// - High performance suitable for non-cryptographic use cases
// - Based on the ChaCha8 algorithm for quality randomness
// - Automatic seeding from runtime information
// Concurrency safety is achieved by utilizing Go's runtime, where each goroutine
// has its own local random number generator. This approach eliminates the need
// for locks or atomic operations, resulting in excellent performance in concurrent scenarios.
// The underlying algorithm is ChaCha8, which provides a good balance of speed and randomness quality.
// Note: While fast and suitable for many applications, this is not cryptographically secure.
func NewRuntime() *Rand {
	return New(_Runtime{})
}

//go:linkname runtime_rand runtime.rand
func runtime_rand() uint64

// _Runtime is a concurrent-safe local random number generator.
type _Runtime struct{}

func (_Runtime) Uint64() uint64 {
	// underlying algorithm is chacha8.
	return runtime_rand()
}

// NewChaCha8 returns a new non-concurrent-safe ChaCha8 random number generator.
// ChaCha8Rng is a variant of the ChaCha stream cipher, designed for
// high-speed encryption and random number generation. It offers:
// - Strong cryptographic properties
// - Excellent performance on a wide range of platforms
// - Resistance to timing attacks
// - Simple and efficient implementation
// This implementation uses PCG to generate the initial seed.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewChaCha8() *Rand {
	seed := BytesN(New(NewPCG()), 32, 0, math.MaxUint8)
	return New(rand.NewChaCha8(*(*[32]byte)(seed)))
}

// NewPCG returns a new non-concurrent-safe PCG random number generator.
// PCG (Permuted Congruential Generator) is a family of simple,
// fast, and statistically excellent random number generators. It offers:
// - High statistical quality output
// - Small state size and fast period advancement
// - Excellent performance on various platforms
// - Multiple, independent streams
// - Simple and efficient implementation
// This implementation uses runtime and system information to generate the initial seeds.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewPCG() *Rand {
	ng := runtime.NumGoroutine()
	pid := os.Getpid()
	micro := time.Now().UnixMicro()
	nano := time.Now().UnixNano()

	seed1 := uint64(micro ^ int64(ng)&int64(pid))
	seed2 := uint64(nano&int64(ng) ^ int64(pid))

	return New(rand.NewPCG(seed1, seed2))
}

// NewZipf returns a new non-concurrent-safe Zipf-distributed random number generator.
// Zipf distribution is a discrete probability distribution commonly used for modeling:
// - Word frequencies in natural languages
// - Population sizes of cities
// - Income distribution in economics
// It offers:
// - Power-law behavior for modeling real-world phenomena
// - Adjustable parameters for different distribution shapes
// - Efficient generation of skewed random numbers
// - Suitable for various applications in data science and simulations
// This implementation uses PCG as the underlying random source with dynamic parameters.
// #nosec G404 (CWE-338): Use of weak random number generator
func NewZipf() *Rand {
	pcg := New(NewPCG())

	s := pcg.Float64Range(1.0+1e-6, 3.0)
	v := 1.0
	imax := uint64(math.MaxUint64)

	return New(rand.NewZipf(pcg.Rand, s, v, imax))
}
