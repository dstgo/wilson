package random

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand_Bool(t *testing.T) {
	r := New(DefaultRng)
	trueCount := 0
	iterations := 10000

	for i := 0; i < iterations; i++ {
		if r.Bool() {
			trueCount++
		}
	}

	ratio := float64(trueCount) / float64(iterations)
	assert.InDelta(t, 0.5, ratio, 0.05, "Bool() should return true and false with roughly equal probability")
}

func TestRand_IntRange(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       int
		max       int
		checkFunc func(t *testing.T, result int)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  10,
			max:  20,
			checkFunc: func(t *testing.T, result int) {
				assert.GreaterOrEqual(t, result, 10)
				assert.Less(t, result, 20)
			},
		},
		{
			name:      "Panic when min > max",
			min:       20,
			max:       10,
			panicFunc: func() { r.IntRange(20, 10) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  10,
			max:  10,
			checkFunc: func(t *testing.T, result int) {
				assert.Equal(t, 10, result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.IntRange(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_Int32Range(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       int32
		max       int32
		checkFunc func(t *testing.T, result int32)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  10,
			max:  20,
			checkFunc: func(t *testing.T, result int32) {
				assert.GreaterOrEqual(t, result, int32(10))
				assert.Less(t, result, int32(20))
			},
		},
		{
			name:      "Panic when min > max",
			min:       20,
			max:       10,
			panicFunc: func() { r.Int32Range(20, 10) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  10,
			max:  10,
			checkFunc: func(t *testing.T, result int32) {
				assert.Equal(t, int32(10), result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.Int32Range(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_Int64Range(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       int64
		max       int64
		checkFunc func(t *testing.T, result int64)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  10,
			max:  20,
			checkFunc: func(t *testing.T, result int64) {
				assert.GreaterOrEqual(t, result, int64(10))
				assert.Less(t, result, int64(20))
			},
		},
		{
			name:      "Panic when min > max",
			min:       20,
			max:       10,
			panicFunc: func() { r.Int64Range(20, 10) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  10,
			max:  10,
			checkFunc: func(t *testing.T, result int64) {
				assert.Equal(t, int64(10), result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.Int64Range(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_UintRange(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       uint
		max       uint
		checkFunc func(t *testing.T, result uint)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  10,
			max:  20,
			checkFunc: func(t *testing.T, result uint) {
				assert.GreaterOrEqual(t, result, uint(10))
				assert.Less(t, result, uint(20))
			},
		},
		{
			name:      "Panic when min > max",
			min:       20,
			max:       10,
			panicFunc: func() { r.UintRange(20, 10) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  10,
			max:  10,
			checkFunc: func(t *testing.T, result uint) {
				assert.Equal(t, uint(10), result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.UintRange(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_Uint32Range(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       uint32
		max       uint32
		checkFunc func(t *testing.T, result uint32)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  10,
			max:  20,
			checkFunc: func(t *testing.T, result uint32) {
				assert.GreaterOrEqual(t, result, uint32(10))
				assert.Less(t, result, uint32(20))
			},
		},
		{
			name:      "Panic when min > max",
			min:       20,
			max:       10,
			panicFunc: func() { r.Uint32Range(20, 10) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  10,
			max:  10,
			checkFunc: func(t *testing.T, result uint32) {
				assert.Equal(t, uint32(10), result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.Uint32Range(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_Uint64Range(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       uint64
		max       uint64
		checkFunc func(t *testing.T, result uint64)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  10,
			max:  20,
			checkFunc: func(t *testing.T, result uint64) {
				assert.GreaterOrEqual(t, result, uint64(10))
				assert.Less(t, result, uint64(20))
			},
		},
		{
			name:      "Panic when min > max",
			min:       20,
			max:       10,
			panicFunc: func() { r.Uint64Range(20, 10) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  10,
			max:  10,
			checkFunc: func(t *testing.T, result uint64) {
				assert.Equal(t, uint64(10), result)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.Uint64Range(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_Float32Range(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       float32
		max       float32
		checkFunc func(t *testing.T, result float32)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  1.0,
			max:  2.0,
			checkFunc: func(t *testing.T, result float32) {
				assert.GreaterOrEqual(t, result, float32(1.0))
				assert.Less(t, result, float32(2.0))
			},
		},
		{
			name:      "Panic when min > max",
			min:       2.0,
			max:       1.0,
			panicFunc: func() { r.Float32Range(2.0, 1.0) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  1.0,
			max:  1.0,
			checkFunc: func(t *testing.T, result float32) {
				assert.Equal(t, float32(1.0), result)
			},
		},
		{
			name:      "Overflow: max - min = +Inf",
			min:       -math.MaxFloat32,
			max:       math.MaxFloat32,
			panicFunc: func() { r.Float32Range(-math.MaxFloat32, math.MaxFloat32) },
			panicMsg:  "overflow: max - min = +Inf",
		},
		{
			name:      "Invalid: max - min = NaN",
			min:       float32(math.NaN()),
			max:       float32(math.NaN()),
			panicFunc: func() { r.Float32Range(float32(math.NaN()), float32(math.NaN())) },
			panicMsg:  "invalid: max - min = NaN",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.Float32Range(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}

func TestRand_Float64Range(t *testing.T) {
	r := New(DefaultRng)
	testCases := []struct {
		name      string
		min       float64
		max       float64
		checkFunc func(t *testing.T, result float64)
		panicFunc func()
		panicMsg  string
	}{
		{
			name: "Normal range",
			min:  1.0,
			max:  2.0,
			checkFunc: func(t *testing.T, result float64) {
				assert.GreaterOrEqual(t, result, 1.0)
				assert.Less(t, result, 2.0)
			},
		},
		{
			name:      "Panic when min > max",
			min:       2.0,
			max:       1.0,
			panicFunc: func() { r.Float64Range(2.0, 1.0) },
			panicMsg:  "min > max",
		},
		{
			name: "Min equals max",
			min:  1.0,
			max:  1.0,
			checkFunc: func(t *testing.T, result float64) {
				assert.Equal(t, 1.0, result)
			},
		},
		{
			name:      "Overflow: max - min = +Inf",
			min:       -math.MaxFloat64,
			max:       math.MaxFloat64,
			panicFunc: func() { r.Float64Range(-math.MaxFloat64, math.MaxFloat64) },
			panicMsg:  "overflow: max - min = +Inf",
		},
		{
			name:      "Invalid: max - min = NaN",
			min:       math.NaN(),
			max:       math.NaN(),
			panicFunc: func() { r.Float64Range(math.NaN(), math.NaN()) },
			panicMsg:  "invalid: max - min = NaN",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panicFunc != nil {
				assert.PanicsWithValue(t, tc.panicMsg, tc.panicFunc)
			} else {
				for i := 0; i < 1000; i++ {
					result := r.Float64Range(tc.min, tc.max)
					tc.checkFunc(t, result)
				}
			}
		})
	}
}
