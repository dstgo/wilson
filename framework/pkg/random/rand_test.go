package random

import (
	"math"
	"testing"
)

func TestNew(t *testing.T) {
	r := New(NewChaCha8())
	if r == nil {
		t.Fatal("New should return a non-nil Rand")
	}
}

func TestBool(t *testing.T) {
	r := New(NewChaCha8())
	results := make(map[bool]int)
	for i := 0; i < 10000; i++ {
		results[r.Bool()]++
	}
	if len(results) != 2 {
		t.Fatal("Bool should return both true and false")
	}
	if results[true] < 4000 || results[true] > 6000 {
		t.Fatalf("Bool distribution seems biased: %v", results)
	}
}

func TestIntRange(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   int
		max   int
		panic bool
	}{
		{"Normal", 0, 10, false},
		{"Negative", -10, 0, false},
		{"Min > Max", 10, 0, true},
		{"Overflow1", -1, math.MaxInt, true},
		{"OVerflow2", math.MinInt, math.MaxInt, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("IntRange(%d, %d) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.IntRange(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("IntRange(%d, %d) returned %d, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestInt32Range(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   int32
		max   int32
		panic bool
	}{
		{"Normal", 0, 10, false},
		{"Negative", -10, 0, false},
		{"Min > Max", 10, 0, true},
		{"Overflow1", -1, math.MaxInt32, true},
		{"Overflow2", math.MinInt32, math.MaxInt32, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Int32Range(%d, %d) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.Int32Range(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("Int32Range(%d, %d) returned %d, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestInt64Range(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   int64
		max   int64
		panic bool
	}{
		{"Normal", 0, 10, false},
		{"Negative", -10, 0, false},
		{"Min > Max", 10, 0, true},
		{"Overflow1", -1, math.MaxInt64, true},
		{"Overflow2", math.MinInt64, math.MaxInt64, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Int64Range(%d, %d) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.Int64Range(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("Int64Range(%d, %d) returned %d, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestUintRange(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   uint
		max   uint
		panic bool
	}{
		{"Normal", 0, 10, false},
		{"Full Range", 0, math.MaxUint, false},
		{"Min > Max", 10, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("UintRange(%d, %d) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.UintRange(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("UintRange(%d, %d) returned %d, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestUint32Range(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   uint32
		max   uint32
		panic bool
	}{
		{"Normal", 0, 10, false},
		{"Full Range", 0, math.MaxUint32, false},
		{"Min > Max", 10, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Uint32Range(%d, %d) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.Uint32Range(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("Uint32Range(%d, %d) returned %d, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestUint64Range(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   uint64
		max   uint64
		panic bool
	}{
		{"Normal", 0, 10, false},
		{"Full Range", 0, math.MaxUint64, false},
		{"Min > Max", 10, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Uint64Range(%d, %d) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.Uint64Range(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("Uint64Range(%d, %d) returned %d, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestFloat32Range(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   float32
		max   float32
		panic bool
	}{
		{"Normal", 0, 1, false},
		{"Negative", -1, 0, false},
		{"Min > Max", 1, 0, true},
		{"Overflow", -math.MaxFloat32, math.MaxFloat32, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Float32Range(%f, %f) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.Float32Range(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("Float32Range(%f, %f) returned %f, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}

func TestFloat64Range(t *testing.T) {
	r := New(NewChaCha8())

	testCases := []struct {
		name  string
		min   float64
		max   float64
		panic bool
	}{
		{"Normal", 0, 1, false},
		{"Negative", -1, 0, false},
		{"Min > Max", 1, 0, true},
		{"Overflow", -math.MaxFloat64, math.MaxFloat64, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.panic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Float64Range(%f, %f) should panic", tc.min, tc.max)
					}
				}()
			}

			for i := 0; i < 1000; i++ {
				result := r.Float64Range(tc.min, tc.max)
				if result < tc.min || result >= tc.max {
					t.Errorf("Float64Range(%f, %f) returned %f, which is out of range", tc.min, tc.max, result)
				}
			}
		})
	}
}
