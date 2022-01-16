package gcf

type mapIterable[T any, R any] struct {
	itb      Iterable[T]
	selector func(T) R
}

type mapIterator[T any, R any] struct {
	it       Iterator[T]
	selector func(T) R
	current  R
}

func Map[T any, R any](itb Iterable[T], f func(T) R) Iterable[R] {
	if itb == nil {
		itb = empty[T]()
	}
	if f == nil {
		r := zero[R]()
		f = func(v T) R { return r }
	}
	return &mapIterable[T, R]{itb, f}
}

func (itb *mapIterable[T, R]) Iterator() Iterator[R] {
	return &mapIterator[T, R]{itb.itb.Iterator(), itb.selector, zero[R]()}
}

func (it *mapIterator[T, R]) MoveNext() bool {
	if !it.it.MoveNext() {
		it.current = zero[R]()
		return false
	}
	it.current = it.selector(it.it.Current())
	return true
}

func (it *mapIterator[T, R]) Current() R {
	return it.current
}
