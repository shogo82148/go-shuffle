package shuffle

// Interface is a type, typically a collection, that satisfies shuffle.Interface can be
// shuffled by the routines in this package.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Ints shuffles a slice of ints.
func Ints(a []int) { Slice(a) }

// Float64s shuffles a slice of float64s.
func Float64s(a []float64) { Slice(a) }

// Strings shuffles a slice of strings.
func Strings(a []string) { Slice(a) }
