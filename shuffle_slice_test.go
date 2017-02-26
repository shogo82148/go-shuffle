//+build go1.8

package shuffle

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestShuffleSlice(t *testing.T) {
	a := make([]int, 20)
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		a[i] = i
		b[i] = i
	}

	Slice(a)
	if reflect.DeepEqual(a, b) {
		t.Errorf("%v has not been shuffled", a)
	}

	sort.Ints(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("got %v\nwant %v", a, b)
	}
}

func BenchmarkSlice(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000} {

		b.Run(fmt.Sprintf("slice %d", n), func(b *testing.B) {
			a := make([]int, n)
			for i := 0; i < b.N; i++ {
				Slice(a)
			}
		})

		b.Run(fmt.Sprintf("ints %d", n), func(b *testing.B) {
			a := make([]int, n)
			for i := 0; i < b.N; i++ {
				Ints(a)
			}
		})

	}
}
