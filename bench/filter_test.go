package gcf_test

import (
	"testing"

	"github.com/meian/gcf"
)

func BenchmarkFilter_100_1(b *testing.B) {
	// usually case
	benchFilter(b, 100, 1, func(i int) func(int) bool {
		return func(v int) bool {
			return v/100 > i
		}
	})
}

func BenchmarkFilter_1000_10(b *testing.B) {
	// usually large case
	benchFilter(b, 1000, 10, func(i int) func(int) bool {
		return func(v int) bool {
			return v/100 > i
		}
	})
}

func BenchmarkFilter_100_100(b *testing.B) {
	benchFilter(b, 100, 100, func(i int) func(int) bool {
		return func(v int) bool {
			return v > i
		}
	})
}

func BenchmarkFilter_200_100(b *testing.B) {
	benchFilter(b, 200, 100, func(i int) func(int) bool {
		return func(v int) bool {
			return v/2 > i
		}
	})
}

func BenchmarkFilter_1000_100(b *testing.B) {
	benchFilter(b, 1000, 100, func(i int) func(int) bool {
		return func(v int) bool {
			return v/10 > i
		}
	})
}

func BenchmarkFilter_100_200(b *testing.B) {
	benchFilter(b, 100, 200, func(i int) func(int) bool {
		return func(v int) bool {
			return v*2 > i
		}
	})
}

func BenchmarkFilter_100_1000(b *testing.B) {
	benchFilter(b, 100, 1000, func(i int) func(int) bool {
		return func(v int) bool {
			return v*10 > i
		}
	})
}

func benchFilter(b *testing.B, sLen, itCnt int, f func(int) func(int) bool) {
	s := make([]int, sLen)
	for i := range s {
		s[i] = i + 1
	}
	itb := gcf.FromSlice(s)
	for i := 0; i < itCnt; i++ {
		i := i
		itb = gcf.Filter(itb, f(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		it := itb.Iterator()
		for it.MoveNext() {
			sum += it.Current()
		}
	}
}

func BenchmarkOnLoop_100_1(b *testing.B) {
	benchOnLoop(b, 100, 1, func(v int) func(int) bool {
		return func(i int) bool {
			return v/100 > i
		}
	})
}

func BenchmarkOnLoop_1000_10(b *testing.B) {
	benchOnLoop(b, 1000, 10, func(v int) func(int) bool {
		return func(i int) bool {
			return v/100 > i
		}
	})
}

func BenchmarkOnLoop_100_100(b *testing.B) {
	benchOnLoop(b, 100, 100, func(v int) func(int) bool {
		return func(i int) bool {
			return v > i
		}
	})
}

func BenchmarkOnLoop_200_100(b *testing.B) {
	benchOnLoop(b, 200, 100, func(v int) func(int) bool {
		return func(i int) bool {
			return v/2 > i
		}
	})
}

func BenchmarkOnLoop_1000_100(b *testing.B) {
	benchOnLoop(b, 1000, 100, func(v int) func(int) bool {
		return func(i int) bool {
			return v/10 > i
		}
	})
}

func BenchmarkOnLoop_100_200(b *testing.B) {
	benchOnLoop(b, 100, 200, func(v int) func(int) bool {
		return func(i int) bool {
			return v*2 > i
		}
	})
}

func BenchmarkOnLoop_100_1000(b *testing.B) {
	benchOnLoop(b, 100, 1000, func(v int) func(int) bool {
		return func(i int) bool {
			return v*10 > i
		}
	})
}

func benchOnLoop(b *testing.B, sLen, lpCnt int, f func(int) func(int) bool) {
	s := make([]int, sLen)
	for i := range s {
		s[i] = i + 1
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum := 0
		for _, v := range s {
			if filterOnLoop(f(v), lpCnt) {
				sum += v
			}
		}
	}
}

func filterOnLoop(f func(int) bool, n int) bool {
	for i := 0; i < n; i++ {
		if !f(i) {
			return false
		}
	}
	return true
}
