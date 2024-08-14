package shuffle

import (
	"math/rand/v2"
)

// Interface is a type, typically a collection, that satisfies shuffle.Interface can be
// shuffled by the routines in this package.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Ints shuffles a slice of ints.
func Ints(a []int) { shuffleSlice(a) }

// Float64s shuffles a slice of float64s.
func Float64s(a []float64) { shuffleSlice(a) }

// Strings shuffles a slice of strings.
func Strings(a []string) { shuffleSlice(a) }

func shuffleSlice[T any](a []T) {
	for i := len(a) - 1; i >= 0; i-- {
		j := rand.IntN(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
