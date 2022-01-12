package gcf

type sliceIteratable[T any] struct {
	slice []T
}

type sliceIterator[T any] struct {
	slice []T
	len   int
	i     int
}

// FromSlice make Iteratable from slice.
//
//   s := []int{1, 2, 3}
//   itb := gcf.FromSlice(s)
//
// By change elements in base slice afrer this called, change is affected to Iterator.
// If you want no affects by change, you can use FromSliceImmutable.
func FromSlice[T any](s []T) Iteratable[T] {
	if s == nil {
		s = []T{}
	}
	return &sliceIteratable[T]{s}
}

// FromSliceImmutable make Iteratable from slice with immutable.
//
//   s := []int{1, 2, 3}
//   itb := gcf.FromSliceImmutable(s)
//
// Input slice is duplicated to make immutable, so have some performance bottleneck.
func FromSliceImmutable[T any](s []T) Iteratable[T] {
	ss := make([]T, 0, len(s))
	return &sliceIteratable[T]{append(ss, s...)}
}

func (itb *sliceIteratable[T]) Iterator() Iterator[T] {
	return &sliceIterator[T]{itb.slice, len(itb.slice), 0}
}

func (it *sliceIterator[T]) MoveNext() bool {
	if it.i > it.len {
		return false
	}
	it.i++
	return it.i <= it.len
}

func (it *sliceIterator[T]) Current() T {
	if it.i <= 0 || it.len < it.i {
		return zero[T]()
	}
	return it.slice[it.i-1]
}

// ToSlice makes slice of elements listed in Iteratable.
func ToSlice[T any](itb Iteratable[T]) []T {
	// shortcut for sliceIteratable
	if sitb, ok := itb.(*sliceIteratable[T]); ok {
		ss := make([]T, 0, len(sitb.slice))
		return append(ss, sitb.slice...)
	}
	s := make([]T, 0)
	it := itb.Iterator()
	for it.MoveNext() {
		s = append(s, it.Current())
	}
	return s
}

func zero[T any]() T {
	var v T
	return v
}

func empty[T any]() Iteratable[T] {
	return FromSliceImmutable([]T{})
}
