//go:build go1.10 && !go1.18
// +build go1.10,!go1.18

// Package shuffle provides primitives for shuffling slices and user-defined
// collections.
package shuffle

import (
	"math/rand"
	"sort"
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
func (s *Shuffler) Shuffle(data Interface) {
	n := data.Len()
	for i := n - 1; i >= 0; i-- {
		j := (*rand.Rand)(s).Intn(i + 1)
		data.Swap(i, j)
	}
}

// Ints shuffles a slice of ints.
func (s *Shuffler) Ints(a []int) { s.Shuffle(sort.IntSlice(a)) }

// Float64s shuffles a slice of float64s.
func (s *Shuffler) Float64s(a []float64) { s.Shuffle(sort.Float64Slice(a)) }

// Strings shuffles a slice of strings.
func (s *Shuffler) Strings(a []string) { s.Shuffle(sort.StringSlice(a)) }
