package gcf

type repeatIteratable[T any] struct {
	itb   Iteratable[T]
	count int
}

type repeatIterator[T any] struct {
	genIt   func() Iterator[T]
	it      Iterator[T]
	count   int
	i       int
	current T
}

func Repeat[T any](itb Iteratable[T], count int) Iteratable[T] {
	if itb == nil {
		itb = empty[T]()
		count = 0
	}
	if count == 1 {
		return itb
	}
	return &repeatIteratable[T]{itb, count}
}

func (itb *repeatIteratable[T]) Iterator() Iterator[T] {
	return &repeatIterator[T]{itb.itb.Iterator, itb.itb.Iterator(), itb.count, 0, zero[T]()}
}

func (it *repeatIterator[T]) MoveNext() bool {
	if it.i >= it.count {
		return false
	}
	for it.i < it.count {
		if it.it.MoveNext() {
			it.current = it.it.Current()
			return true
		}
		it.it = it.genIt()
		it.i++
	}
	it.current = zero[T]()
	return false
}

func (it *repeatIterator[T]) Current() T {
	return it.current
}
