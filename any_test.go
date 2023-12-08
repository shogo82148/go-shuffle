//go:build go1.18
// +build go1.18

package shuffle

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestAny(t *testing.T) {
	a := make([]int, 20)
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		a[i] = i
		b[i] = i
	}

	Any(a)
	if reflect.DeepEqual(a, b) {
		t.Errorf("%v has not been shuffled", a)
	}

	sort.Ints(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("got %v\nwant %v", a, b)
	}
}

func TestAnyRand(t *testing.T) {
	a := make([]int, 20)
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		a[i] = i
		b[i] = i
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	AnyRand(rnd, a)
	if reflect.DeepEqual(a, b) {
		t.Errorf("%v has not been shuffled", a)
	}

	sort.Ints(a)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("got %v\nwant %v", a, b)
	}
}

func BenchmarkAny(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000} {

		b.Run(fmt.Sprintf("slice %d", n), func(b *testing.B) {
			a := make([]int, n)
			for i := 0; i < b.N; i++ {
				Any(a)
			}
		})

	}
}

func BenchmarkAnyRand(b *testing.B) {
	for _, n := range []int{1, 10, 100, 1000, 10000} {

		b.Run(fmt.Sprintf("slice %d", n), func(b *testing.B) {
			a := make([]int, n)
			rnd := rand.New(rand.NewSource(42))
			for i := 0; i < b.N; i++ {
				AnyRand(rnd, a)
			}
		})

	}
}
