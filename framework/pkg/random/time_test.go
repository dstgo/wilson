package random

import (
	"testing"
	"time"
)

func TestRand_Time(t *testing.T) {
	r := runtimeRng
	testCases := []struct {
		name  string
		since time.Time
		until time.Time
	}{
		{
			name:  "Normal range",
			since: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			until: time.Date(2023, 12, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			name:  "Same time",
			since: time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC),
			until: time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			name:  "1 second apart",
			since: time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC),
			until: time.Date(2023, 6, 1, 12, 0, 1, 0, time.UTC),
		},
		{
			name:  "Different timezone",
			since: time.Date(2023, 6, 1, 12, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
			until: time.Date(2023, 6, 2, 12, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := 0; i < 100; i++ {
				result := r.Time(tc.since, tc.until)
				if !result.Equal(tc.since) && (result.Before(tc.since) || result.After(tc.until) || result.Equal(tc.until)) {
					t.Errorf("Rand.Time() returned time %v, which is outside the range [%v, %v)", result, tc.since, tc.until)
				}
				if result.Location() != tc.since.Location() {
					t.Errorf("Rand.Time() returned time with location %v, expected %v", result.Location(), tc.since.Location())
				}
			}
		})
	}
}

func TestRand_TimePanic(t *testing.T) {
	r := runtimeRng
	since := time.Date(2023, 6, 1, 12, 0, 0, 0, time.UTC)
	until := time.Date(2023, 6, 2, 12, 0, 0, 0, time.FixedZone("UTC+8", 8*60*60))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Rand.Time() did not panic with different locations")
		}
	}()

	r.Time(since, until)
}
