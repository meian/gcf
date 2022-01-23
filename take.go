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

type takeWhileIterable[T any] struct {
	itb       Iterable[T]
	whileFunc func(v T) bool
}

type takeWhileIterator[T any] struct {
	it        Iterator[T]
	whileFunc func(v T) bool
	done      bool
	current   T
}

// TakeWhile makes Iterable with elements in which whileFunc is true from ahead.
//
//   itb := gcf.FromSlice([]{1, 2, 3})
//   itb = gcf.TakeWhile(itb, func(v int) bool { return v <= 2 })
//
// If whileFunc is nil, returns original Iterable.
func TakeWhile[T any](itb Iterable[T], whileFunc func(v T) bool) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	if whileFunc == nil {
		return itb
	}
	return &takeWhileIterable[T]{itb, whileFunc}
}

func (itb *takeWhileIterable[T]) Iterator() Iterator[T] {
	return &takeWhileIterator[T]{itb.itb.Iterator(), itb.whileFunc, false, zero[T]()}
}

func (it *takeWhileIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if it.it.MoveNext() && it.whileFunc(it.it.Current()) {
		it.current = it.it.Current()
		return true
	}
	it.done = true
	it.current = zero[T]()
	return false
}

func (it *takeWhileIterator[T]) Current() T {
	return it.current
}
