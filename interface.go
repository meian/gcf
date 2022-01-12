package gcf

// Iteratable provide Iterator creation.
type Iteratable[T any] interface {
	// Iterator create Iterator[T] instance.
	Iterator() Iterator[T]
}

// Iterator provide iterative process for collections.
type Iterator[T any] interface {
	// MoveNext proceed element position to next element.
	//
	//   var itb Iterator[int]
	//   it := itb.Iterator()
	//   for it.MoveNext() {
	//   	v := it.Current()
	//   	// some processes for iteration values
	//   }
	//
	// Return true if next value is exists or false if no next value.
	MoveNext() bool
	// Current return current element value.
	//
	// Return current value if iteratable position or should get zero value if out of iteratable position.
	// Note that return zero value before MoveNext is called.
	Current() T
}
