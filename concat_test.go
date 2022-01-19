package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	itb1 := gcf.FromSlice([]int{1, 2, 3})
	itb2 := gcf.FromSlice([]int{4, 5, 6})
	itbe := gcf.FromSlice[int](nil)
	tests := []struct {
		name string
		itb1 gcf.Iterable[int]
		itb2 gcf.Iterable[int]
		want []int
	}{
		{
			name: "normal",
			itb1: itb1,
			itb2: itb2,
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "itb1 is empty",
			itb1: itbe,
			itb2: itb2,
			want: []int{4, 5, 6},
		},
		{
			name: "itb2 is empty",
			itb1: itb1,
			itb2: itbe,
			want: []int{1, 2, 3},
		},
		{
			name: "itb1 is nil",
			itb1: nil,
			itb2: itb2,
			want: []int{4, 5, 6},
		},
		{
			name: "itb2 is nil",
			itb1: itb1,
			itb2: nil,
			want: []int{1, 2, 3},
		},
		{
			name: "both empty",
			itb1: itbe,
			itb2: itbe,
			want: []int{},
		},
		{
			name: "both nil",
			itb1: nil,
			itb2: nil,
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb := gcf.Concat(tt.itb1, tt.itb2)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}
}

func ExampleConcat() {
	itb1 := gcf.FromSlice([]int{1, 2, 3})
	itb2 := gcf.FromSlice([]int{5, 6, 7})
	itb := gcf.Concat(itb1, itb2)
	fmt.Println(gcf.ToSlice(itb))
	// [ 1 2 3 5 6 7 ]
}
