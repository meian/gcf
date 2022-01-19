package gcf_test

import (
	"fmt"
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
