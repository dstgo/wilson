package random

import (
	"math"
)

// BytesN returns a random slice of n bytes in the range [low, high].
// #nosec G404 (CWE-338): Use of weak random number generator
func BytesN(rng *Rand, n int, low, high byte) []byte {
	if low > high {
		panic("low > high")
	}

	bs := make([]byte, n)
	// overflow protects, such as math.MaxUint8 - 0 + 1
	delta := uint32(high) - uint32(low) + 1
	for i := 0; i < n; i++ {
		bs[i] = byte(rng.Uint32N(delta)) + low
	}
	return bs
}

// Bytes returns a random slice of n bytes in the range [0, 255].
// #nosec G404 (CWE-338): Use of weak random number generator
func Bytes(n int) []byte {
	return BytesN(DefaultRand, n, 0, math.MaxUint8)
}

// BytesSeq returns a random slice of n bytes as a string.
// #nosec G404 (CWE-338): Use of weak random number generator
func BytesSeq(n int) string {
	return string(Bytes(n))
}

// RunesN returns a random slice of n runes in the range [low, high].
// #nosec G404 (CWE-338): Use of weak random number generator
func RunesN(rng *Rand, n int, low, high rune) []rune {
	if low > high {
		panic("low > high")
	}

	rs := make([]rune, n)
	// overflow protects, such as math.MaxInt32 - 0 + 1
	delta := uint32(high) - uint32(low) + 1
	for i := 0; i < n; i++ {
		rs[i] = rune(rng.Uint32N(delta) + uint32(low))
	}
	return rs
}

// Runes returns a random slice of n runes in the range [0, math.MaxInt32].
// #nosec G404 (CWE-338): Use of weak random number generator
func Runes(n int) []rune {
	return RunesN(DefaultRand, n, 0, math.MaxInt32)
}

// RunesSeq returns a random slice of n runes as a string.
// #nosec G404 (CWE-338): Use of weak random number generator
func RunesSeq(n int) string {
	return string(Runes(n))
}

// DigitsN returns a random string of n digits.
// #nosec G404 (CWE-338): Use of weak random number generator
func DigitsN(rng *Rand, n int) string {
	return string(RunesN(rng, n, '0', '9'))
}

// Digits returns a random string of n digits using the default random source.
// #nosec G404 (CWE-338): Use of weak random number generator
func Digits(n int) string {
	return DigitsN(DefaultRand, n)
}

// LowerN returns a random string of n lowercase letters.
// #nosec G404 (CWE-338): Use of weak random number generator
func LowerN(rng *Rand, n int) string {
	return string(RunesN(rng, n, 'a', 'z'))
}

// Lower returns a random string of n lowercase letters using the default random source.
// #nosec G404 (CWE-338): Use of weak random number generator
func Lower(n int) string {
	return LowerN(DefaultRand, n)
}

// UpperN returns a random string of n uppercase letters.
// #nosec G404 (CWE-338): Use of weak random number generator
func UpperN(rng *Rand, n int) string {
	return string(RunesN(rng, n, 'A', 'Z'))
}

// Upper returns a random string of n uppercase letters using the default random source.
// #nosec G404 (CWE-338): Use of weak random number generator
func Upper(n int) string {
	return UpperN(DefaultRand, n)
}

// LettersN returns a random string of n letters (mixed case).
// #nosec G404 (CWE-338): Use of weak random number generator
func LettersN(rng *Rand, n int) string {
	letters := make([]rune, n)
	for i := 0; i < n; i++ {
		if rng.Uint64()&1 == 1 {
			letters[i] = rng.Int32N(26) + 'A'
		} else {
			letters[i] = rng.Int32N(26) + 'a'
		}
	}
	return string(letters)
}

// Letters returns a random string of n letters (mixed case) using the default random source.
// #nosec G404 (CWE-338): Use of weak random number generator
func Letters(n int) string {
	return LettersN(DefaultRand, n)
}

// AlphaNumericN returns a random string of n alphanumeric characters.
// #nosec G404 (CWE-338): Use of weak random number generator
func AlphaNumericN(rng *Rand, n int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := range result {
		result[i] = charset[rng.IntN(len(charset))]
	}
	return string(result)
}

// AlphaNumeric returns a random string of n alphanumeric characters using the default random source.
// #nosec G404 (CWE-338): Use of weak random number generator
func AlphaNumeric(n int) string {
	return AlphaNumericN(DefaultRand, n)
}
