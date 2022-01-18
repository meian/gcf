package gcf

type filterIterable[T any] struct {
	itb       Iterable[T]
	predicate func(v T) bool
}

type filterIterator[T any] struct {
	iterator  Iterator[T]
	predicate func(v T) bool
	current   T
}

// Filter makes Iterable contains elements which predicate is true.
//
//   itb := gcf.FromSlice([]int{1, 2, 3})
//   itb = gcf.Filter(itb, func(v int) bool { return v%2 > 0 })
//
// If predicate is nil, return original Iteratable.
func Filter[T any](itb Iterable[T], predicate func(v T) bool) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	if predicate == nil {
		return itb
	}
	return &filterIterable[T]{itb, predicate}
}

func (itb *filterIterable[T]) Iterator() Iterator[T] {
	return &filterIterator[T]{itb.itb.Iterator(), itb.predicate, zero[T]()}
}

func (it *filterIterator[T]) MoveNext() bool {
	for true {
		if !it.iterator.MoveNext() {
			it.current = zero[T]()
			return false
		}
		c := it.iterator.Current()
		if it.predicate(c) {
			it.current = c
			return true
		}
	}
	panic("unreachable")
}

func (it *filterIterator[T]) Current() T {
	return it.current
}
