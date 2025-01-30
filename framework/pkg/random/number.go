package random

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	"math"
)

// Int returns a non-negative random int.
// #nosec G404 (CWE-338): Use of weak random number generator
func Int() int {
	return DefaultRand.Int()
}

// Int32 returns a random 32-bit integer as an int32.
// #nosec G404 (CWE-338): Use of weak random number generator
func Int32() int32 {
	return DefaultRand.Int32()
}

// Int64 returns a random 64-bit integer as an int64.
// #nosec G404 (CWE-338): Use of weak random number generator
func Int64() int64 {
	return DefaultRand.Int64()
}

// Uint returns a random unsigned int.
// #nosec G404 (CWE-338): Use of weak random number generator
func Uint() uint {
	return DefaultRand.Uint()
}

// Uint32 returns a random 32-bit unsigned integer as a uint32.
// #nosec G404 (CWE-338): Use of weak random number generator
func Uint32() uint32 {
	return DefaultRand.Uint32()
}

// Uint64 returns a random 64-bit unsigned integer as a uint64.
// #nosec G404 (CWE-338): Use of weak random number generator
func Uint64() uint64 {
	return DefaultRand.Uint64()
}

// IntN returns a random integer in the range [0, max).
// #nosec G404 (CWE-338): Use of weak random number generator
func IntN(max int) int {
	return DefaultRand.IntN(max)
}

// Int32N returns a random 32-bit integer in the range [0, max).
// #nosec G404 (CWE-338): Use of weak random number generator
func Int32N(max int32) int32 {
	return DefaultRand.Int32N(max)
}

// Int64N returns a random 64-bit integer in the range [0, max).
// #nosec G404 (CWE-338): Use of weak random number generator
func Int64N(max int64) int64 {
	return DefaultRand.Int64N(max)
}

// UintN returns a random unsigned integer in the range [0, max).
// #nosec G404 (CWE-338): Use of weak random number generator
func UintN(max uint) uint {
	return DefaultRand.UintN(max)
}

// Uint32N returns a random 32-bit unsigned integer in the range [0, max).
// #nosec G404 (CWE-338): Use of weak random number generator
func Uint32N(max uint32) uint32 {
	return DefaultRand.Uint32N(max)
}

// Uint64N returns a random 64-bit unsigned integer in the range [0, max).
// #nosec G404 (CWE-338): Use of weak random number generator
func Uint64N(max uint64) uint64 {
	return DefaultRand.Uint64N(max)
}

// Float32 returns a random float32 in the range [0.0, 1.0).
// #nosec G404 (CWE-338): Use of weak random number generator
func Float32() float32 {
	return DefaultRand.Float32()
}

// Float64 returns a random float64 in the range [0.0, 1.0).
// #nosec G404 (CWE-338): Use of weak random number generator
func Float64() float64 {
	return DefaultRand.Float64()
}

// SecInt returns a secure random non-negative int.
func SecInt() int {
	return int(uint(SecUint64()) << 1 >> 1)
}

// SecInt32 returns a secure random 32-bit integer.
func SecInt32() int32 {
	var buf [binary.MaxVarintLen32]byte
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		return NewChaCha8().Int32()
	}
	return int32(binary.BigEndian.Uint32(buf[:]) &^ (1 << 31))
}

// SecInt64 returns a secure random 64-bit integer.
func SecInt64() int64 {
	var buf [binary.MaxVarintLen64]byte
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		return NewChaCha8().Int64()
	}
	return int64(binary.BigEndian.Uint64(buf[:]) &^ (1 << 63))
}

// SecUint returns a secure random unsigned int.
func SecUint() uint {
	return uint(SecUint64())
}

// SecUint32 returns a secure random 32-bit unsigned integer.
func SecUint32() uint32 {
	var buf [binary.MaxVarintLen32]byte
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		return NewChaCha8().Uint32()
	}
	return binary.BigEndian.Uint32(buf[:])
}

// SecUint64 returns a secure random 64-bit unsigned integer.
func SecUint64() uint64 {
	var buf [binary.MaxVarintLen64]byte
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		return NewChaCha8().Uint64()
	}
	return binary.BigEndian.Uint64(buf[:])
}

// SecFloat32 returns a secure random float32 in the range [0.0, 1.0).
func SecFloat32() float32 {
	// There are exactly 1<<24 float32s in [0,1).
	return float32(SecUint32()<<8>>8) / (1 << 24)
}

// SecFloat64 returns a secure random float64 in the range [0.0, 1.0).
func SecFloat64() float64 {
	// There are exactly 1<<53 float64s in [0,1).
	return float64(SecUint64()<<11>>11) / (1 << 53)
}

// SecIntN returns a secure random integer in the range [0, max).
func SecIntN(max int) int {
	if max <= 0 {
		panic("max must be positive")
	}

	if max <= math.MaxInt32 {
		return int(SecInt32N(int32(max)))
	}
	return int(SecInt64N(int64(max)))
}

