package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		name  string
		v     int
		count int
		want  []int
	}{
		{
			name:  "3 times",
			v:     1,
			count: 3,
			want:  []int{1, 1, 1},
		},
		{
			name:  "1 times",
			v:     1,
			count: 1,
			want:  []int{1},
		},
		{
			name:  "0 times",
			v:     1,
			count: 0,
			want:  []int{},
		},
		{
			name:  "negative times",
			v:     1,
			count: -1,
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb := gcf.Repeat(tt.v, tt.count)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}
}

func ExampleRepeat() {
	itb := gcf.Repeat(1, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 1 1]
}

func TestRepeatIterable(t *testing.T) {
	itb := gcf.FromSlice([]int{1, 2, 3})
	itbe := gcf.FromSlice[int](nil)
	tests := []struct {
		name  string
		itb   gcf.Iterable[int]
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
			name:  "empty iterable",
			itb:   itbe,
			count: 3,
			want:  []int{},
		},
		{
			name:  "nil iterable",
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
			itb := gcf.RepeatIterable(tt.itb, tt.count)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}
}

func ExampleRepeatIterable() {
	s := []int{1, 2, 3}
	itb := gcf.FromSlice(s)
	itb = gcf.RepeatIterable(itb, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 2 3 1 2 3 1 2 3]
}
