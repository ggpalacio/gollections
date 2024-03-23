package gollections

type collectionIterator[T comparable] struct {
	collection   Collection[T]
	elements     []T
	currentIndex int
	removed      bool
}

func (ref *collectionIterator[T]) HasNext() bool {
	return ref.currentIndex < len(ref.elements)
}

func (ref *collectionIterator[T]) Next() T {
	if !ref.HasNext() {
		panic("has no more elements to iterate")
	}
	ref.removed = false
	currentElement := ref.elements[ref.currentIndex]
	ref.currentIndex++
	return currentElement
}

func (ref *collectionIterator[T]) Remove() {
	if ref.currentIndex == 0 {
		panic("next method has not been called")
	}
	if ref.removed {
		panic("remove method has already been called after the last call to the next method")
	}
	ref.collection.Remove(ref.elements[ref.currentIndex-1])
	ref.removed = true
}
