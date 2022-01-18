package gcf

type concatIterable[T any] struct {
	itb1 Iterable[T]
	itb2 Iterable[T]
}

type concatIterator[T any] struct {
	it1     Iterator[T]
	it2     Iterator[T]
	current T
}

func Concat[T any](itb1 Iterable[T], itb2 Iterable[T]) Iterable[T] {
	switch {
	case itb1 == nil && itb2 == nil:
		return empty[T]()
	case itb2 == nil:
		return itb1
	case itb1 == nil:
		return itb2
	}
	return &concatIterable[T]{itb1, itb2}
}

func (itb *concatIterable[T]) Iterator() Iterator[T] {
	return &concatIterator[T]{itb.itb1.Iterator(), itb.itb2.Iterator(), zero[T]()}
}

func (it *concatIterator[T]) MoveNext() bool {
	if it.it1.MoveNext() {
		it.current = it.it1.Current()
		return true
	}
	if it.it2.MoveNext() {
		it.current = it.it2.Current()
		return true
	}
	it.current = zero[T]()
	return false
}

func (it *concatIterator[T]) Current() T {
	return it.current
}
