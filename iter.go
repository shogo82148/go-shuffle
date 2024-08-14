//go:build go1.23

package shuffle

import (
	"iter"
	"slices"
)

// Iter shuffles seq.
func Iter[T any](seq iter.Seq[T]) []T {
	s := slices.Collect(seq)
	Slice(s)
	return s
}
