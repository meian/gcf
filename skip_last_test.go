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

func TestSkipLastWhile(t *testing.T) {
	type args struct {
		itb       gcf.Iterable[int]
		whileFunc func(v int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "all true",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return true },
			},
			want: []int{},
		},
		{
			name: "all false",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return false },
			},
			want: []int{1, 2, 3},
		},
		{
			name: "partial true from end",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v >= 2 },
			},
			want: []int{1},
		},
		{
			name: "partial true not from end",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v <= 2 },
			},
			want: []int{1, 2, 3},
		},
		{
			name: "1 element with filter true",
			args: args{
				itb:       gcf.FromSlice([]int{1}),
				whileFunc: func(v int) bool { return true },
			},
			want: []int{},
		},
		{
			name: "1 element with filter false",
			args: args{
				itb:       gcf.FromSlice([]int{1}),
				whileFunc: func(v int) bool { return false },
			},
			want: []int{1},
		},
		{
			name: "nil func",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: nil,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "nil Iterable",
			args: args{
				itb:       nil,
				whileFunc: func(v int) bool { return true },
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itb := gcf.SkipLastWhile(tt.args.itb, tt.args.whileFunc)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.SkipLastWhile(itb, func(v int) bool { return v > 2 })
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.SkipLastWhile(itb, func(v int) bool { return true })
	})
}

func ExampleSkipLastWhile() {
	itb := gcf.FromSlice([]int{1, 3, 5, 7, 2, 4, 6, 8, 9})
	itb = gcf.SkipLastWhile(itb, func(v int) bool { return v > 3 })
	s := gcf.ToSlice(itb)
	fmt.Println(s)
	// Output:
	// [1 3 5 7 2]
}
