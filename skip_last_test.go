package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestSkipLast(t *testing.T) {
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
			name: "skip last 1 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 1,
			},
			want: []int{1, 2},
		},
		{
			name: "skip last 2 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 2,
			},
			want: []int{1},
		},
		{
			name: "skip last 3 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 3,
			},
			want: []int{},
		},
		{
			name: "skip last 4 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 4,
			},
			want: []int{},
		},
		{
			name: "skip last 0 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 0,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "skip last negative",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: -1,
			},
			want: []int{1, 2, 3},
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
			itb := gcf.SkipLast(tt.args.itb, tt.args.count)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.SkipLast(itb, 2)
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.SkipLast(itb, 2)
	})
}

func ExampleSkipLast() {
	itb := gcf.FromSlice([]int{1, 2, 3, 4, 5})
	itb = gcf.SkipLast(itb, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 2]
}
