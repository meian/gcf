package gcf

type repeatIteratableIteratable[T any] struct {
	itb   Iteratable[T]
	count int
}

type repeatIteratableIterator[T any] struct {
	genIt   func() Iterator[T]
	it      Iterator[T]
	count   int
	i       int
	current T
}

// RepeatIteratable makes Iteratable that repeat elements in itb a count times.
//
//   s := []int{1, 2, 3}
//   itb := gcf.FromSlice(s)
//   itb = gcf.RepeatIteratable(itb, 3)
//
// If count is 0 or negative, return Iteratable with no element.
func RepeatIteratable[T any](itb Iteratable[T], count int) Iteratable[T] {
	switch {
	case itb == nil:
		return empty[T]()
	case count < 1:
		return empty[T]()
	case count == 1:
		return itb
	}
	return &repeatIteratableIteratable[T]{itb, count}
}

func (itb *repeatIteratableIteratable[T]) Iterator() Iterator[T] {
	return &repeatIteratableIterator[T]{itb.itb.Iterator, itb.itb.Iterator(), itb.count, 0, zero[T]()}
}

func (it *repeatIteratableIterator[T]) MoveNext() bool {
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

func (it *repeatIteratableIterator[T]) Current() T {
	return it.current
}
