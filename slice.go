package gcf

type sliceIterable[T any] struct {
	slice []T
}

type sliceIterator[T any] struct {
	slice []T
	len   int
	i     int
}

// FromSlice makes Iterable from slice.
//
//   s := []int{1, 2, 3}
//   itb := gcf.FromSlice(s)
//
// By change elements in base slice afrer this called, change is affected to Iterator.
// If you want no affects by change, you can use FromSliceImmutable.
func FromSlice[T any](s []T) Iterable[T] {
	if s == nil {
		s = []T{}
	}
	return &sliceIterable[T]{s}
}

// FromSliceImmutable makes Iterable from slice with immutable.
//
//   s := []int{1, 2, 3}
//   itb := gcf.FromSliceImmutable(s)
//
// Input slice is duplicated to make immutable, so have some performance bottleneck.
func FromSliceImmutable[T any](s []T) Iterable[T] {
	ss := make([]T, 0, len(s))
	return &sliceIterable[T]{append(ss, s...)}
}

func (itb *sliceIterable[T]) Iterator() Iterator[T] {
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

// ToSlice makes slice of elements listed in Iterable.
func ToSlice[T any](itb Iterable[T]) []T {
	// shortcut for sliceIterable
	if sitb, ok := itb.(*sliceIterable[T]); ok {
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

func empty[T any]() Iterable[T] {
	return FromSliceImmutable([]T{})
}
