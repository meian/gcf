package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestDistinct(t *testing.T) {
	tests := []struct {
		name string
		itb  gcf.Iterable[int]
		want []int
	}{
		{
			name: "uniques",
			itb:  gcf.FromSlice([]int{1, 4, 3, 2, 5}),
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "duplicates",
			itb:  gcf.FromSlice([]int{1, 2, 3, 2, 5}),
			want: []int{1, 2, 3, 5},
		},
		{
			name: "single",
			itb:  gcf.FromSlice([]int{1}),
			want: []int{1},
		},
		{
			name: "blank",
			itb:  gcf.FromSlice([]int{}),
			want: []int{},
		},
		{
			name: "nil",
			itb:  nil,
			want: []int{},
		},
		{
			name: "Distinct",
			itb:  gcf.Distinct(gcf.FromSlice([]int{1, 2, 3, 2})),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itb := gcf.Distinct(tt.itb)
			s := gcf.ToSlice(itb)
			assert.ElementsMatch(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3, 2, 4})
	itb = gcf.Distinct(itb)
	testBeforeAndAfter(t, itb)

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.Distinct(itb)
	})
}

func ExampleDistinct() {
	itb := gcf.FromSlice([]int{1, 4, 2, 3, 2, 3, 1, 2})
	itb = gcf.Distinct(itb)
	for it := itb.Iterator(); it.MoveNext(); {
		fmt.Println(it.Current())
	}
	// Unordered output:
	// 1
	// 2
	// 3
	// 4
}
