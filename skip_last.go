package gcf

type skipLastIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type skipLastIterator[T any] struct {
	it    Iterator[T]
	count int
	built bool
	iteratorItem[T]
}

// SkipLast makes Iterable with elements excepting counted elements from end.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.SkipLast(itb, 2)
//
// If count is 0, returns original Iterable.
// If count is negative, raises panic.
func SkipLast[T any](itb Iterable[T], count int) Iterable[T] {
	if count < 0 {
		panic("count for SkipLast must not be negative.")
	}
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count == 0 {
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

type skipLastWhileIterable[T any] struct {
	itb       Iterable[T]
	whileFunc func(v T) bool
}

type skipLastWhileIterator[T any] struct {
	it        Iterator[T]
	whileFunc func(v T) bool
	built     bool
	iteratorItem[T]
}

// SkipLastWhile makes Iterable with elements excepting elements that whileFunc is true from end.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.SkipLastWhile(itb, func(v int) bool { return v <= 2 })
//
// If whileFunc is nil, returns original Iterable.
func SkipLastWhile[T any](itb Iterable[T], whileFunc func(v T) bool) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if whileFunc == nil {
		return itb
	}
	return &skipLastWhileIterable[T]{
		itb:       itb,
		whileFunc: whileFunc,
	}
}

func (itb *skipLastWhileIterable[T]) Iterator() Iterator[T] {
	return &skipLastWhileIterator[T]{
		it:        itb.itb.Iterator(),
		whileFunc: itb.whileFunc,
	}
}

func (it *skipLastWhileIterator[T]) MoveNext() bool {
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

func (it *skipLastWhileIterator[T]) Current() T {
	return it.current
}

func (it *skipLastWhileIterator[T]) build() {
	slice := iteratorToSlice(it.it)
	if len(slice) == 0 {
		it.it = emptyIter[T]()
		it.built = true
		return
	}
	if !it.whileFunc(slice[len(slice)-1]) {
		it.it = makeSliceIterator(slice)
		it.built = true
		return
	}
	sLen := len(slice)
	for i := sLen - 2; i >= 0; i-- {
		if !it.whileFunc(slice[i]) {
			slice = slice[:i+1]
			break
		}
	}
	if len(slice) == sLen {
		it.it = emptyIter[T]()
	} else {
		it.it = makeSliceIterator(slice)
	}
	it.built = true
}
