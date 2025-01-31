package random

import (
	"time"
)

// Time returns a random time within the given range [since, until).
func (r *Rand) Time(since, until time.Time) time.Time {
	if since.Location().String() != until.Location().String() {
		panic("Time ranges must have the same location")
	}

	sinceTs := since.Unix()
	untilTs := until.Unix()

	return time.Unix(r.Int64Range(sinceTs, untilTs), 0).In(since.Location())
}

// Time returns a random time within the given range [since, until) using the default source.
func Time(since, until time.Time) time.Time {
	return runtimeRng.Time(since, until)
}

// SecTime returns a random time within the given range [since, until) using the secure source.
func SecTime(since, until time.Time) time.Time {
	return cryptoRng.Time(since, until)
}
