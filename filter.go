package gcf

type filterIterable[T any] struct {
	itb        Iterable[T]
	filterFunc func(v T) bool
}

type filterIterator[T any] struct {
	iterator   Iterator[T]
	filterFunc func(v T) bool
	current    T
}

// Filter makes Iterable with elements which filterFunc is true.
//
//   itb := gcf.FromSlice([]int{1, 2, 3})
//   itb = gcf.Filter(itb, func(v int) bool { return v%2 > 0 })
//
// If filterFunc is nil, returns original Iteratable.
func Filter[T any](itb Iterable[T], filterFunc func(v T) bool) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	if filterFunc == nil {
		return itb
	}
	return &filterIterable[T]{itb, filterFunc}
}

func (itb *filterIterable[T]) Iterator() Iterator[T] {
	return &filterIterator[T]{itb.itb.Iterator(), itb.filterFunc, zero[T]()}
}

func (it *filterIterator[T]) MoveNext() bool {
	for it.iterator.MoveNext() {
		c := it.iterator.Current()
		if it.filterFunc(c) {
			it.current = c
			return true
		}
	}
	it.current = zero[T]()
	return false
}

func (it *filterIterator[T]) Current() T {
	return it.current
}
