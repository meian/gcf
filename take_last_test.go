package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestTakeLast(t *testing.T) {
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
			name: "take last 1 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 1,
			},
			want: []int{3},
		},
		{
			name: "take last 2 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 2,
			},
			want: []int{2, 3},
		},
		{
			name: "take last 3 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 3,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "take last 4 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 4,
			},
			want: []int{1, 2, 3},
		},
		{
			name: "take last 0 from slice 3",
			args: args{
				itb:   gcf.FromSlice([]int{1, 2, 3}),
				count: 0,
			},
			want: []int{},
		},
		{
			name: "take last negative",
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
					_ = gcf.TakeLast(tt.args.itb, tt.args.count)
				})
				return
			}
			itb := gcf.TakeLast(tt.args.itb, tt.args.count)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.TakeLast(itb, 2)
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.TakeLast(itb, 2)
	})
}

func ExampleTakeLast() {
	itb := gcf.FromSlice([]int{1, 2, 3, 4, 5})
	itb = gcf.TakeLast(itb, 3)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [3 4 5]
}

func TestTakeLastWhile(t *testing.T) {
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
			name: "partial true from end",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v >= 2 },
			},
			want: []int{2, 3},
		},
		{
			name: "partial true not from end",
			args: args{
				itb:       gcf.FromSlice([]int{1, 2, 3}),
				whileFunc: func(v int) bool { return v <= 2 },
			},
			want: []int{},
		},
		{
			name: "1 element with filter true",
			args: args{
				itb:       gcf.FromSlice([]int{1}),
				whileFunc: func(v int) bool { return true },
			},
			want: []int{1},
		},
		{
			name: "1 element with filter false",
			args: args{
				itb:       gcf.FromSlice([]int{1}),
				whileFunc: func(v int) bool { return false },
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
			itb := gcf.TakeLastWhile(tt.args.itb, tt.args.whileFunc)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.TakeLastWhile(itb, func(v int) bool { return v > 2 })
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.TakeLastWhile(itb, func(v int) bool { return true })
	})
}

func ExampleTakeLastWhile() {
	itb := gcf.FromSlice([]int{1, 3, 5, 7, 2, 4, 6, 8, 9})
	itb = gcf.TakeLastWhile(itb, func(v int) bool { return v > 3 })
	s := gcf.ToSlice(itb)
	fmt.Println(s)
	// Output:
	// [4 6 8 9]
}
