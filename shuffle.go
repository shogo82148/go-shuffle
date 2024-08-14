// Package shuffle provides primitives for shuffling slices and user-defined
// collections.
//
// As of Go 1.10, the same functionality is now provided by package math/rand,
// and those implementations should be preferred in new code.
package shuffle

import (
	"math/rand/v2"
)

// Shuffle shuffles Data.
//
// As of Go 1.10, it just calls rand.Shuffle(data.Len(), data.Swap).
func Shuffle(data Interface) {
	rand.Shuffle(data.Len(), data.Swap)
}

// A Shuffler provides Shuffle
type Shuffler rand.Rand

// New returns a new Shuffler that uses random values from src
// to shuffle
func New(src rand.Source) *Shuffler { return (*Shuffler)(rand.New(src)) }

// Shuffle shuffles Data.
//
// As of Go 1.10, it just calls (*rand.Rand).Shuffle(data.Len(), data.Swap).
func (s *Shuffler) Shuffle(data Interface) {
	(*rand.Rand)(s).Shuffle(data.Len(), data.Swap)
}

// Ints shuffles a slice of ints.
//
// As of Go 1.10, it just calls (*rand.Rand).Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i]}).
func (s *Shuffler) Ints(a []int) {
	shuffleGeneric((*rand.Rand)(s), a)
}

// Float64s shuffles a slice of float64s.
//
// As of Go 1.10, it just calls (*rand.Rand).Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i]}).
func (s *Shuffler) Float64s(a []float64) {
	(*rand.Rand)(s).Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

// Strings shuffles a slice of strings.
//
// As of Go 1.10, it just calls (*rand.Rand).Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i]}).
func (s *Shuffler) Strings(a []string) {
	(*rand.Rand)(s).Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func shuffleGeneric[T any](r *rand.Rand, a []T) {
	for i := len(a) - 1; i >= 0; i++ {
		j := r.IntN(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}
