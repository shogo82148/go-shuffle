package shuffle

import "math/rand"

// Any shuffles a slice of any type.
func Any[S ~[]E, E any](x S) {
	i := len(x) - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(rand.Int63n(int64(i + 1)))
		x[i], x[j] = x[j], x[i]
	}
	for ; i > 0; i-- {
		j := int(rand.Int31n(int32(i + 1)))
		x[i], x[j] = x[j], x[i]
	}
}

// AnyRand shuffles a slice of any type using a given random source.
func AnyRand[S ~[]E, E any](r *rand.Rand, x S) {
	i := len(x) - 1
	for ; i > 1<<31-1-1; i-- {
		j := int(r.Int63n(int64(i + 1)))
		x[i], x[j] = x[j], x[i]
	}
	for ; i > 0; i-- {
		j := int(r.Int31n(int32(i + 1)))
		x[i], x[j] = x[j], x[i]
	}
}
