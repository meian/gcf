package gcf_test

import (
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestFromSlice_int(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
	}{
		{
			name:  "normal slice",
			slice: []int{1, 2, 3},
		},
		{
			name:  "ampty slice",
			slice: []int{},
		},
		{
			name:  "nil",
			slice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			// test for iterated values
			itb := gcf.FromSlice(tt.slice)
			it := itb.Iterator()
			assert.Zero(it.Current())
			for i, v := range tt.slice {
				assert.True(it.MoveNext(), "i=%d", i)
				assert.Equal(v, it.Current(), "i=%d", i)
			}

			// test for move count
			it = itb.Iterator()
			itCnt := 0
			for ; it.MoveNext(); itCnt++ {
			}
			assert.Len(tt.slice, itCnt)
		})
	}

	itb := gcf.FromSlice([]int{1, 2, 3})
	testBeforeAndAfter(t, itb)
}

func TestFromSlice_pointer(t *testing.T) {
	i1, i2, i3 := 1, 2, 3
	tests := []struct {
		name  string
		slice []*int
	}{
		{
			name:  "normal slice",
			slice: []*int{&i1, &i2, &i3},
		},
		{
			name:  "ampty slice",
			slice: []*int{},
		},
		{
			name:  "nil",
			slice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			// test for iterated values
			itb := gcf.FromSlice(tt.slice)
			it := itb.Iterator()
			assert.Zero(it.Current())
			for i, v := range tt.slice {
				assert.True(it.MoveNext(), "i=%d", i)
				assert.Equal(v, it.Current(), "i=%d", i)
			}

			// test for move count
			it = itb.Iterator()
			itCnt := 0
			for ; it.MoveNext(); itCnt++ {
			}
			assert.Len(tt.slice, itCnt)
		})
	}

	itb := gcf.FromSlice([]*int{&i1, &i2, &i3})
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.FromSlice([]int{})
	})
}

func TestFromSliceImmutable(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	tests := []struct {
		name    string
		genIter func([]string) gcf.Iterable[string]
		want    string
	}{
		{
			name:    "FromSlice",
			genIter: gcf.FromSlice[string],
			want:    "fbcde",
		},
		{
			name:    "FromSliceImmutable",
			genIter: gcf.FromSliceImmutable[string],
			want:    "abcde",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			ss := append([]string{}, s...)
			itb := tt.genIter(ss)
			ss[0] = "f"
			actual := ""
			it := itb.Iterator()
			for it.MoveNext() {
				actual += it.Current()
			}
			assert.Equal(tt.want, actual)
		})
	}

	t.Run("nil slice", func(t *testing.T) {
		itb := gcf.FromSliceImmutable[string](nil)
		actual := ""
		it := itb.Iterator()
		for it.MoveNext() {
			actual += it.Current()
		}
		assert.Empty(t, actual)
	})

	itb := gcf.FromSliceImmutable([]int{1, 2, 3})
	testBeforeAndAfter(t, itb)

	testEmpties(t, func(itb gcf.Iterable[int]) gcf.Iterable[int] {
		return gcf.FromSliceImmutable([]int{})
	})
}

func TestToSlice(t *testing.T) {
	// test for coverage "shortcut for sliceIterable"
	s := []uint16{2, 3, 4, 5, 6}
	itb := gcf.FromSlice(s)
	ss := gcf.ToSlice(itb)
	assert.Equal(t, s, ss)
}
