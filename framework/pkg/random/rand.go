package random

import (
	// #nosec G404 (CWE-338): Use of weak random number generator
	"math"
	"math/rand/v2"
)

// Source is an interface for generating random numbers.
// #nosec G404 (CWE-338): Use of weak random number generator
type Source interface {
	rand.Source
}

// New returns a new Rand using the provided source.
// #nosec G404 (CWE-338): Use of weak random number generator
func New(source Source) *Rand {
	return &Rand{rand.New(source)}
}

// Rand is an enhance number generator for *math/rand/v2.Rand.
// #nosec G404 (CWE-338): Use of weak random number generator
type Rand struct {
	*rand.Rand
}

// Bool returns a random boolean value.
func (r *Rand) Bool() bool {
	return r.Uint64()&1 == 1
}

// IntRange returns a random int in the range [min, max).
func (r *Rand) IntRange(min, max int) int {
	if min > max {
		panic("min > max")
	}

	if min == max {
		return min
	}

	diff := uint(max) - uint(min)

	return int(uint(min) + r.UintN(diff))
}

// Int32Range returns a random int32 in the range [min, max).
func (r *Rand) Int32Range(min, max int32) int32 {
	if min > max {
		panic("min > max")
	}

	if min == max {
		return min
	}

	diff := uint32(max) - uint32(min)

	return int32(uint32(min) + r.Uint32N(diff))
}

// Int64Range returns a random int64 in the range [min, max).
func (r *Rand) Int64Range(min, max int64) int64 {
	if min > max {
		panic("min > max")
	}

	if min == max {
		return min
	}

	diff := uint64(max) - uint64(min)

	return int64(uint64(min) + r.Uint64N(diff))
}

// UintRange returns a random uint in the range [min, max).
func (r *Rand) UintRange(min, max uint) uint {
	if min > max {
		panic("min > max")
	}

	if min == max {
		return min
	}

	return min + r.UintN(max-min)
}

// Uint32Range returns a random uint32 in the range [min, max).
func (r *Rand) Uint32Range(min, max uint32) uint32 {
	if min > max {
		panic("min > max")
	}

	if min == max {
		return min
	}

	return min + r.Uint32N(max-min)
}

// Uint64Range returns a random uint64 in the range [min, max).
func (r *Rand) Uint64Range(min, max uint64) uint64 {
	if min > max {
		panic("min > max")
	}

	if min == max {
		return min
	}

	return min + r.Uint64N(max-min)
}

// Float32Range returns a random float32 in the range [min, max).
func (r *Rand) Float32Range(min, max float32) float32 {
	if min > max {
		panic("min > max")
	}

	diff := float64(max - min)

	if math.IsInf(diff, 0) {
		panic("overflow: max - min = +Inf")
	}

	if math.IsNaN(diff) {
		panic("invalid: max - min = NaN")
	}

	return min + r.Float32()*(max-min)
}

// Float64Range returns a random float64 in the range [min, max).
func (r *Rand) Float64Range(min, max float64) float64 {
	if min > max {
		panic("min > max")
	}

	diff := max - min

	if math.IsInf(diff, 0) {
		panic("overflow: max - min = +Inf")
	}

	if math.IsNaN(diff) {
		panic("invalid: max - min = NaN")
	}

	return min + r.Float64()*(max-min)
}
