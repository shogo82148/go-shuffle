package shuffle

import (
	"math/rand/v2"
)

// Slice shuffles the slice.
func Slice[T any](slice []T) {
	for i := len(slice) - 1; i >= 0; i-- {
		j := rand.IntN(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
