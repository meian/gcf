package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		itb  gcf.Iterable[int]
		want []int
	}{
		{
			name: "3 elements",
			itb:  gcf.FromSlice([]int{1, 2, 3}),
			want: []int{3, 2, 1},
		},
		{
			name: "1 elements",
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
			want: []int{},
		},
		{
			name: "reverse",
			itb:  gcf.Reverse(gcf.FromSlice([]int{1, 2, 3})),
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itb := gcf.Reverse(tt.itb)
			assert.Equal(t, tt.want, gcf.ToSlice(itb))
		})
	}
}

func ExampleReverse() {
	itb := gcf.FromSlice([]int{1, 3, 2, 4})
	itb = gcf.Reverse(itb)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [4 2 3 1]
}
