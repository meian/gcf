package gcf

type mapIterable[T any, R any] struct {
	itb     Iterable[T]
	mapFunc func(T) R
}

type mapIterator[T any, R any] struct {
	it      Iterator[T]
	mapFunc func(T) R
	current R
}

// Map makes Iterable in elements convert by mapFunc.
//
//   itbs := gcf.Func([]string{"a", "ab", "abc"})
//   itbi := gcf.Map(itbs, func(v string) int { return len(v) })
//
// If mapFunc is nil, return Iterable in zero value elements.
func Map[T any, R any](itb Iterable[T], mapFunc func(v T) R) Iterable[R] {
	if itb == nil {
		return empty[R]()
	}
	if mapFunc == nil {
		r := zero[R]()
		mapFunc = func(v T) R { return r }
	}
	return &mapIterable[T, R]{itb, mapFunc}
}

func (itb *mapIterable[T, R]) Iterator() Iterator[R] {
	return &mapIterator[T, R]{itb.itb.Iterator(), itb.mapFunc, zero[R]()}
}

func (it *mapIterator[T, R]) MoveNext() bool {
	if !it.it.MoveNext() {
		it.current = zero[R]()
		return false
	}
	it.current = it.mapFunc(it.it.Current())
	return true
}

func (it *mapIterator[T, R]) Current() R {
	return it.current
}

type flatMapIterable[T any, R any] struct {
	itb     Iterable[T]
	mapFunc func(T) []R
}

type flatMapIterator[T any, R any] struct {
	it      Iterator[T]
	mapFunc func(T) []R
	its     Iterator[R]
	current R
}

// FlatMap makes Iterable in elements in slice converted by mapFunc.
//
//   itbs := gcf.Func([]string{"a", "ab", "abc"})
//   itbi := gcf.Map(itbs, func(v string) int[] {
//       var r := make([]int, 0)
//       for _, c := range []rune(v) {
//           r = append(r, int(c))
//       }
//   })
//
// If mapFunc is nil, return empty Iterable.
func FlatMap[T any, R any](itb Iterable[T], mapFunc func(v T) []R) Iterable[R] {
	if itb == nil {
		return empty[R]()
	}
	if mapFunc == nil {
		return empty[R]()
	}
	return &flatMapIterable[T, R]{itb, mapFunc}
}

func (itb *flatMapIterable[T, R]) Iterator() Iterator[R] {
	return &flatMapIterator[T, R]{itb.itb.Iterator(), itb.mapFunc, emptyIter[R](), zero[R]()}
}

func (it *flatMapIterator[T, R]) MoveNext() bool {
	for {
		if it.its.MoveNext() {
			it.current = it.its.Current()
			return true
		}
		if !it.it.MoveNext() {
			break
		}
		if s := it.mapFunc(it.it.Current()); len(s) > 0 {
			it.its = makeSliceIterator(s)
		}
	}
	it.current = zero[R]()
	return false
}

func (it *flatMapIterator[T, R]) Current() R {
	return it.current
}
