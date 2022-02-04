package gcf

type takeLastIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type takeLastIterator[T any] struct {
	it      Iterator[T]
	count   int
	i       int
	built   bool
	current T
}

// TakeLast makes Iterable with count elements from end.
//
//   itb := gcf.FromSlice([]{1, 2, 3})
//   itb = gcf.TakeLast(itb, 2)
//
// If count is 0 or negative, returns empty Iterable.
func TakeLast[T any](itb Iterable[T], count int) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count < 1 {
		return empty[T]()
	}
	return &takeLastIterable[T]{itb, count}
}

func (itb *takeLastIterable[T]) Iterator() Iterator[T] {
	return &takeLastIterator[T]{itb.itb.Iterator(), itb.count, 0, false, zero[T]()}
}

func (it *takeLastIterator[T]) MoveNext() bool {
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

func (it *takeLastIterator[T]) Current() T {
	return it.current
}

func (it *takeLastIterator[T]) build() {
	s := iteratorToSlice(it.it)
	if len(s) <= it.count {
		it.it = makeSliceIterator(s)
	} else {
		it.it = makeSliceIterator(s[len(s)-it.count:])
	}
	it.built = true
}
