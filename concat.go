package gcf

type concatIteratable[T any] struct {
	itb1 Iteratable[T]
	itb2 Iteratable[T]
}

type concatIterator[T any] struct {
	it1     Iterator[T]
	it2     Iterator[T]
	current T
}

func Concat[T any](itb1 Iteratable[T], itb2 Iteratable[T]) Iteratable[T] {
	if itb1 == nil {
		itb1 = empty[T]()
	}
	if itb2 == nil {
		itb2 = empty[T]()
	}
	return &concatIteratable[T]{itb1, itb2}
}

func (itb *concatIteratable[T]) Iterator() Iterator[T] {
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
