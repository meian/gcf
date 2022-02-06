package gcf

type skipLastIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type skipLastIterator[T any] struct {
	it    Iterator[T]
	count int
	i     int
	built bool
	iteratorItem[T]
}

// Skip makes Iterable with elements excepting counted elements from end.
//
//   itb := gcf.FromSlice([]{1, 2, 3})
//   itb = gcf.SkipLast(itb, 2)
//
// If count is 0 or negative, returns original Iterable.
func SkipLast[T any](itb Iterable[T], count int) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count < 1 {
		return itb
	}
	return &skipLastIterable[T]{
		itb:   itb,
		count: count,
	}
}

func (itb *skipLastIterable[T]) Iterator() Iterator[T] {
	return &skipLastIterator[T]{
		it:    itb.itb.Iterator(),
		count: itb.count,
	}
}

func (it *skipLastIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if !it.built {
		it.build()
	}
	if !it.it.MoveNext() {
		it.MarkDone()
		return false
	}
	it.current = it.it.Current()
	return true
}

func (it *skipLastIterator[T]) Current() T {
	return it.current
}

func (it *skipLastIterator[T]) build() {
	s := iteratorToSlice(it.it)
	if len(s) <= it.count {
		it.it = emptyIter[T]()
	} else {
		it.it = makeSliceIterator(s[:len(s)-it.count])
	}
	it.built = true
}
