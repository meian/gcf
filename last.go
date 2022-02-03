package gcf

type lastIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type lastIterator[T any] struct {
	it      Iterator[T]
	count   int
	i       int
	built   bool
	current T
}

// Last makes Iterable with count elements from end.
//
//   itb := gcf.FromSlice([]{1, 2, 3})
//   itb = gcf.Last(itb, 2)
//
// If count is 0 or negative, returns empty Iterable.
func Last[T any](itb Iterable[T], count int) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count < 1 {
		return empty[T]()
	}
	return &lastIterable[T]{itb, count}
}

func (itb *lastIterable[T]) Iterator() Iterator[T] {
	return &lastIterator[T]{itb.itb.Iterator(), itb.count, 0, false, zero[T]()}
}

func (it *lastIterator[T]) MoveNext() bool {
	if !it.built {
		it.build()
	}
	if !it.it.MoveNext() {
		it.current = zero[T]()
		return false
	}
	it.current = it.it.Current()
	return true
}

func (it *lastIterator[T]) Current() T {
	return it.current
}

func (it *lastIterator[T]) build() {
	s := iteratorToSlice(it.it)
	if len(s) <= it.count {
		it.it = makeSliceIterator(s)
	} else {
		it.it = makeSliceIterator(s[len(s)-it.count:])
	}
	it.built = true
}
