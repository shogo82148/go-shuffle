package shuffle

import (
	"reflect"
	"sort"
	"testing"
)

func TestShuffle(t *testing.T) {
	a := make([]int, 20)
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		a[i] = i
		b[i] = i
	}

	Ints(a)
	if reflect.DeepEqual(a, b) {
		t.Errorf("%v has not been shuffled", a)
	}

	sort.Ints(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("got %v\nwant %v", a, b)
	}
}

func BenchmarkInts3(b *testing.B) {
	a := make([]int, 3)
	for n := b.N; n > 0; n-- {
		Ints(a)
	}
}

func BenchmarkInts20(b *testing.B) {
	a := make([]int, 20)
	for n := b.N; n > 0; n-- {
		Ints(a)
	}
}
