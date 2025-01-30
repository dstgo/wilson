package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChaCha8(t *testing.T) {
	rng := NewChaCha8()
	assert.NotNil(t, rng, "NewChaCha8 should return a non-nil random number generator")

	// Generate a sequence of numbers and ensure they're not all the same
	values := make([]int64, 1000)
	for i := range values {
		values[i] = rng.Int64()
	}
	assert.NotEqual(t, values[0], values[1], "Generated values should not be identical")
}

func TestNewPCG(t *testing.T) {
	rng := NewPCG()
	assert.NotNil(t, rng, "NewPCG should return a non-nil random number generator")

	// Generate a sequence of numbers and ensure they're not all the same
	values := make([]uint64, 1000)
	for i := range values {
		values[i] = rng.Uint64()
	}
	assert.NotEqual(t, values[0], values[1], "Generated values should not be identical")
}

func TestNewZipf(t *testing.T) {
	rng := NewZipf()
	assert.NotNil(t, rng, "NewZipf should return a non-nil random number generator")

	// Generate a sequence of numbers
	values := make([]uint64, 10000)
	for i := range values {
		values[i] = rng.Uint64()
	}

	// Check that values follow Zipf distribution (very basic check)
	// In a Zipf distribution, lower values should occur more frequently
	counts := make(map[uint64]int)
	for _, v := range values {
		counts[v]++
	}

	var prevCount int
	for i := uint64(1); i <= 10; i++ {
		count := counts[i]
		if i > 1 {
			assert.GreaterOrEqual(t, prevCount, count, "In Zipf distribution, lower values should occur more frequently")
		}
		prevCount = count
	}
}
