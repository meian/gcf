package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestRepeatIteratable(t *testing.T) {
	itb := gcf.FromSlice([]int{1, 2, 3})
	itbe := gcf.FromSlice[int](nil)
	tests := []struct {
		name  string
		itb   gcf.Iteratable[int]
		count int
		want  []int
	}{
		{
			name:  "3 times",
			itb:   itb,
			count: 3,
			want:  []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
		},
		{
			name:  "empty iteratable",
			itb:   itbe,
			count: 3,
			want:  []int{},
		},
		{
			name:  "nil iteratable",
			itb:   nil,
			count: 3,
			want:  []int{},
		},
		{
			name:  "1 times",
			itb:   itb,
			count: 1,
			want:  []int{1, 2, 3},
		},
		{
			name:  "0 times",
			itb:   itb,
			count: 0,
			want:  []int{},
		},
		{
			name:  "negative times",
			itb:   itb,
			count: -1,
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb := gcf.RepeatIteratable(tt.itb, tt.count)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}
}

func ExampleRepeatIteratable() {
	s := []int{1, 2, 3}
	itb := gcf.FromSlice(s)
	itb = gcf.RepeatIteratable(itb, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 2 3 1 2 3 1 2 3]
}
