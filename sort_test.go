package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestSortAsc(t *testing.T) {
	type args struct {
		itb gcf.Iterable[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sorted slice",
			args: args{
				itb: gcf.FromSlice([]int{1, 3, 5, 7, 9}),
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "reverse slice",
			args: args{
				itb: gcf.FromSlice([]int{9, 7, 5, 3, 1}),
			},
			want: []int{1, 3, 5, 7, 9},
		},
		{
			name: "duplicated slice",
			args: args{
				itb: gcf.FromSlice([]int{2, 4, 3, 5, 2, 4, 7, 6, 8}),
			},
			want: []int{2, 2, 3, 4, 4, 5, 6, 7, 8},
		},
		{
			name: "same elements only",
			args: args{
				itb: gcf.FromSlice([]int{2, 2, 2, 2, 2}),
			},
			want: []int{2, 2, 2, 2, 2},
		},
		{
			name: "1 element",
			args: args{
				itb: gcf.FromSlice([]int{2}),
			},
			want: []int{2},
		},
		{
			name: "empty",
			args: args{
				itb: gcf.FromSlice[int](nil),
			},
			want: []int{},
		},
		{
			name: "nil",
			args: args{
				itb: nil,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itb := gcf.SortAsc(tt.args.itb)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.SortAsc(itb)
	testBeforeAndAfter(t, itb)
}

func FuzzSortAsc(f *testing.F) {
	tests := [][]byte{
		{1, 2, 3},
		{3, 2, 1},
		{1, 3, 2},
		{1, 1, 1},
	}
	for _, tt := range tests {
		f.Add(tt)
	}
	f.Fuzz(func(t *testing.T, s []byte) {
		assert := assert.New(t)
		if len(s) < 2 {
			return
		}
		itb := gcf.FromSlice(s)
		itba := gcf.SortAsc(itb)
		sa := gcf.ToSlice(itba)
		for i := range sa[:len(sa)-1] {
			v0, v1 := sa[i], sa[i+1]
			assert.LessOrEqualf(v0, v1, "src: %v, i: %d", s, i)
		}
	})
}

func ExampleSortAsc() {
	itb := gcf.FromSlice([]int{3, 6, 7, 1, 5, 6, 2, 4, 5})
	itb = gcf.SortAsc(itb)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 2 3 4 5 5 6 6 7]
}
