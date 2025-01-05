package slicex

import (
	"cmp"

	"github.com/samber/lo"
)

// ToSetOrdered converts a slice of ordered elements to a set-like map.
func ToSetOrdered[E cmp.Ordered, S ~[]E](s S) map[E]struct{} {
	return lo.SliceToMap(s, func(item E) (E, struct{}) {
		return item, struct{}{}
	})
}

// ToBoolSetOrdered converts a slice of ordered elements to a map where the keys are the slice elements
// and the values are all set to true. This effectively creates a set-like structure using a map.
func ToBoolSetOrdered[E cmp.Ordered, S ~[]E](s S) map[E]bool {
	return lo.SliceToMap(s, func(item E) (E, bool) {
		return item, true
	})
}
