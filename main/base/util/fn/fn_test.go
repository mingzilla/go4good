package fn

import (
	"reflect"
	"testing"
)

func TestMapSlice(t *testing.T) {
	runWith := func(in []int, want []int) {
		mapFn := func(x int) int { return x + 1 }
		got := MapSlice(in, mapFn)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("MapSlice() = %v, want %v", got, want)
		}
	}

	runWith([]int{1, 2}, []int{2, 3})
	runWith([]int{}, []int{})
	runWith(nil, []int{})
}

func TestFilterSlice(t *testing.T) {
	runWith := func(in []int, want []int) {
		filterFn := func(x int) bool { return x > 2 }
		got := FilterSlice(in, filterFn)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("FilterSlice() = %v, want %v", got, want)
		}
	}

	runWith([]int{1, 2, 3}, []int{3})
	runWith([]int{1, 2}, []int{})
	runWith(nil, []int{})
}
