package gcf

type filterIteratable[T any] struct {
	itb       Iteratable[T]
	predicate func(v T) bool
}

type filterIterator[T any] struct {
	iterator  Iterator[T]
	predicate func(v T) bool
	current   T
}

func Filter[T any](itb Iteratable[T], predicate func(v T) bool) Iteratable[T] {
	if itb == nil {
		itb = empty[T]()
	}
	if predicate == nil {
		predicate = func(v T) bool {
			return true
		}
	}
	return &filterIteratable[T]{itb, predicate}
}

func (itb *filterIteratable[T]) Iterator() Iterator[T] {
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
