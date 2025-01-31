package random

import (
	"math"
)

// Byte returns a random byte.
func (r *Rand) Byte() byte {
	return byte(r.Uint32Range(0, math.MaxUint8+1))
}

// Byte returns a random byte using the default random source.
func Byte() byte {
	return runtimeRng.Byte()
}

// SecByte returns a cryptographically secure random byte.
func SecByte() byte {
	return cryptoRng.Byte()
}

// BytesN generates a random slice of n bytes in the range [low, high].
func (r *Rand) BytesN(n int, low, high byte) []byte {
	if low > high {
		panic("low > high")
	}

	bs := make([]byte, n)
	// overflow protection, such as math.MaxUint8 - 0 + 1
	delta := uint32(high) - uint32(low) + 1
	for i := 0; i < n; i++ {
		bs[i] = byte(r.Uint32N(delta)) + low
	}
	return bs
}

// BytesN generates a random slice of n bytes in the range [low, high] using the default random source.
func BytesN(n int, low, high byte) []byte {
	return runtimeRng.BytesN(n, low, high)
}

// SecBytesN generates a cryptographically secure random slice of n bytes in the range [low, high].
func SecBytesN(n int, low, high byte) []byte {
	return cryptoRng.BytesN(n, low, high)
}

// BytesSeqN generates a random slice of n bytes as a string.
func BytesSeqN(n int, low, high byte) string {
	return string(BytesN(n, low, high))
}

// SecBytesSeqN generates a cryptographically secure random slice of n bytes as a string.
func SecBytesSeqN(n int, low, high byte) string {
	return string(SecBytesN(n, low, high))
}

// Rune returns a random rune.
func (r *Rand) Rune() rune {
	return rune(r.Uint32Range(0, math.MaxInt32+1))
}

// Rune returns a random rune using the default random source.
func Rune() rune {
	return runtimeRng.Rune()
}

// SecRune returns a cryptographically secure random rune.
func SecRune() rune {
	return cryptoRng.Rune()
}

// RunesN generates a random slice of n runes in the range [low, high].
func (r *Rand) RunesN(n int, low, high rune) []rune {
	if low > high {
		panic("low > high")
	}

	rs := make([]rune, n)
	// overflow protection, such as math.MaxInt32 - 0 + 1
	delta := uint32(high) - uint32(low) + 1
	for i := 0; i < n; i++ {
		rs[i] = rune(r.Uint32N(delta) + uint32(low))
	}
	return rs
}

// RunesN generates a random slice of n runes in the range [low, high] using the default random source.
func RunesN(n int, low, high rune) []rune {
	return runtimeRng.RunesN(n, low, high)
}

// SecRunesN generates a cryptographically secure random slice of n runes in the range [low, high].
func SecRunesN(n int, low, high rune) []rune {
	return cryptoRng.RunesN(n, low, high)
}

// RunesSeq generates a random slice of n runes as a string.
func RunesSeq(n int, low, high rune) string {
	return string(RunesN(n, low, high))
}

// SecRunesSeq generates a cryptographically secure random slice of n runes as a string.
func SecRunesSeq(n int, low, high rune) string {
	return string(SecRunesN(n, low, high))
}

// DigitsN generates a random string of n digits.
func (r *Rand) DigitsN(n int) string {
	return string(r.BytesN(n, '0', '9'))
}

// DigitsN generates a random string of n digits using the default random source.
func DigitsN(n int) string {
	return runtimeRng.DigitsN(n)
}

// SecDigitsN generates a cryptographically secure random string of n digits.
func SecDigitsN(n int) string {
	return cryptoRng.DigitsN(n)
}

// LowerN generates a random string of n lowercase letters.
func (r *Rand) LowerN(n int) string {
	return string(r.RunesN(n, 'a', 'z'))
}

// LowerN generates a random string of n lowercase letters using the default random source.
func LowerN(n int) string {
	return runtimeRng.LowerN(n)
}

// SecLowerN generates a cryptographically secure random string of n lowercase letters.
func SecLowerN(n int) string {
	return cryptoRng.LowerN(n)
}

// UpperN generates a random string of n uppercase letters.
func (r *Rand) UpperN(n int) string {
	return string(r.RunesN(n, 'A', 'Z'))
}

// Upper generates a random string of n uppercase letters using the default random source.
func Upper(n int) string {
	return runtimeRng.UpperN(n)
}

// SecUpperN generates a cryptographically secure random string of n uppercase letters.
func SecUpperN(n int) string {
	return cryptoRng.UpperN(n)
}

// LettersN generates a random string of n letters (mixed case).
func (r *Rand) LettersN(n int) string {
	letters := make([]rune, n)
	for i := 0; i < n; i++ {
		if r.Uint64()&1 == 1 {
			letters[i] = r.Int32N(26) + 'A'
		} else {
			letters[i] = r.Int32N(26) + 'a'
		}
	}
	return string(letters)
}

// Letters generates a random string of n letters (mixed case) using the default random source.
func Letters(n int) string {
	return runtimeRng.LettersN(n)
}

// SecLetters generates a cryptographically secure random string of n letters (mixed case).
func SecLetters(n int) string {
	return cryptoRng.LettersN(n)
}

const alphaNumericCharSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// AlphaNumericN generates a random string of n alphanumeric characters.
func (r *Rand) AlphaNumericN(n int) string {
	result := make([]byte, n)
	for i := range result {
		result[i] = alphaNumericCharSet[r.IntN(len(alphaNumericCharSet))]
	}
	return string(result)
}

// AlphaNumeric generates a random string of n alphanumeric characters using the default random source.
func AlphaNumeric(n int) string {
	return runtimeRng.AlphaNumericN(n)
}

// SecAlphaNumeric generates a cryptographically secure random string of n alphanumeric characters.
func SecAlphaNumeric(n int) string {
	return cryptoRng.AlphaNumericN(n)
}
