package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestSkip(t *testing.T) {
	type args struct {
		itb   gcf.Iterable[int]
		count int
	}
	tests := []struct {
		name      string
		args      args
		want      []int
		wantPanic bool
	}{
		{
			name: "skip 1 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 1,
			},
			want: []int{2, 3},
		},
		{
			name: "skip 2 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 2,
			},
			want: []int{3},
		},
		{
			name: "skip 3 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 3,
			},
			want: []int{},
		},
		{
			name: "skip 4 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 4,
			},
			want: []int{},
		},
		{
			name: "skip 0 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 0,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "skip negative",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: -1,
			},
			wantPanic: true,
		},
		{
			name: "nil Iterable",
			args: args{
				itb:   nil,
				count: 3,
			},
			want: []int{},
		},
		{
			name: "nil and negative",
			args: args{
				itb:   nil,
				count: -1,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = gcf.Skip(tt.args.itb, tt.args.count)
				})
				return
			}
			itb := gcf.Skip(tt.args.itb, tt.args.count)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.Skip(itb, 2)
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.Skip(itb, 0)
	})
}

func ExampleSkip() {
	itb := gcf.FromSlice([]int{1, 2, 3, 4, 5})
	itb = gcf.Skip(itb, 2)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [3 4 5]
}

func TestSkipWhile(t *testing.T) {
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
			name: "partial true from ahead",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v <= 2 },
			},
			want: []int{3},
		},
		{
			name: "partial true not from ahead",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v >= 2 },
			},
			want: []int{1, 2, 3},
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
			itb := gcf.SkipWhile(tt.args.itb, tt.args.whileFunc)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.SkipWhile(itb, func(v int) bool { return v < 2 })
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.SkipWhile(itb, func(v int) bool { return true })
	})
}

func ExampleSkipWhile() {
	itb := gcf.FromSlice([]int{1, 3, 5, 7, 2, 4, 6, 8, 9})
	itb = gcf.SkipWhile(itb, func(v int) bool { return v%2 > 0 })
	s := gcf.ToSlice(itb)
	fmt.Println(s)
	// Output:
	// [2 4 6 8 9]
}
