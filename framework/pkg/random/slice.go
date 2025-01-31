package random

// Shuffle shuffles the first n elements of data using the default source.
func Shuffle(n int, swap func(i, j int)) {
	runtimeRng.Shuffle(n, swap)
}

// SecShuffle shuffles the first n elements of data using a cryptographically secure random number generator.
func SecShuffle(n int, swap func(i, j int)) {
	cryptoRng.Shuffle(n, swap)
}

// Perm returns a slice of n distinct random non-negative integers in [0, n)
func Perm(n int) []int {
	return runtimeRng.Perm(n)
}

// SecPerm returns a slice of n distinct random non-negative integers in [0, n) using a cryptographically secure random number generator.
func SecPerm(n int) []int {
	return cryptoRng.Perm(n)
}
