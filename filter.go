package gcf

type filterIterable[T any] struct {
	itb        Iterable[T]
	filterFunc func(v T) bool
}

type filterIterator[T any] struct {
	it         Iterator[T]
	filterFunc func(v T) bool
	iteratorItem[T]
}

// Filter makes Iterable with elements which filterFunc is true.
//
//   itb := gcf.FromSlice([]int{1, 2, 3})
//   itb = gcf.Filter(itb, func(v int) bool { return v%2 > 0 })
//
// If filterFunc is nil, returns original Iteratable.
func Filter[T any](itb Iterable[T], filterFunc func(v T) bool) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if filterFunc == nil {
		return itb
	}
	return &filterIterable[T]{itb, filterFunc}
}

func (itb *filterIterable[T]) Iterator() Iterator[T] {
	return &filterIterator[T]{
		it:         itb.itb.Iterator(),
		filterFunc: itb.filterFunc,
	}
}

func (it *filterIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	for it.it.MoveNext() {
		c := it.it.Current()
		if it.filterFunc(c) {
			it.current = c
			return true
		}
	}
	it.MarkDone()
	return false
}

func (it *filterIterator[T]) Current() T {
	return it.current
}
