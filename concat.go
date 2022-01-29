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

// Concat makes Iterable elements concatenated of itb1 and itb2.
//
//   itb1 := gcf.FromSlice([]int{1, 2, 3})
//   itb2 := gcf.FromSlice([]int{4, 5, 6})
//   itbc := gcf.Concat(itb1, itb2)
func Concat[T any](itb1 Iterable[T], itb2 Iterable[T]) Iterable[T] {
	if isEmpty(itb1) && isEmpty(itb2) {
		return empty[T]()
	}
	if isEmpty(itb2) {
		return itb1
	}
	if isEmpty(itb1) {
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
