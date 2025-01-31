package random

// Bool returns a random boolean value.
func Bool() bool {
	return runtimeRng.Uint64()&1 == 1
}

// Int returns a non-negative random int.
func Int() int {
	return runtimeRng.Int()
}

// Int32 returns a random 32-bit integer as an int32.
func Int32() int32 {
	return runtimeRng.Int32()
}

// Int64 returns a random 64-bit integer as an int64.
func Int64() int64 {
	return runtimeRng.Int64()
}

// Uint returns a random unsigned int.
func Uint() uint {
	return runtimeRng.Uint()
}

// Uint32 returns a random 32-bit unsigned integer as a uint32.
func Uint32() uint32 {
	return runtimeRng.Uint32()
}

// Uint64 returns a random 64-bit unsigned integer as a uint64.
func Uint64() uint64 {
	return runtimeRng.Uint64()
}

// IntN returns a random integer in the range [0, max).
func IntN(max int) int {
	return runtimeRng.IntN(max)
}

// Int32N returns a random 32-bit integer in the range [0, max).
func Int32N(max int32) int32 {
	return runtimeRng.Int32N(max)
}

// Int64N returns a random 64-bit integer in the range [0, max).
func Int64N(max int64) int64 {
	return runtimeRng.Int64N(max)
}

// UintN returns a random unsigned integer in the range [0, max).
func UintN(max uint) uint {
	return runtimeRng.UintN(max)
}

// Uint32N returns a random 32-bit unsigned integer in the range [0, max).
func Uint32N(max uint32) uint32 {
	return runtimeRng.Uint32N(max)
}

// Uint64N returns a random 64-bit unsigned integer in the range [0, max).
func Uint64N(max uint64) uint64 {
	return runtimeRng.Uint64N(max)
}

// Float32 returns a random float32 in the range [0.0, 1.0).
func Float32() float32 {
	return runtimeRng.Float32()
}

// Float64 returns a random float64 in the range [0.0, 1.0).
func Float64() float64 {
	return runtimeRng.Float64()
}

// IntRange returns a random integer in the range [min, max).
func IntRange(min, max int) int {
	return runtimeRng.IntRange(min, max)
}

// Int32Range returns a random 32-bit integer in the range [min, max).
func Int32Range(min, max int32) int32 {
	return runtimeRng.Int32Range(min, max)
}

// Int64Range returns a random 64-bit integer in the range [min, max).
func Int64Range(min, max int64) int64 {
	return runtimeRng.Int64Range(min, max)
}

// UintRange returns a random unsigned integer in the range [min, max).
func UintRange(min, max uint) uint {
	return runtimeRng.UintRange(min, max)
}

// Uint32Range returns a random 32-bit unsigned integer in the range [min, max).
func Uint32Range(min, max uint32) uint32 {
	return runtimeRng.Uint32Range(min, max)
}

// Uint64Range returns a random 64-bit unsigned integer in the range [min, max).
func Uint64Range(min, max uint64) uint64 {
	return runtimeRng.Uint64Range(min, max)
}

// Float32Range returns a random float32 in the range [min, max).
func Float32Range(min, max float32) float32 {
	return runtimeRng.Float32Range(min, max)
}

// Float64Range returns a random float64 in the range [min, max).
func Float64Range(min, max float64) float64 {
	return runtimeRng.Float64Range(min, max)
}

// SecBool returns a secure random boolean value.
func SecBool() bool {
	return runtimeRng.Uint64()&1 == 1
}

// SecInt returns a secure random non-negative int.
func SecInt() int {
	return cryptoRng.Int()
}

// SecInt32 returns a secure random 32-bit integer.
func SecInt32() int32 {
	return cryptoRng.Int32()
}

// SecInt64 returns a secure random 64-bit integer.
func SecInt64() int64 {
	return cryptoRng.Int64()
}

// SecUint returns a secure random unsigned int.
func SecUint() uint {
	return cryptoRng.Uint()
}

// SecUint32 returns a secure random 32-bit unsigned integer.
func SecUint32() uint32 {
	return cryptoRng.Uint32()
}

// SecUint64 returns a secure random 64-bit unsigned integer.
func SecUint64() uint64 {
	return cryptoRng.Uint64()
}

// SecFloat32 returns a secure random float32 in the range [0.0, 1.0).
func SecFloat32() float32 {
	return cryptoRng.Float32()
}

// SecFloat64 returns a secure random float64 in the range [0.0, 1.0).
func SecFloat64() float64 {
	return cryptoRng.Float64()
}

// SecIntN returns a secure random integer in the range [0, n).
func SecIntN(n int) int {
	return cryptoRng.IntN(n)
}

// SecInt32N returns a secure random 32-bit integer in the range [0, n).
func SecInt32N(n int32) int32 {
	return cryptoRng.Int32N(n)
}

// SecInt64N returns a secure random 64-bit integer in the range [0, n).
func SecInt64N(n int64) int64 {
	return cryptoRng.Int64N(n)
}

// SecUintN returns a secure random unsigned integer in the range [0, n).
func SecUintN(n uint) uint {
	return cryptoRng.UintN(n)
}

// SecUint32N returns a secure random 32-bit unsigned integer in the range [0, n).
func SecUint32N(n uint32) uint32 {
	return cryptoRng.Uint32N(n)
}

// SecUint64N returns a secure random 64-bit unsigned integer in the range [0, n).
func SecUint64N(max uint64) uint64 {
	return cryptoRng.Uint64N(max)
}

// SecIntRange returns a secure random integer in the range [min, max).
func SecIntRange(min, max int) int {
	return cryptoRng.IntRange(min, max)
}

// SecInt32Range returns a secure random 32-bit integer in the range [min, max).
func SecInt32Range(min, max int32) int32 {
	return cryptoRng.Int32Range(min, max)
}

// SecInt64Range returns a secure random 64-bit integer in the range [min, max).
func SecInt64Range(min, max int64) int64 {
	return cryptoRng.Int64Range(min, max)
}

// SecUintRange returns a secure random unsigned integer in the range [min, max).
func SecUintRange(min, max uint) uint {
	return cryptoRng.UintRange(min, max)
}

// SecUint32Range returns a secure random 32-bit unsigned integer in the range [min, max).
func SecUint32Range(min, max uint32) uint32 {
	return cryptoRng.Uint32Range(min, max)
}

// SecUint64Range returns a secure random 64-bit unsigned integer in the range [min, max).
func SecUint64Range(min, max uint64) uint64 {
	return cryptoRng.Uint64Range(min, max)
}

// SecFloat32Range returns a secure random float32 in the range [min, max).
func SecFloat32Range(min, max float32) float32 {
	return cryptoRng.Float32Range(min, max)
}

// SecFloat64Range returns a secure random float64 in the range [min, max).
func SecFloat64Range(min, max float64) float64 {
	return cryptoRng.Float64Range(min, max)
}
