package gcf

type emptyIterable[T any] struct {
}

type emptyIterator[T any] struct {
	current T
}

func empty[T any]() Iterable[T] {
	return emptyIterable[T]{}
}

func (itb emptyIterable[T]) Iterator() Iterator[T] {
	return emptyIter[T]()
}

func emptyIter[T any]() Iterator[T] {
	return &emptyIterator[T]{}
}

func (it *emptyIterator[T]) MoveNext() bool {
	return false
}

func (it *emptyIterator[T]) Current() T {
	return it.current
}
