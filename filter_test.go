package gcf_test

import (
	"testing"

	"github.com/meian/gcf"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int) bool
		want      []int
	}{
		{
			name:      "ture all",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v < 10 },
			want:      []int{1, 2, 3, 4, 5},
		},
		{
			name:      "false all",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v > 10 },
			want:      []int{},
		},
		{
			name:      "true or false conditional",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: func(v int) bool { return v%2 > 0 },
			want:      []int{1, 3, 5},
		},
		{
			name:      "nil predicate",
			slice:     []int{1, 2, 3, 4, 5},
			predicate: nil,
			want:      []int{1, 2, 3, 4, 5},
		},
		{
			name:      "nil slice",
			slice:     nil,
			predicate: func(v int) bool { return v > 10 },
			want:      []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			itb := gcf.Filter(gcf.FromSlice(tt.slice), tt.predicate)
			s := gcf.ToSlice(itb)
			assert.Equal(tt.want, s)
		})
	}

	t.Run("nil iterable", func(t *testing.T) {
		itb := gcf.Filter(nil, func(v int) bool { return true })
		assert.Equal(t, []int{}, gcf.ToSlice(itb))
	})
}
