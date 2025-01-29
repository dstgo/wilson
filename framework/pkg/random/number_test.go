package random

import (
	"math"
	"testing"
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
