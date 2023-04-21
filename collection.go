package gollections

type Collection[E comparable] interface {
	Iterable[E]

	// Add ensures that this collection contains the specified element.
	Add(element E) bool

	// AddAll adds all the elements in the specified collection to this collection.
	AddAll(collection Collection[E]) bool

	// Clear removes all the elements from this collection.
	Clear()

	// Contains returns true if this collection contains the specified element.
	Contains(element E) bool

	// ContainsAll returns true if this collection contains all the elements in the specified collection.
	ContainsAll(collection Collection[E])

	// IsEmpty returns true if this collection contains no elements.
	IsEmpty() bool

	// ParallelStream returns a possibly parallel Stream with this collection as its source.
	//ParallelStream() Stream[E]

	// Remove removes a single instance of the specified element from this collection, if it is present
	Remove(element E) bool

	// RemoveAll removes all of this collection's elements that are also contained in the specified collection.
	RemoveAll(collection Collection[E]) bool

	// RemoveIf removes all the elements of this collection that satisfy the given predicate.
	RemoveIf(filter Predicate[E]) bool

	// RetainAll retains only the elements in this collection that are contained in the specified collection.
	RetainAll(collection Collection[E])

	// Size returns the number of elements in this collection.
	Size() int

	// Stream returns a sequential Stream with this collection as its source.
	//Stream() Stream[E]

	// ToArray returns an array containing all the elements in this collection.
	ToArray() []E
}
