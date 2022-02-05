package gcf

import (
	"constraints"
)

type rangeIterable[T constraints.Integer] struct {
	start T
	end   T
	step  T
}

type rangeIncrementIterator[T constraints.Integer] struct {
	start   T
	end     T
	step    T
	started bool
	iteratorItem[T]
}

type rangeDecrementIterator[T constraints.Integer] struct {
	start   T
	end     T
	step    T
	started bool
	iteratorItem[T]
}

// Range makes Iterable with increasing or decreasing elements according to step.
//
//   itb := gcf.Range(1, 10, 3)
//
// If step is positive, elements is enumerated from start to end with increment by step.
// If step is negative, elements is enumerated from start to end with decrement by step.
// If step is zero, returns error.
// If start equals to end, makes Iterable with one element.
// If direction of step is opposite to start and end, returns empty Iterable.
func Range[T constraints.Integer](start, end, step T) (Iterable[T], error) {
	if step == 0 {
		return nil, errStep0
	}
	if start == end {
		return FromSlice([]T{start}), nil
	}
	if step > 0 && start > end {
		return empty[T](), nil
	}
	if step < 0 && start < end {
		return empty[T](), nil
	}
	return &rangeIterable[T]{start, end, step}, nil
}

func (itb *rangeIterable[T]) Iterator() Iterator[T] {
	if itb.step > 0 {
		return &rangeIncrementIterator[T]{
			start: itb.start,
			end:   itb.end,
			step:  itb.step,
		}
	}
	return &rangeDecrementIterator[T]{
		start: itb.start,
		end:   itb.end,
		step:  itb.step,
	}
}

func (it *rangeIncrementIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if !it.started {
		it.started = true
		it.current = it.start
		return true
	}
	if it.current+it.step > it.end {
		it.MarkDone()
		return false
	}
	it.current += it.step
	return true
}

func (it *rangeIncrementIterator[T]) Current() T {
	return it.current
}

func (it *rangeDecrementIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	if !it.started {
		it.started = true
		it.current = it.start
		return true
	}
	it.started = true
	if it.current+it.step < it.end {
		it.MarkDone()
		return false
	}
	it.current += it.step
	return true
}

func (it *rangeDecrementIterator[T]) Current() T {
	return it.current
}
