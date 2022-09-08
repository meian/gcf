package gcf_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	itb := gcf.FromSlice([]string{"a", "bc", "def"})
	itbe := gcf.FromSlice[string](nil)
	mf := func(s string) int { return len(s) }
	tests := []struct {
		name    string
		itb     gcf.Iterable[string]
		mapFunc func(string) int
		want    []int
	}{
		{
			name:    "normal",
			itb:     itb,
			mapFunc: mf,
			want:    []int{1, 2, 3},
		},
		{
			name:    "empty iterable",
			itb:     itbe,
			mapFunc: mf,
			want:    []int{},
		},
		{
			name:    "nil iterable",
			itb:     nil,
			mapFunc: mf,
			want:    []int{},
		},
		{
			name:    "nil func",
			itb:     itb,
			mapFunc: nil,
			want:    []int{0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb2 := gcf.Map(tt.itb, tt.mapFunc)
			s := gcf.ToSlice(itb2)
			assert.Equal(tt.want, s)
		})
	}

	itbi := gcf.Map(itb, func(v string) int { return len(v) })
	testBeforeAndAfter(t, itbi)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.Map(itb, func(v int) int { return v })
	})
}

func ExampleMap() {
	s := []string{"dog", "cat", "mouse", "rabbit"}
	itb := gcf.FromSlice(s)
	itbs := gcf.Map(itb, func(v string) string {
		return fmt.Sprintf("%s:%d", v, len(v))
	})
	for it := itbs.Iterator(); it.MoveNext(); {
		fmt.Println(it.Current())
	}
	// Output:
	// dog:3
	// cat:3
	// mouse:5
	// rabbit:6
}

func TestFlatMap(t *testing.T) {
	mf := func(v string) []int {
		r := make([]int, 0)
		for i := 0; i < len(v); i++ {
			r = append(r, i+1)
		}
		return r
	}
	mfe := func(v string) []int {
		return make([]int, 0)
	}
	mfn := func(v string) []int {
		return nil
	}
	tests := []struct {
		name    string
		itb     gcf.Iterable[string]
		mapFunc func(v string) []int
		want    []int
	}{
		{
			name:    "normal",
			itb:     gcf.FromSlice([]string{"a", "bc", "def", "ghij"}),
			mapFunc: mf,
			want:    []int{1, 1, 2, 1, 2, 3, 1, 2, 3, 4},
		},
		{
			name:    "empty iterable",
			itb:     gcf.FromSlice([]string{}),
			mapFunc: mf,
			want:    []int{},
		},
		{
			name:    "nil iterable",
			itb:     nil,
			mapFunc: mf,
			want:    []int{},
		},
		{
			name:    "empty return func",
			itb:     gcf.FromSlice([]string{"a", "bc", "def", "ghij"}),
			mapFunc: mfe,
			want:    []int{},
		},
		{
			name:    "nil return func",
			itb:     gcf.FromSlice([]string{"a", "bc", "def", "ghij"}),
			mapFunc: mfn,
			want:    []int{},
		},
		{
			name:    "nil func",
			itb:     gcf.FromSlice([]string{"a", "bc", "def", "ghij"}),
			mapFunc: nil,
			want:    []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itbi := gcf.FlatMap(tt.itb, tt.mapFunc)
			s := gcf.ToSlice(itbi)
			assert.Equal(t, tt.want, s)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3, 4})
	itbs := gcf.FlatMap(itb, func(v int) []string {
		return []string{strings.Repeat("a", v)}
	})
	testBeforeAndAfter(t, itbs)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.FlatMap(itb, func(v int) []int { return []int{v, v, v} })
	})
}

func ExampleFlatMap() {
	s := []string{"dog", "cat", "mouse", "rabbit"}
	itb := gcf.FromSlice(s)
	itbs := gcf.FlatMap(itb, func(v string) []string {
		r := make([]string, 0, len(v))
		for _, c := range v {
			r = append(r, fmt.Sprintf("%c", c))
		}
		return r
	})
	fmt.Println(gcf.ToSlice(itbs))
	// Output:
	// [d o g c a t m o u s e r a b b i t]
}
