package timex

import (
	"time"
)

func UnixSeconds() uint64 {
	return uint64(time.Now().Unix())
}

func UnixMilliseconds() uint64 {
	return uint64(time.Now().UnixNano()) / 1e6
}

func UnixMicroseconds() uint64 {
	return uint64(time.Now().UnixNano()) / 1e3
}

func UnixNanoseconds() uint64 {
	return uint64(time.Now().UnixNano())
}
