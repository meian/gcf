package gcf

type reverseIterable[T any] struct {
	itb Iterable[T]
}

type reverseIterator[T any] struct {
	it    Iterator[T]
	built bool
	iteratorItem[T]
}

// Reverse makes Iterable with reverse order elements.
//
//   itb := gcf.FromSlice([]int{1, 2, 3})
//   itb = gcf.Reverse(itb)
func Reverse[T any](itb Iterable[T]) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	// reverse in reverse is original Iterable.
	if itbr, ok := itb.(*reverseIterable[T]); ok {
		return itbr.itb
	}
	return &reverseIterable[T]{itb}
}

func (itb *reverseIterable[T]) Iterator() Iterator[T] {
	return &reverseIterator[T]{
		it: itb.itb.Iterator(),
	}
}

func (it *reverseIterator[T]) MoveNext() bool {
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

func (it *reverseIterator[T]) Current() T {
	return it.current
}

func (it *reverseIterator[T]) build() {
	s := iteratorToSlice(it.it)
	len := len(s)
	for i := 0; i < len/2; i++ {
		s[i], s[len-i-1] = s[len-i-1], s[i]
	}
	it.it = makeSliceIterator(s)
	it.built = true
}
