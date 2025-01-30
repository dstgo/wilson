package random

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := Int(); n < 0 {
			t.Errorf("Int() = %d; want non-negative", n)
		}
	}
}

func TestInt32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := Int32(); n < math.MinInt32 || n > math.MaxInt32 {
			t.Errorf("Int32() = %d; want value in range [%d, %d]", n, math.MinInt32, math.MaxInt32)
		}
	}
}

func TestInt64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := Int64(); n < math.MinInt64 || n > math.MaxInt64 {
			t.Errorf("Int64() = %d; want value in range [%d, %d]", n, math.MinInt64, math.MaxInt64)
		}
	}
}

func TestUint(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := Uint(); n < 0 {
			t.Errorf("Uint() = %d; want non-negative", n)
		}
	}
}

func TestUint32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := Uint32(); n > math.MaxUint32 {
			t.Errorf("Uint32() = %d; want value in range [0, %d]", n, math.MaxUint32)
		}
	}
}

func TestUint64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := Uint64(); n > math.MaxUint64 {
			t.Errorf("Uint64() = %d; want value in range [0, %d]", n, uint64(math.MaxUint64))
		}
	}
}

func TestSecInt(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := SecInt(); n < 0 {
			t.Errorf("SecInt() = %d; want non-negative", n)
		}
	}
}

func TestSecInt32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := SecInt32(); n < 0 {
			t.Errorf("SecInt32() = %d; want non-negative", n)
		}
	}
}

func TestSecInt64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if n := SecInt64(); n < 0 {
			t.Errorf("SecInt64() = %d; want non-negative", n)
		}
	}
}

func TestSecUint(t *testing.T) {
	for i := 0; i < 1000; i++ {
		SecUint() // Just ensure it doesn't panic
	}
}

func TestSecUint32(t *testing.T) {
	for i := 0; i < 1000; i++ {
		SecUint32() // Just ensure it doesn't panic
	}
}

func TestSecUint64(t *testing.T) {
	for i := 0; i < 1000; i++ {
		SecUint64() // Just ensure it doesn't panic
	}
}

