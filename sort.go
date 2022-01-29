package gcf

import (
	"constraints"
	"sort"
)

type sortIterable[T constraints.Ordered] struct {
	itb      Iterable[T]
	toSorter func([]T) sorter[T]
}

type sortIterator[T constraints.Ordered] struct {
	it       Iterator[T]
	toSorter func([]T) sorter[T]
	built    bool
	current  T
}

type sorter[T any] interface {
	sort.Interface
	Sort()
}

type ascSlice[T constraints.Ordered] []T

func (s ascSlice[T]) Len() int { return len(s) }

func (s ascSlice[T]) Less(i, j int) bool { return s[i] < s[j] }

func (s ascSlice[T]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s ascSlice[T]) Sort() { sort.Sort(s) }

// SortAsc makes Iterable with sorted by ascending elements.
//
//   itb := gcf.FromSlice([]int{1, 3, 2})
//   itb = gcf.SortAsc(itb)
func SortAsc[T constraints.Ordered](itb Iterable[T]) Iterable[T] {
	if itb == nil {
		return empty[T]()
	}
	toSorter := func(s []T) sorter[T] { return ascSlice[T](s) }
	return &sortIterable[T]{itb, toSorter}
}

func (itb *sortIterable[T]) Iterator() Iterator[T] {
	return &sortIterator[T]{itb.itb.Iterator(), itb.toSorter, false, zero[T]()}
}

func (it *sortIterator[T]) MoveNext() bool {
	if !it.built {
		it.buildSort()
	}
	if !it.it.MoveNext() {
		it.current = zero[T]()
		return false
	}
	it.current = it.it.Current()
	return true
}

func (it *sortIterator[T]) Current() T {
	return it.current
}

func (it *sortIterator[T]) buildSort() {
	s := iteratorToSlice(it.it)
	it.toSorter(s).Sort()
	it.it = makeSliceIterator(s)
	it.built = true
}
