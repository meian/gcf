package gcf

type distinctIterable[T comparable] struct {
	itb Iterable[T]
}

type distinctIterator[T comparable] struct {
	it   Iterator[T]
	past map[T]struct{}
	iteratorItem[T]
}

// Distinct makes Iterable contains unique elements.
// Inner elements is restrict by comparable constraint.
//
//	itb := gcf.FromSlice([]int{1, 2, 3, 3, 4, 2, 5})
//	itb = gcf.Distinct(itb)
//
// Currently, result order is determined, but on spec, is undefined.
func Distinct[T comparable](itb Iterable[T]) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	// Because no change has with Distinct combined Distinct, original Iterable is returned
	if itbd, ok := itb.(*distinctIterable[T]); ok {
		return itbd
	}
	return &distinctIterable[T]{itb}
}

func (itb *distinctIterable[T]) Iterator() Iterator[T] {
	return &distinctIterator[T]{
		it:   itb.itb.Iterator(),
		past: map[T]struct{}{},
	}
}

func (it *distinctIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	for it.it.MoveNext() {
		c := it.it.Current()
		if _, ok := it.past[c]; ok {
			continue
		}
		it.past[c] = struct{}{}
		it.current = c
		return true
	}
	it.MarkDone()
	return false
}

func (it *distinctIterator[T]) Current() T {
	return it.current
}
