package gcf

type repeatIterable[T any] struct {
	v     T
	count int
}

type repeatIterator[T any] struct {
	v       T
	count   int
	i       int
	current T
}

type repeatIterableIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type repeatIterableIterator[T any] struct {
	genIt   func() Iterator[T]
	it      Iterator[T]
	count   int
	i       int
	current T
}

// Repeat makes Iterable that repeat v a count times.
//
//   itb = gcf.Repeat(1, 3)
//
// If count is 0 or negative, return Iterable with no element.
func Repeat[T any](v T, count int) Iterable[T] {
	switch {
	case count < 1:
		return empty[T]()
	case count == 1:
		return FromSlice([]T{v})
	}
	return &repeatIterable[T]{v, count}
}

func (itb *repeatIterable[T]) Iterator() Iterator[T] {
	return &repeatIterator[T]{itb.v, itb.count, 0, zero[T]()}
}

func (it *repeatIterator[T]) MoveNext() bool {
	if it.i >= it.count {
		it.current = zero[T]()
		return false
	}
	it.current = it.v
	it.i++
	return true
}

func (it *repeatIterator[T]) Current() T {
	return it.current
}

// RepeatIterable makes Iterable that repeat elements in itb a count times.
//
//   s := []int{1, 2, 3}
//   itb := gcf.FromSlice(s)
//   itb = gcf.RepeatIterable(itb, 3)
//
// If count is 0 or negative, return Iterable with no element.
func RepeatIterable[T any](itb Iterable[T], count int) Iterable[T] {
	switch {
	case itb == nil:
		return empty[T]()
	case count < 1:
		return empty[T]()
	case count == 1:
		return itb
	}
	return &repeatIterableIterable[T]{itb, count}
}

func (itb *repeatIterableIterable[T]) Iterator() Iterator[T] {
	return &repeatIterableIterator[T]{itb.itb.Iterator, itb.itb.Iterator(), itb.count, 0, zero[T]()}
}

func (it *repeatIterableIterator[T]) MoveNext() bool {
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

func (it *repeatIterableIterator[T]) Current() T {
	return it.current
}
