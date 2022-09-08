package gcf

type takeIterable[T any] struct {
	itb   Iterable[T]
	count int
}

type takeIterator[T any] struct {
	it    Iterator[T]
	count int
	i     int
	iteratorItem[T]
}

// Take makes Iterable with count elements from ahead.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.Take(itb, 2)
//
// If count is 0, returns empty Iterable.
// If count is negative, raises panic.
func Take[T any](itb Iterable[T], count int) Iterable[T] {
	if count < 0 {
		panic("count for Take must not be negative.")
	}
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if count == 0 {
		return empty[T]()
	}
	return &takeIterable[T]{itb, count}
}

func (itb *takeIterable[T]) Iterator() Iterator[T] {
	return &takeIterator[T]{
		it:    itb.itb.Iterator(),
		count: itb.count,
	}
}

func (it *takeIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if it.count <= it.i {
		it.MarkDone()
		return false
	}
	if it.it.MoveNext() {
		it.i++
		it.current = it.it.Current()
		return true
	}
	it.MarkDone()
	return false
}

func (it *takeIterator[T]) Current() T {
	return it.current
}

type takeWhileIterable[T any] struct {
	itb       Iterable[T]
	whileFunc func(v T) bool
}

type takeWhileIterator[T any] struct {
	it        Iterator[T]
	whileFunc func(v T) bool
	iteratorItem[T]
}

// TakeWhile makes Iterable with elements in which whileFunc is true from ahead.
//
//	itb := gcf.FromSlice([]{1, 2, 3})
//	itb = gcf.TakeWhile(itb, func(v int) bool { return v <= 2 })
//
// If whileFunc is nil, returns empty Iterable.
func TakeWhile[T any](itb Iterable[T], whileFunc func(v T) bool) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	if whileFunc == nil {
		return empty[T]()
	}
	return &takeWhileIterable[T]{itb, whileFunc}
}

func (itb *takeWhileIterable[T]) Iterator() Iterator[T] {
	return &takeWhileIterator[T]{
		it:        itb.itb.Iterator(),
		whileFunc: itb.whileFunc,
	}
}

func (it *takeWhileIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if !it.it.MoveNext() || !it.whileFunc(it.it.Current()) {
		it.MarkDone()
		return false
	}
	it.current = it.it.Current()
	return true
}

func (it *takeWhileIterator[T]) Current() T {
	return it.current
}
