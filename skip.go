package gcf

type skipIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type skipIterator[T any] struct {
	it      Iterator[T]
	count   int
	i       int
	current T
}

// Skip makes Iterable with elements excepting counted elements from ahead.
//
//   itb := gcf.FromSlice([]{1, 2, 3})
//   itb = gcf.Skip(itb, 2)
//
// If count is 0 or negative, returns original Iterable.
func Skip[T any](itb Iterable[T], count int) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	if count < 1 {
		return itb
	}
	return &skipIterable[T]{itb, count}
}

func (itb *skipIterable[T]) Iterator() Iterator[T] {
	return &skipIterator[T]{itb.itb.Iterator(), itb.count, 0, zero[T]()}
}

func (it *skipIterator[T]) MoveNext() bool {
	for it.i < it.count && it.it.MoveNext() {
		it.i++
	}
	if it.it.MoveNext() {
		it.current = it.it.Current()
		return true
	}
	it.current = zero[T]()
	return false
}

func (it *skipIterator[T]) Current() T {
	return it.current
}