func TestIntN(t *testing.T) {
	upperBound := 100
	for i := 0; i < 1000; i++ {
		if n := IntN(upperBound); n < 0 || n >= upperBound {
			t.Errorf("IntN(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestInt32N(t *testing.T) {
	upperBound := int32(100)
	for i := 0; i < 1000; i++ {
		if n := Int32N(upperBound); n < 0 || n >= upperBound {
			t.Errorf("Int32N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestInt64N(t *testing.T) {
	upperBound := int64(100)
	for i := 0; i < 1000; i++ {
		if n := Int64N(upperBound); n < 0 || n >= upperBound {
			t.Errorf("Int64N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestUintN(t *testing.T) {
	upperBound := uint(100)
	for i := 0; i < 1000; i++ {
		if n := UintN(upperBound); n >= upperBound {
			t.Errorf("UintN(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestUint32N(t *testing.T) {
	upperBound := uint32(100)
	for i := 0; i < 1000; i++ {
		if n := Uint32N(upperBound); n >= upperBound {
			t.Errorf("Uint32N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestUint64N(t *testing.T) {
	upperBound := uint64(100)
	for i := 0; i < 1000; i++ {
		if n := Uint64N(upperBound); n >= upperBound {
			t.Errorf("Uint64N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestSecIntN(t *testing.T) {
	upperBound := 100
	for i := 0; i < 1000; i++ {
		if n := SecIntN(upperBound); n < 0 || n >= upperBound {
			t.Errorf("SecIntN(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestSecInt32N(t *testing.T) {
	upperBound := int32(100)
	for i := 0; i < 1000; i++ {
		if n := SecInt32N(upperBound); n < 0 || n >= upperBound {
			t.Errorf("SecInt32N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestSecInt64N(t *testing.T) {
	upperBound := int64(100)
	for i := 0; i < 1000; i++ {
		if n := SecInt64N(upperBound); n < 0 || n >= upperBound {
			t.Errorf("SecInt64N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestSecUintN(t *testing.T) {
	upperBound := uint(100)
	for i := 0; i < 1000; i++ {
		if n := SecUintN(upperBound); n >= upperBound {
			t.Errorf("SecUintN(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestSecUint32N(t *testing.T) {
	upperBound := uint32(100)
	for i := 0; i < 1000; i++ {
		if n := SecUint32N(upperBound); n >= upperBound {
			t.Errorf("SecUint32N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestSecUint64N(t *testing.T) {
	upperBound := uint64(100)
	for i := 0; i < 1000; i++ {
		if n := SecUint64N(upperBound); n >= upperBound {
			t.Errorf("SecUint64N(%d) = %d; want value in range [0, %d)", upperBound, n, upperBound)
		}
	}
}

func TestFloat32(t *testing.T) {
	const iterations = 100000
	const epsilon = 1e-1

	smallestObserved := float32(1.0)
	largestObserved := float32(0.0)
	totalSum := float32(0.0)

	for i := 0; i < iterations; i++ {
		value := Float32()
		assert.GreaterOrEqual(t, value, float32(0.0), "Float32() should return values >= 0")
		assert.Less(t, value, float32(1.0), "Float32() should return values < 1")

		smallestObserved = float32(math.Min(float64(smallestObserved), float64(value)))
		largestObserved = float32(math.Max(float64(largestObserved), float64(value)))
		totalSum += value
	}

	averageValue := totalSum / float32(iterations)
	assert.InDelta(t, averageValue, 0.5, epsilon, "Average value should be close to 0.5")
	assert.InDelta(t, smallestObserved, 0.0, epsilon, "Smallest observed value should be close to 0")
	assert.InDelta(t, largestObserved, 1.0, epsilon, "Largest observed value should be close to 1")
}

func TestFloat64(t *testing.T) {
	const iterations = 100000
	const epsilon = 1e-1

	smallestObserved := 1.0
	largestObserved := 0.0
	totalSum := 0.0

	for i := 0; i < iterations; i++ {
		value := Float64()
		assert.GreaterOrEqual(t, value, 0.0, "Float64() should return values >= 0")
		assert.Less(t, value, 1.0, "Float64() should return values < 1")

		smallestObserved = math.Min(smallestObserved, value)
		largestObserved = math.Max(largestObserved, value)
		totalSum += value
	}

	averageValue := totalSum / float64(iterations)
	assert.InDelta(t, averageValue, 0.5, epsilon, "Average value should be close to 0.5")
	assert.InDelta(t, smallestObserved, 0.0, epsilon, "Smallest observed value should be close to 0")
	assert.InDelta(t, largestObserved, 1.0, epsilon, "Largest observed value should be close to 1")
}

func TestSecFloat32(t *testing.T) {
	const iterations = 100000
	const epsilon = 1e-1

	smallestObserved := float32(1.0)
	largestObserved := float32(0.0)
	totalSum := float32(0.0)

	for i := 0; i < iterations; i++ {
		value := SecFloat32()
		assert.GreaterOrEqual(t, value, float32(0.0), "SecFloat32() should return values >= 0")
		assert.Less(t, value, float32(1.0), "SecFloat32() should return values < 1")

		smallestObserved = float32(math.Min(float64(smallestObserved), float64(value)))
		largestObserved = float32(math.Max(float64(largestObserved), float64(value)))
		totalSum += value
	}

	averageValue := totalSum / float32(iterations)
	assert.InDelta(t, averageValue, 0.5, epsilon, "Average value should be close to 0.5")
	assert.InDelta(t, smallestObserved, 0.0, epsilon, "Smallest observed value should be close to 0")
	assert.InDelta(t, largestObserved, 1.0, epsilon, "Largest observed value should be close to 1")
}

func TestSecFloat64(t *testing.T) {
	const iterations = 100000
	const epsilon = 1e-1

	smallestObserved := 1.0
	largestObserved := 0.0
	totalSum := 0.0

	for i := 0; i < iterations; i++ {
		value := SecFloat64()
		assert.GreaterOrEqual(t, value, 0.0, "SecFloat64() should return values >= 0")
		assert.Less(t, value, 1.0, "SecFloat64() should return values < 1")

		smallestObserved = math.Min(smallestObserved, value)
		largestObserved = math.Max(largestObserved, value)
		totalSum += value
	}

	averageValue := totalSum / float64(iterations)
	assert.InDelta(t, averageValue, 0.5, epsilon, "Average value should be close to 0.5")
	assert.InDelta(t, smallestObserved, 0.0, epsilon, "Smallest observed value should be close to 0")
	assert.InDelta(t, largestObserved, 1.0, epsilon, "Largest observed value should be close to 1")
}

func TestPanicConditions(t *testing.T) {
	testCases := []struct {
		name string
		f    func()
	}{
		{"SecIntN(0)", func() { SecIntN(0) }},
		{"SecInt32N(0)", func() { SecInt32N(0) }},
		{"SecInt64N(0)", func() { SecInt64N(0) }},
		{"SecUintN(0)", func() { SecUintN(0) }},
		{"SecUint32N(0)", func() { SecUint32N(0) }},
		{"SecUint64N(0)", func() { SecUint64N(0) }},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("%s did not panic", tc.name)
				}
			}()
			tc.f()
		})
	}
}
