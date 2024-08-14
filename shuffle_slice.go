package shuffle

import (
	"math/rand/v2"
	"reflect"
)

// Slice shuffles the slice.
func Slice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	rand.Shuffle(rv.Len(), swap)
}

// Slice shuffles the slice.
func (s *Shuffler) Slice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	(*rand.Rand)(s).Shuffle(rv.Len(), swap)
}
