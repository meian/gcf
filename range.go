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
	end     T
	step    T
	started bool
	done    bool
	current T
}

type rangeDecrementIterator[T constraints.Integer] struct {
	end     T
	step    T
	started bool
	done    bool
	current T
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
		return &rangeIncrementIterator[T]{itb.end, itb.step, false, false, itb.start - itb.step}
	}
	return &rangeDecrementIterator[T]{itb.end, itb.step, false, false, itb.start - itb.step}
}

func (it *rangeIncrementIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	it.started = true
	it.current += it.step
	if it.current > it.end {
		it.done = true
		return false
	}
	return true
}

func (it *rangeIncrementIterator[T]) Current() T {
	if !it.started || it.done {
		return zero[T]()
	}
	return it.current
}

func (it *rangeDecrementIterator[T]) MoveNext() bool {
	if it.done {
		return false
	}
	it.started = true
	it.current += it.step
	if it.current < it.end {
		it.done = true
		return false
	}
	return true
}

func (it *rangeDecrementIterator[T]) Current() T {
	if !it.started || it.done {
		return zero[T]()
	}
	return it.current
}
