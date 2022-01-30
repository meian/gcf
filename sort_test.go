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

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.SortAsc(itb)
	})
}

func ExampleSortAsc() {
	itb := gcf.FromSlice([]int{3, 6, 7, 1, 5, 6, 2, 4, 5})
	itb = gcf.SortAsc(itb)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [1 2 3 4 5 5 6 6 7]
}

func TestSortDesc(t *testing.T) {
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
			want: []int{9, 7, 5, 3, 1},
		},
		{
			name: "reverse slice",
			args: args{
				itb: gcf.FromSlice([]int{9, 7, 5, 3, 1}),
			},
			want: []int{9, 7, 5, 3, 1},
		},
		{
			name: "duplicated slice",
			args: args{
				itb: gcf.FromSlice([]int{2, 4, 3, 5, 2, 4, 7, 6, 8}),
			},
			want: []int{8, 7, 6, 5, 4, 4, 3, 2, 2},
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
			itb := gcf.SortDesc(tt.args.itb)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	itb = gcf.SortDesc(itb)
	testBeforeAndAfter(t, itb)

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.SortDesc(itb)
	})
}

func ExampleSortDesc() {
	itb := gcf.FromSlice([]int{3, 6, 7, 1, 5, 6, 2, 4, 5})
	itb = gcf.SortDesc(itb)
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [7 6 6 5 5 4 3 2 1]
}

func TestSortBy(t *testing.T) {
	type data struct{ v int }
	type args struct {
		itb  gcf.Iterable[data]
		less func(x, y data) bool
	}
	tests := []struct {
		name string
		args args
		want []data
	}{
		{
			name: "sorted slice",
			args: args{
				itb:  gcf.FromSlice([]data{{1}, {3}, {5}, {7}, {9}}),
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{{1}, {3}, {5}, {7}, {9}},
		},
		{
			name: "sorted slice with reverse func",
			args: args{
				itb:  gcf.FromSlice([]data{{1}, {3}, {5}, {7}, {9}}),
				less: func(x, y data) bool { return x.v > y.v },
			},
			want: []data{{9}, {7}, {5}, {3}, {1}},
		},
		{
			name: "reverse slice",
			args: args{
				itb:  gcf.FromSlice([]data{{9}, {7}, {5}, {3}, {1}}),
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{{1}, {3}, {5}, {7}, {9}},
		},
		{
			name: "duplicated slice",
			args: args{
				itb:  gcf.FromSlice([]data{{2}, {4}, {3}, {5}, {2}, {4}, {7}, {6}, {8}}),
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{{2}, {2}, {3}, {4}, {4}, {5}, {6}, {7}, {8}},
		},
		{
			name: "same elements only",
			args: args{
				itb:  gcf.FromSlice([]data{{2}, {2}, {2}, {2}, {2}}),
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{{2}, {2}, {2}, {2}, {2}},
		},
		{
			name: "1 element",
			args: args{
				itb:  gcf.FromSlice([]data{{2}}),
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{{2}},
		},
		{
			name: "empty",
			args: args{
				itb:  gcf.FromSlice[data](nil),
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{},
		},
		{
			name: "nil",
			args: args{
				itb:  nil,
				less: func(x, y data) bool { return x.v < y.v },
			},
			want: []data{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itb := gcf.SortBy(tt.args.itb, tt.args.less)
			s := gcf.ToSlice(itb)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]data{{1}, {2}, {3}})
	itb = gcf.SortBy(itb, func(x, y data) bool { return x.v < y.v })
	testBeforeAndAfter(t, itb)

	testEmptyChain(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.SortBy(itb, func(x, y int) bool { return true })
	})
}

func ExampleSortBy() {
	type data struct{ v int }
	itbi := gcf.FromSlice([]int{3, 6, 7, 1, 5, 6, 2, 4, 5})
	itb := gcf.Map(itbi, func(v int) data { return data{v} })
	itb = gcf.SortBy(itb, func(x, y data) bool { return x.v < y.v })
	fmt.Println(gcf.ToSlice(itb))
	// Output:
	// [{1} {2} {3} {4} {5} {5} {6} {6} {7}]
}

func FuzzSort(f *testing.F) {
	type data struct{ v byte }
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
		itbd := gcf.SortDesc(itb)
		sd := gcf.ToSlice(itbd)
		for i := range sd[:len(sd)-1] {
			v0, v1 := sd[i], sd[i+1]
			assert.GreaterOrEqualf(v0, v1, "src: %v, i: %d", s, i)
		}
		itbs := gcf.Map(itb, func(v byte) data { return data{v} })
		itbs = gcf.SortBy(itbs, func(x, y data) bool { return x.v > y.v })
		ss := gcf.ToSlice(itbs)
		for i := range ss[:len(ss)-1] {
			v0, v1 := ss[i], ss[i+1]
			assert.GreaterOrEqualf(v0.v, v1.v, "src: %v, i: %d", s, i)
		}
	})
}
