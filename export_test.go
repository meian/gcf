package gcf

// IsEmptyIterable detects itb is emptyIterable.
func IsEmptyIterable[T any](itb Iterable[T]) bool {
	return isEmpty(itb)
}
