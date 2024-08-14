//go:build go1.23

package shuffle_test

import (
	"fmt"

	"github.com/shogo82148/go-shuffle/v2"
)

func ExampleIter() {
	seq := func(yield func(int) bool) {
		for i := range 5 {
			if !yield(i) {
				break
			}
		}
	}

	shuffled := shuffle.Iter(seq)
	for _, value := range shuffled {
		fmt.Println(value)
	}
	// Unordered output:
	// 0
	// 1
	// 2
	// 3
	// 4
}