// SecInt32N returns a secure random 32-bit integer in the range [0, max).
func SecInt32N(max int32) int32 {
	if max <= 0 {
		panic("max must be positive")
	}

	// Create a mask that is the next power of 2 minus 1
	mask := uint32(max - 1)
	// Fill in all bits below the highest set bit
	// This creates a mask with all bits set to 1 from the highest bit down
	mask |= mask >> 1  // Set all bits below the highest bit
	mask |= mask >> 2  // Set all bits below the 2nd highest bit
	mask |= mask >> 4  // Set all bits below the 4th highest bit
	mask |= mask >> 8  // Set all bits below the 8th highest bit
	mask |= mask >> 16 // Set all remaining bits

	for {
		var buf [binary.MaxVarintLen32]byte
		_, err := cryptorand.Read(buf[:])
		if err != nil {
			return NewChaCha8().Int32N(max)
		}

		// Apply the mask to ensure the number is within the desired range
		n := int32(binary.BigEndian.Uint32(buf[:]) & mask)
		// Reject if n is outside the desired range
		if n < max {
			return n
		}
	}
}

// SecInt64N returns a secure random 64-bit integer in the range [0, max).
func SecInt64N(max int64) int64 {
	if max <= 0 {
		panic("max must be positive")
	}

	// Create a mask that is the next power of 2 minus 1
	mask := uint64(max - 1)
	// Fill in all bits below the highest set bit
	// This creates a mask with all bits set to 1 from the highest bit down
	mask |= mask >> 1  // Set all bits below the highest bit
	mask |= mask >> 2  // Set all bits below the 2nd highest bit
	mask |= mask >> 4  // Set all bits below the 4th highest bit
	mask |= mask >> 8  // Set all bits below the 8th highest bit
	mask |= mask >> 16 // Set all bits below the 16th highest bit
	mask |= mask >> 32 // Set all remaining bits

	for {
		var buf [binary.MaxVarintLen64]byte
		_, err := cryptorand.Read(buf[:])
		if err != nil {
			return NewChaCha8().Int64N(max)
		}
		// Apply the mask to ensure the number is within the desired range
		n := int64(binary.BigEndian.Uint64(buf[:]) & mask)
		// Reject if n is outside the desired range
		if n < max {
			return n
		}
	}
}

// SecUintN returns a secure random unsigned integer in the range [0, max).
func SecUintN(max uint) uint {
	if max == 0 {
		panic("max must be positive")
	}

	if max <= math.MaxUint32 {
		return uint(SecUint32N(uint32(max)))
	}
	return uint(SecUint64N(uint64(max)))
}

// SecUint32N returns a secure random 32-bit unsigned integer in the range [0, max).
func SecUint32N(max uint32) uint32 {
	if max == 0 {
		panic("max must be positive")
	}

	// Create a mask that is the next power of 2 minus 1
	mask := max - 1
	// Fill in all bits below the highest set bit
	// This creates a mask with all bits set to 1 from the highest bit down
	mask |= mask >> 1  // Set all bits below the highest bit
	mask |= mask >> 2  // Set all bits below the 2nd highest bit
	mask |= mask >> 4  // Set all bits below the 4th highest bit
	mask |= mask >> 8  // Set all bits below the 8th highest bit
	mask |= mask >> 16 // Set all remaining bits

	for {
		var buf [binary.MaxVarintLen32]byte
		_, err := cryptorand.Read(buf[:])
		if err != nil {
			return NewChaCha8().Uint32N(max)
		}
		// Apply the mask to ensure the number is within the desired range
		n := binary.BigEndian.Uint32(buf[:]) & mask
		// Reject if n is outside the desired range
		if n < max {
			return n
		}
	}
}

// SecUint64N returns a secure random 64-bit unsigned integer in the range [0, max).
func SecUint64N(max uint64) uint64 {
	if max == 0 {
		panic("max must be positive")
	}

	// Create a mask that is the next power of 2 minus 1
	mask := max - 1
	// Fill in all bits below the highest set bit
	// This creates a mask with all bits set to 1 from the highest bit down
	mask |= mask >> 1  // Set all bits below the highest bit
	mask |= mask >> 2  // Set all bits below the 2nd highest bit
	mask |= mask >> 4  // Set all bits below the 4th highest bit
	mask |= mask >> 8  // Set all bits below the 8th highest bit
	mask |= mask >> 16 // Set all bits below the 16th highest bit
	mask |= mask >> 32 // Set all remaining bits

	for {
		var buf [binary.MaxVarintLen64]byte
		_, err := cryptorand.Read(buf[:])
		if err != nil {
			return NewChaCha8().Uint64N(max)
		}
		// Apply the mask to ensure the number is within the desired range
		n := binary.BigEndian.Uint64(buf[:]) & mask
		// Reject if n is outside the desired range
		if n < max {
			return n
		}
	}
}
