package bench_test

import (
	"fmt"
	"testing"

	"github.com/meian/gcf"
)

func BenchmarkFilter_Volumes(b *testing.B) {
	sLens := []int{10}
	fLens := []int{1, 5, 10, 100, 1000}
	for _, sLen := range sLens {
		for _, fLen := range fLens {
			name := fmt.Sprintf("slice=%d/filter=%d", sLen, fLen)
			b.Run(name, func(b *testing.B) {
				s := make([]int, sLen)
				for i := range s {
					s[i] = i + 1
				}
				itb := gcf.FromSlice(s)
				for i := 0; i < fLen; i++ {
					itb = gcf.Filter(itb, func(v int) bool {
						return true
					})
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					it := itb.Iterator()
					for it.MoveNext() {
					}
				}
			})
		}
	}
}

func BenchmarkFilter_Compare(b *testing.B) {
	s := make([]int, 100)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	f13 := func(v int) bool {
		return v%13 > 0
	}
	f11 := func(v int) bool {
		return v%11 > 0
	}
	f7 := func(v int) bool {
		return v%7 > 0
	}

	b.Run("filter", func(b *testing.B) {
		s := append([]int{}, s...)
		itb := gcf.FromSlice(s)
		itb = gcf.Filter(itb, f13)
		itb = gcf.Filter(itb, f11)
		itb = gcf.Filter(itb, f7)
		b.ResetTimer()
		n := 0
		for i := 0; i < b.N; i++ {
			n = 0
			it := itb.Iterator()
			for it.MoveNext() {
				n++
			}
		}
	})

	b.Run("if-func", func(b *testing.B) {
		s := append([]int{}, s...)
		b.ResetTimer()
		n := 0
		for i := 0; i < b.N; i++ {
			n = 0
			for _, v := range s {
				if !f13(v) {
					continue
				}
				if !f11(v) {
					continue
				}
				if !f7(v) {
					continue
				}
				n++
			}
		}
	})

	b.Run("if-inline", func(b *testing.B) {
		s := append([]int{}, s...)
		b.ResetTimer()
		n := 0
		for i := 0; i < b.N; i++ {
			n = 0
			for _, v := range s {
				if v%13 == 0 {
					continue
				}
				if v%11 == 0 {
					continue
				}
				if v%7 == 0 {
					continue
				}
				n++
			}
		}
	})

	b.Run("chan", func(b *testing.B) {
		s := append([]int{}, s...)
		b.ResetTimer()
		n := 0
		for i := 0; i < b.N; i++ {
			n = 0
			c13 := make(chan int)
			c11 := make(chan int)
			c7 := make(chan int)
			go func() {
				defer close(c13)
				for _, v := range s {
					if v%13 > 0 {
						c13 <- v
					}
				}
			}()
			go func() {
				defer close(c11)
				for v := range c13 {
					if v%11 > 0 {
						c11 <- v
					}
				}
			}()
			go func() {
				defer close(c7)
				for v := range c11 {
					if v%7 > 0 {
						c7 <- v
					}
				}
			}()
			for range c7 {
				n++
			}
		}
	})
}
