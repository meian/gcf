package bench_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
)

func BenchmarkRange(b *testing.B) {
	tests := []struct {
		start int
		end   int
		step  int
	}{
		{
			start: 1,
			end:   1000000,
			step:  1,
		},
		{
			start: 1,
			end:   1000000,
			step:  10,
		},
		{
			start: 1,
			end:   1000000,
			step:  100,
		},
		{
			start: 1000000,
			end:   1,
			step:  -1,
		},
		{
			start: 1000000,
			end:   1,
			step:  -10,
		},
		{
			start: 1000000,
			end:   1,
			step:  -100,
		},
	}
	for _, tt := range tests {
		itb, _ := gcf.Range(tt.start, tt.end, tt.step)
		c := 0
		it := itb.Iterator()
		for it.MoveNext() {
			c++
		}
		name := fmt.Sprintf("start=%d/end=%d/step=%d/count=%d", tt.start, tt.end, tt.step, c)
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				it := itb.Iterator()
				for it.MoveNext() {
				}
			}
		})
	}
}
