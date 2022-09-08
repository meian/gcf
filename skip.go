package gcf

type skipIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type skipIterator[T any] struct {
	it    Iterator[T]
	count int
	i     int
	iteratorItem[T]
}

// Skip makes Iterable with elements excepting counted elements from ahead.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.Skip(itb, 2)
//
// If count is 0, returns original Iterable.
// If count is negative, raises panic.
func Skip[T any](itb Iterable[T], count int) Iterable[T] {
	if count < 0 {
		panic("count for Skip must not be negative.")
	}
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count == 0 {
		return itb
	}
	return &skipIterable[T]{itb, count}
}

func (itb *skipIterable[T]) Iterator() Iterator[T] {
	return &skipIterator[T]{
		it:    itb.itb.Iterator(),
		count: itb.count,
	}
}

func (it *skipIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	for it.i < it.count && it.it.MoveNext() {
		it.i++
	}
	if !it.it.MoveNext() {
		it.MarkDone()
		return false
	}
	it.current = it.it.Current()
	return true
}

func (it *skipIterator[T]) Current() T {
	return it.current
}

type skipWhileIterable[T any] struct {
	itb       Iterable[T]
	whileFunc func(v T) bool
}

type skipWhileIterator[T any] struct {
	it        Iterator[T]
	whileFunc func(v T) bool
	skipDone  bool
	iteratorItem[T]
}

// SkipWhile makes Iterable with elements excepting elements that whileFunc is true from ahead.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.SkipWhile(itb, func(v int) bool { return v <= 2 })
//
// If whileFunc is nil, returns original Iterable.
func SkipWhile[T any](itb Iterable[T], whileFunc func(v T) bool) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if whileFunc == nil {
		return itb
	}
	return &skipWhileIterable[T]{itb, whileFunc}
}

func (itb *skipWhileIterable[T]) Iterator() Iterator[T] {
	return &skipWhileIterator[T]{
		it:        itb.itb.Iterator(),
		whileFunc: itb.whileFunc,
	}
}

func (it *skipWhileIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if !it.skipDone {
		for it.it.MoveNext() {
			if it.whileFunc(it.it.Current()) {
				continue
			}
			it.skipDone = true
			it.current = it.it.Current()
			return true
		}
	}
	if !it.it.MoveNext() {
		it.MarkDone()
		return false
	}
	it.current = it.it.Current()
	return true
}

func (it *skipWhileIterator[T]) Current() T {
	return it.current
}
