package gollections

type Iterator[E comparable] interface {

	// ForEachRemaining performs the given action for each remaining element until all elements have been processed.
	ForEachRemaining(action Consumer[E])

	// HasNext returns true if the iteration has more elements.
	HasNext() bool

	// Next returns the next element in the iteration.
	Next() E

	// Remove removes from the underlying collection the last element returned by this iterator.
	Remove()
}
