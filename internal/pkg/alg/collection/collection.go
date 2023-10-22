package collection

import "slices"

// IsSubset
// if the s2 contains all the elements in s1 return true
// or return false
// time O(n) space O(n), another way is sort the s2 firstly
// then use binary search in each element in s1.
func IsSubset[E comparable](s1, s2 []E) bool {
	if len(s1) > len(s2) {
		return false
	}

	dict := make(map[E]struct{}, len(s2))
	for _, e := range s2 {
		dict[e] = struct{}{}
	}

	for _, e := range s1 {
		if _, exist := dict[e]; !exist {
			return false
		}
	}
	return true
}

// ComplementSet return the complement set of s2 relative to s1
func ComplementSet[E comparable](s1, s2 []E) []E {
	var result []E
	// must be subset
	if !IsSubset(s1, s2) {
		return result
	}

	for _, e := range s2 {
		if !slices.Contains(s1, e) {
			result = append(result, e)
		}
	}
	return result
}

// DifferenceSet return the difference set of s2 relative to s1
func DifferenceSet[E comparable](s1, s2 []E) []E {
	var result []E
	for _, e := range s2 {
		if !slices.Contains(s1, e) {
			result = append(result, e)
		}
	}
	return result
}
