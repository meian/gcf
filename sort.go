package gcf

import (
	"constraints"
	"sort"
)

type sortIterable[T any] struct {
	itb      Iterable[T]
	toSorter func([]T) sorter[T]
}

type sortIterator[T any] struct {
	it       Iterator[T]
	toSorter func([]T) sorter[T]
	built    bool
	iteratorItem[T]
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

type descSlice[T constraints.Ordered] []T

func (s descSlice[T]) Len() int { return len(s) }

func (s descSlice[T]) Less(i, j int) bool { return s[i] > s[j] }

func (s descSlice[T]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s descSlice[T]) Sort() { sort.Sort(s) }

type funcSorter[T any] struct {
	slice []T
	less  func(x, y T) bool
}

func (s funcSorter[T]) Len() int { return len(s.slice) }

func (s funcSorter[T]) Less(i, j int) bool { return s.less(s.slice[i], s.slice[j]) }

func (s funcSorter[T]) Swap(i, j int) { s.slice[i], s.slice[j] = s.slice[j], s.slice[i] }

func (s funcSorter[T]) Sort() { sort.Sort(s) }

// SortAsc makes Iterable with sorted by ascending elements.
//
//   itb := gcf.FromSlice([]int{1, 3, 2})
//   itb = gcf.SortAsc(itb)
func SortAsc[T constraints.Ordered](itb Iterable[T]) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	toSorter := func(s []T) sorter[T] { return ascSlice[T](s) }
	return &sortIterable[T]{itb, toSorter}
}

// SortDesc makes Iterable with sorted by descending elements.
//
//   itb := gcf.FromSlice([]int{1, 3, 2})
//   itb = gcf.SortDesc(itb)
func SortDesc[T constraints.Ordered](itb Iterable[T]) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	toSorter := func(s []T) sorter[T] { return descSlice[T](s) }
	return &sortIterable[T]{itb, toSorter}
}

// SortBy makes iterable with elements sorted by provided less function.
//
//   type data struct { id int }
//   itb := gcf.FromSlice([]data{{1}, {3}, {2}})
//   itb = gcf.SortBy(itb, func(x, y data) bool { return x.id < y.id })
//
// The less function takes x element and y element, and returns true if x is less than y.
func SortBy[T any](itb Iterable[T], less func(x, y T) bool) Iterable[T] {
	if isEmpty(itb) {
		return orEmpty(itb)
	}
	toSorter := func(s []T) sorter[T] { return funcSorter[T]{s, less} }
	return &sortIterable[T]{itb, toSorter}
}

func (itb *sortIterable[T]) Iterator() Iterator[T] {
	return &sortIterator[T]{
		it:       itb.itb.Iterator(),
		toSorter: itb.toSorter,
	}
}

func (it *sortIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if !it.built {
		it.buildSort()
	}
	if !it.it.MoveNext() {
		it.MarkDone()
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
