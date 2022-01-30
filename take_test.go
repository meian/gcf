package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
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
			name: "take 1 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 1,
			},
			want: []int{1},
		},
		{
			name: "take 2 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 2,
			},
			want: []int{1, 2},
		},
		{
			name: "take 3 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "take 4 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 4,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "take 0 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 0,
			},
			want: []int{},
		},
		{
			name: "take negative",
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
			itb := gcf.Take(tt.args.itb, tt.args.count)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.Take(itb, 2)
	testBeforeAndAfter(t, itb)

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.Take(itb, 1)
	})
}

func ExampleTake() {
	itb := gcf.FromSlice([]int{1, 2, 3, 4, 5})
	itb = gcf.Take(itb, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 2 3]
}

func TestTakeWhile(t *testing.T) {
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
			want: []int{1, 2, 3},
		},
		{
			name: "all false",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return false },
			},
			want: []int{},
		},
		{
			name: "partial true from ahead",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v <= 2 },
			},
			want: []int{1, 2},
		},
		{
			name: "partial true not from ahead",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v >= 2 },
			},
			want: []int{},
		},
		{
			name: "nil func",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: nil,
			},
			want: []int{},
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
			itb := gcf.TakeWhile(tt.args.itb, tt.args.whileFunc)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.TakeWhile(itb, func(v int) bool { return v < 2 })
	testBeforeAndAfter(t, itb)

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.TakeWhile(itb, func(v int) bool { return true })
	})
}

func ExampleTakeWhile() {
	itb := gcf.FromSlice([]int{1, 3, 5, 7, 2, 4, 6, 8, 9})
	itb = gcf.TakeWhile(itb, func(v int) bool { return v%2 > 0 })
	s := gcf.ToSlice(itb)
	fmt.Println(s)
	// Output:
	// [1 3 5 7]
}
