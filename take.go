package gcf

type takeIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type takeIterator[T any] struct {
	it      Iterator[T]
	count   int
	i       int
	done    bool
	current T
}

// Take makes Iterable with count elements from ahead.
//
//   itb := gcf.FromSlice([]{1, 2, 3})
//   itb = gcf.Take(itb, 2)
//
// If count is 0 or negative, returns empty Iterable.
func Take[T any](itb Iterable[T], count int) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	if count < 1 {
		return empty[T]()
	}
	return &takeIterable[T]{itb, count}
}

func (itb *takeIterable[T]) Iterator() Iterator[T] {
	return &takeIterator[T]{itb.itb.Iterator(), itb.count, 0, false, zero[T]()}
}

func (it *takeIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if it.count <= it.i {
		it.done = true
		it.current = zero[T]()
		return false
	}
	it.i++
	if it.it.MoveNext() {
		it.current = it.it.Current()
		return true
	}
	it.done = true
	it.current = zero[T]()
	return false
}

func (it *takeIterator[T]) Current() T {
	return it.current
}
