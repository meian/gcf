package gcf

// iteratorItem provide Iterator standard common items.
type iteratorItem[T any] struct {
	current T
	done    bool
}

// MarkDone sets state Iterator has been read to end.
// By call MarkDone, done is set true and current is set zero value.
func (iti *iteratorItem[T]) MarkDone() {
	iti.done = true
	iti.current = zero[T]()
}
