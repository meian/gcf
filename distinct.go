package gcf

type distinctIterable[T comparable] struct {
	itb Iterable[T]
}

type distinctIterator[T comparable] struct {
	it      Iterator[T]
	past    map[T]struct{}
	current T
}

// Distinct makes Iterable contains unique elements.
//
//   itb := gcf.FromSlice([]int{1, 2, 3, 3, 4, 2, 5})
//   itb = gcf.Distinct(itb)
//
// Currently, result order is determined, but on spec, is undefined.
func Distinct[T comparable](itb Iterable[T]) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	// Because no change has with Distinct combined Distinct, original Iterable is returned
	if itbd, ok := itb.(*distinctIterable[T]); ok {
		return itbd
	}
	return &distinctIterable[T]{itb}
}

func (itb *distinctIterable[T]) Iterator() Iterator[T] {
	return &distinctIterator[T]{itb.itb.Iterator(), map[T]struct{}{}, zero[T]()}
}

func (it *distinctIterator[T]) MoveNext() bool {
	for it.it.MoveNext() {
		c := it.it.Current()
		if _, ok := it.past[c]; ok {
			continue
		}
		it.past[c] = struct{}{}
		it.current = c
		return true
	}
	it.current = zero[T]()
	return false
}

func (it *distinctIterator[T]) Current() T {
	return it.current
}
