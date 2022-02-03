package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestLast(t *testing.T) {
	type args struct {
		itb   gcf.Iterable[int]
		count int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "last 1 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 1,
			},
			want: []int{3},
		},
		{
			name: "last 2 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 2,
			},
			want: []int{2, 3},
		},
		{
			name: "last 3 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "last 4 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 4,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "last 0 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 0,
			},
			want: []int{},
		},
		{
			name: "last negative",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: -1,
			},
			want: []int{},
		},
		{
			name: "nil Iterable",
			args: args{
				itb:   nil,
				count: 3,
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itb := gcf.Last(tt.args.itb, tt.args.count)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.Last(itb, 2)
	testBeforeAndAfter(t, itb)

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.Last(itb, 2)
	})
}

func ExampleLast() {
	itb := gcf.FromSlice([]int{1, 2, 3, 4, 5})
	itb = gcf.Last(itb, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [3 4 5]
}
