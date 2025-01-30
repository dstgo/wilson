package random

// Bool returns a random boolean value.
func Bool() bool {
	return DefaultRand.Uint64()&1 == 1
}

// Int returns a non-negative random int.
func Int() int {
	return DefaultRand.Int()
}

// Int32 returns a random 32-bit integer as an int32.
func Int32() int32 {
	return DefaultRand.Int32()
}

// Int64 returns a random 64-bit integer as an int64.
func Int64() int64 {
	return DefaultRand.Int64()
}

// Uint returns a random unsigned int.
func Uint() uint {
	return DefaultRand.Uint()
}

// Uint32 returns a random 32-bit unsigned integer as a uint32.
func Uint32() uint32 {
	return DefaultRand.Uint32()
}

// Uint64 returns a random 64-bit unsigned integer as a uint64.
func Uint64() uint64 {
	return DefaultRand.Uint64()
}

// IntN returns a random integer in the range [0, max).
func IntN(max int) int {
	return DefaultRand.IntN(max)
}

// Int32N returns a random 32-bit integer in the range [0, max).
func Int32N(max int32) int32 {
	return DefaultRand.Int32N(max)
}

// Int64N returns a random 64-bit integer in the range [0, max).
func Int64N(max int64) int64 {
	return DefaultRand.Int64N(max)
}

// UintN returns a random unsigned integer in the range [0, max).
func UintN(max uint) uint {
	return DefaultRand.UintN(max)
}

// Uint32N returns a random 32-bit unsigned integer in the range [0, max).
func Uint32N(max uint32) uint32 {
	return DefaultRand.Uint32N(max)
}

// Uint64N returns a random 64-bit unsigned integer in the range [0, max).
func Uint64N(max uint64) uint64 {
	return DefaultRand.Uint64N(max)
}

// Float32 returns a random float32 in the range [0.0, 1.0).
func Float32() float32 {
	return DefaultRand.Float32()
}

// Float64 returns a random float64 in the range [0.0, 1.0).
func Float64() float64 {
	return DefaultRand.Float64()
}

// SecBool returns a secure random boolean value.
func SecBool() bool {
	return DefaultRand.Uint64()&1 == 1
}

// SecInt returns a secure random non-negative int.
func SecInt() int {
	return CryptoRand.Int()
}

// SecInt32 returns a secure random 32-bit integer.
func SecInt32() int32 {
	return CryptoRand.Int32()
}

// SecInt64 returns a secure random 64-bit integer.
func SecInt64() int64 {
	return CryptoRand.Int64()
}

// SecUint returns a secure random unsigned int.
func SecUint() uint {
	return CryptoRand.Uint()
}

// SecUint32 returns a secure random 32-bit unsigned integer.
func SecUint32() uint32 {
	return CryptoRand.Uint32()
}

// SecUint64 returns a secure random 64-bit unsigned integer.
func SecUint64() uint64 {
	return CryptoRand.Uint64()
}

// SecFloat32 returns a secure random float32 in the range [0.0, 1.0).
func SecFloat32() float32 {
	return CryptoRand.Float32()
}

// SecFloat64 returns a secure random float64 in the range [0.0, 1.0).
func SecFloat64() float64 {
	return CryptoRand.Float64()
}

// SecIntN returns a secure random integer in the range [0, n).
func SecIntN(n int) int {
	return CryptoRand.IntN(n)
}

// SecInt32N returns a secure random 32-bit integer in the range [0, n).
func SecInt32N(n int32) int32 {
	return CryptoRand.Int32N(n)
}

// SecInt64N returns a secure random 64-bit integer in the range [0, n).
func SecInt64N(n int64) int64 {
	return CryptoRand.Int64N(n)
}

// SecUintN returns a secure random unsigned integer in the range [0, n).
func SecUintN(n uint) uint {
	return CryptoRand.UintN(n)
}

// SecUint32N returns a secure random 32-bit unsigned integer in the range [0, n).
func SecUint32N(n uint32) uint32 {
	return CryptoRand.Uint32N(n)
}

// SecUint64N returns a secure random 64-bit unsigned integer in the range [0, n).
func SecUint64N(max uint64) uint64 {
	return CryptoRand.Uint64N(max)
}
