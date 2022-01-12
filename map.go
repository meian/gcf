package gcf

type mapIteratable[T any, R any] struct {
	itb      Iteratable[T]
	selector func(T) R
}

type mapIterator[T any, R any] struct {
	it       Iterator[T]
	selector func(T) R
	current  R
}

func Map[T any, R any](itb Iteratable[T], f func(T) R) Iteratable[R] {
	if itb == nil {
		itb = empty[T]()
	}
	if f == nil {
		r := zero[R]()
		f = func(v T) R { return r }
	}
	return &mapIteratable[T, R]{itb, f}
}

func (itb *mapIteratable[T, R]) Iterator() Iterator[R] {
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
