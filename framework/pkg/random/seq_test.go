package random

import (
	"testing"
	"unicode"
)

func TestRand_BytesN(t *testing.T) {
	r := runtimeRng
	testCases := []struct {
		n, low, high byte
	}{
		{10, 0, 255},
		{20, 65, 90},
		{30, 97, 122},
	}

	for _, tc := range testCases {
		result := r.BytesN(int(tc.n), tc.low, tc.high)
		if len(result) != int(tc.n) {
			t.Errorf("Rand.BytesN(%d, %d, %d): expected length %d, got %d", tc.n, tc.low, tc.high, tc.n, len(result))
		}
		for _, b := range result {
			if b < tc.low || b > tc.high {
				t.Errorf("Rand.BytesN(%d, %d, %d): byte %d out of range", tc.n, tc.low, tc.high, b)
			}
		}
	}
}

func TestRand_RunesN(t *testing.T) {
	r := runtimeRng
	testCases := []struct {
		n          int
		low, high  rune
		checkRange func(rune) bool
	}{
		{10, 'A', 'Z', unicode.IsUpper},
		{20, 'a', 'z', unicode.IsLower},
		{30, '0', '9', unicode.IsDigit},
	}

	for _, tc := range testCases {
		result := r.RunesN(tc.n, tc.low, tc.high)
		if len(result) != tc.n {
			t.Errorf("Rand.RunesN(%d, %d, %d): expected length %d, got %d", tc.n, tc.low, tc.high, tc.n, len(result))
		}
		for _, r := range result {
			if r < tc.low || r > tc.high || !tc.checkRange(r) {
				t.Errorf("Rand.RunesN(%d, %d, %d): rune %d out of range or doesn't match expected type", tc.n, tc.low, tc.high, r)
			}
		}
	}
}

func TestRand_DigitsN(t *testing.T) {
	r := runtimeRng
	testCases := []int{5, 10, 15}

	for _, n := range testCases {
		result := r.DigitsN(n)
		if len(result) != n {
			t.Errorf("Rand.DigitsN(%d): expected length %d, got %d", n, n, len(result))
		}
		for _, r := range result {
			if !unicode.IsDigit(r) {
				t.Errorf("Rand.DigitsN(%d): character %c is not a digit", n, r)
			}
		}
	}
}

func TestRand_LowerN(t *testing.T) {
	r := runtimeRng
	testCases := []int{5, 10, 15}

	for _, n := range testCases {
		result := r.LowerN(n)
		if len(result) != n {
			t.Errorf("Rand.LowerN(%d): expected length %d, got %d", n, n, len(result))
		}
		for _, r := range result {
			if !unicode.IsLower(r) {
				t.Errorf("Rand.LowerN(%d): character %c is not lowercase", n, r)
			}
		}
	}
}

func TestRand_UpperN(t *testing.T) {
	r := runtimeRng
	testCases := []int{5, 10, 15}

	for _, n := range testCases {
		result := r.UpperN(n)
		if len(result) != n {
			t.Errorf("Rand.UpperN(%d): expected length %d, got %d", n, n, len(result))
		}
		for _, r := range result {
			if !unicode.IsUpper(r) {
				t.Errorf("Rand.UpperN(%d): character %c is not uppercase", n, r)
			}
		}
	}
}

func TestRand_LettersN(t *testing.T) {
	r := runtimeRng
	testCases := []int{5, 10, 15}

	for _, n := range testCases {
		result := r.LettersN(n)
		if len(result) != n {
			t.Errorf("Rand.LettersN(%d): expected length %d, got %d", n, n, len(result))
		}
		for _, r := range result {
			if !unicode.IsLetter(r) {
				t.Errorf("Rand.LettersN(%d): character %c is not a letter", n, r)
			}
		}
	}
}

func TestRand_AlphaNumericN(t *testing.T) {
	r := runtimeRng
	testCases := []int{5, 10, 15}

	for _, n := range testCases {
		result := r.AlphaNumericN(n)
		if len(result) != n {
			t.Errorf("Rand.AlphaNumericN(%d): expected length %d, got %d", n, n, len(result))
		}
		for _, r := range result {
			if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
				t.Errorf("Rand.AlphaNumericN(%d): character %c is not alphanumeric", n, r)
			}
		}
	}
}
