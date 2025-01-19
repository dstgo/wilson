package timex

import (
	"time"
)

func Yesterday() time.Time {
	return Today().AddDate(0, 0, -1)
}

func Today() time.Time {
	return time.Now().Truncate(24 * time.Hour)
}

func Tomorrow() time.Time {
	return Today().AddDate(0, 0, 1)
}
