package random

import (
	"math/rand/v2"
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestBytesN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := BytesN(rng, 10, 0, 255)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, b := range result {
		if b < 0 || b > 255 {
			t.Errorf("Byte out of range: %d", b)
		}
	}
}

func TestBytes(t *testing.T) {
	result := Bytes(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
}

func TestBytesSeq(t *testing.T) {
	result := BytesSeq(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
}

func TestRunesN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := RunesN(rng, 10, 'A', 'Z')
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if r < 'A' || r > 'Z' {
			t.Errorf("Rune out of range: %c", r)
		}
	}
}

func TestRunes(t *testing.T) {
	result := Runes(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
}

func TestRunesSeq(t *testing.T) {
	result := RunesSeq(10)
	n := utf8.RuneCountInString(result)
	if n != 10 {
		t.Errorf("Expected length 10, got %d", n)
	}
}

func TestDigitsN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := DigitsN(rng, 10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if r < '0' || r > '9' {
			t.Errorf("Character is not a digit: %c", r)
		}
	}
}

func TestDigits(t *testing.T) {
	result := Digits(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsDigit(r) {
			t.Errorf("Character is not a digit: %c", r)
		}
	}
}

func TestLowerN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := LowerN(rng, 10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsLower(r) {
			t.Errorf("Character is not lowercase: %c", r)
		}
	}
}

func TestLower(t *testing.T) {
	result := Lower(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsLower(r) {
			t.Errorf("Character is not lowercase: %c", r)
		}
	}
}

func TestUpperN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := UpperN(rng, 10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsUpper(r) {
			t.Errorf("Character is not uppercase: %c", r)
		}
	}
}

func TestUpper(t *testing.T) {
	result := Upper(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsUpper(r) {
			t.Errorf("Character is not uppercase: %c", r)
		}
	}
}

func TestLettersN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := LettersN(rng, 10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsLetter(r) {
			t.Errorf("Character is not a letter: %c", r)
		}
	}
}

func TestLetters(t *testing.T) {
	result := Letters(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsLetter(r) {
			t.Errorf("Character is not a letter: %c", r)
		}
	}
}

func TestAlphaNumericN(t *testing.T) {
	rng := New(rand.NewPCG(1, 2))
	result := AlphaNumericN(rng, 10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			t.Errorf("Character is not alphanumeric: %c", r)
		}
	}
}

func TestAlphaNumeric(t *testing.T) {
	result := AlphaNumeric(10)
	if len(result) != 10 {
		t.Errorf("Expected length 10, got %d", len(result))
	}
	for _, r := range result {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			t.Errorf("Character is not alphanumeric: %c", r)
		}
	}
}
