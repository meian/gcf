package gcf_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []int
		filterFunc func(int) bool
		want       []int
	}{
		{
			name:       "all true",
			slice:      []int{1, 2, 3, 4, 5},
			filterFunc: func(v int) bool { return v < 10 },
			want:       []int{1, 2, 3, 4, 5},
		},
		{
			name:       "all false",
			slice:      []int{1, 2, 3, 4, 5},
			filterFunc: func(v int) bool { return v > 10 },
			want:       []int{},
		},
		{
			name:       "partial true",
			slice:      []int{1, 2, 3, 4, 5},
			filterFunc: func(v int) bool { return v%2 > 0 },
			want:       []int{1, 3, 5},
		},
		{
			name:       "nil func",
			slice:      []int{1, 2, 3, 4, 5},
			filterFunc: nil,
			want:       []int{1, 2, 3, 4, 5},
		},
		{
			name:       "nil slice",
			slice:      nil,
			filterFunc: func(v int) bool { return v > 10 },
			want:       []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb := gcf.Filter(gcf.FromSlice(tt.slice), tt.filterFunc)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}

	t.Run("nil iterable", func(t *testing.T) {
		itb := gcf.Filter(nil, func(v int) bool { return true })
		assert.Equal(t, []int{}, gcf.ToSlice(itb))
	})

	itb := gcf.FromSlice([]int{1, 2, 3, 4})
	itb = gcf.Filter(itb, func(v int) bool { return v%2 == 0 })
	testBeforeAndAfter(t, itb)
}

func ExampleFilter() {
	s := []string{"dog", "cat", "mouse", "rabbit"}
	itb := gcf.FromSlice(s)
	itb = gcf.Filter(itb, func(v string) bool {
		return len(v) > 3
	})
	for it := itb.Iterator(); it.MoveNext(); {
		fmt.Println(it.Current())
	}
	// Output:
	// mouse
	// rabbit
}
