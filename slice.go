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
	return makeSliceIterator(itb.slice)
}

func makeSliceIterator[T any](s []T) *sliceIterator[T] {
	return &sliceIterator[T]{s, len(s), 0}
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
//
//   itb := gcf.FromSlice([]int{1, 2, 3})
//   s := gcf.ToSlice(itb)
func ToSlice[T any](itb Iterable[T]) []T {
	// shortcut for sliceIterable
	if sitb, ok := itb.(*sliceIterable[T]); ok {
		ss := make([]T, 0, len(sitb.slice))
		return append(ss, sitb.slice...)
	}
	return iteratorToSlice(itb.Iterator())
}

func iteratorToSlice[T any](it Iterator[T]) []T {
	s := make([]T, 0)
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
	return FromSlice[T](nil)
}
