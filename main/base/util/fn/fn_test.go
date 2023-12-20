package fn

import (
	"go4good/main/base/util/testfn"
	"testing"
)

func TestMapSlice(t *testing.T) {
	runWith := func(in []int, want []int) {
		mapFn := func(x int) int { return x + 1 }
		got := MapSlice(in, mapFn)
		testfn.VerifyEq(t, got, want)
	}

	runWith([]int{1, 2}, []int{2, 3})
	runWith([]int{}, []int{})
	runWith(nil, []int{})
}

func TestFilterSlice(t *testing.T) {
	runWith := func(in []int, want []int) {
		filterFn := func(x int) bool { return x > 2 }
		got := FilterSlice(in, filterFn)
		testfn.VerifyEq(t, got, want)
	}

	runWith([]int{1, 2, 3}, []int{3})
	runWith([]int{1, 2}, []int{})
	runWith(nil, []int{})
}
