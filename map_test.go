package gcf_test

import (
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
		itb     gcf.Iteratable[string]
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
			name:    "empty iteratable",
			itb:     itbe,
			mapFunc: mf,
			want:    []int{},
		},
		{
			name:    "nil iteratable",
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
}
