package gollections

type ListIterator[E comparable] interface {
	Iterator[E]

	// Add inserts the specified element into the list.
	Add(element E)

	// HasPrevious returns true if this list iterator has more elements when traversing the list in the reverse direction.
	HasPrevious() bool

	// NextIndex returns the index of the element that would be returned by a subsequent call to Next().
	NextIndex() int

	// Previous returns the previous element in the list and moves the cursor position backwards.
	Previous() E

	// PreviousIndex returns the index of the element that would be returned by a subsequent call to Previous().
	PreviousIndex() int

	// Set replaces the last element returned by Next() or Previous() with the specified element.
	Set(element E)
}
