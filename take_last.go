package gcf

type takeLastIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type takeLastIterator[T any] struct {
	it    Iterator[T]
	count int
	built bool
	iteratorItem[T]
}

// TakeLast makes Iterable with count elements from end.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.TakeLast(itb, 2)
//
// If count is 0, returns empty Iterable.
// If count is negative, raises panic.
func TakeLast[T any](itb Iterable[T], count int) Iterable[T] {
	if count < 0 {
		panic("count for TakeLast must not be negative.")
	}
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count == 0 {
		return empty[T]()
	}
	return &takeLastIterable[T]{itb, count}
}

func (itb *takeLastIterable[T]) Iterator() Iterator[T] {
	return &takeLastIterator[T]{
		it:    itb.itb.Iterator(),
		count: itb.count,
	}
}

func (it *takeLastIterator[T]) MoveNext() bool {
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

type takeLastWhileIterable[T any] struct {
	itb       Iterable[T]
	whileFunc func(v T) bool
}

type takeLastWhileIterator[T any] struct {
	it        Iterator[T]
	whileFunc func(v T) bool
	built     bool
	iteratorItem[T]
}

// TakeLastWhile makes Iterable with elements in which whileFunc is true from end.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.TakeLastWhile(itb, func(v int) bool { return v >= 2 })
//
// If whileFunc is nil, returns empty Iterable.
func TakeLastWhile[T any](itb Iterable[T], whileFunc func(v T) bool) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if whileFunc == nil {
		return empty[T]()
	}
	return &takeLastWhileIterable[T]{itb, whileFunc}
}

func (itb *takeLastWhileIterable[T]) Iterator() Iterator[T] {
	return &takeLastWhileIterator[T]{
		it:        itb.itb.Iterator(),
		whileFunc: itb.whileFunc,
	}
}

func (it *takeLastWhileIterator[T]) MoveNext() bool {
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

func (it *takeLastWhileIterator[T]) Current() T {
	return it.current
}

func (it *takeLastWhileIterator[T]) build() {
	slice := iteratorToSlice(it.it)
	if len(slice) == 0 {
		it.it = emptyIter[T]()
		it.built = true
		return
	}
	if !it.whileFunc(slice[len(slice)-1]) {
		it.it = emptyIter[T]()
		it.built = true
		return
	}
	for i := len(slice) - 2; i >= 0; i-- {
		if !it.whileFunc(slice[i]) {
			slice = slice[i+1:]
			break
		}
	}
	it.it = makeSliceIterator(slice)
	it.built = true
}
