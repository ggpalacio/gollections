package gollections

type Iterable[T comparable] interface {

	// ForEach performs the given action for each element of the Iterable until all elements have been processed.
	ForEach(action Consumer[T])

	// Iterator returns an iterator over elements of type T.
	Iterator() Iterator[T]

	// Spliterator creates a Spliterator over the elements described by this Iterable.
	//Spliterator() Spliterator[T]
}
