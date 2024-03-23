package gollections

type Collection[T any] interface {
	Add(element T, elements ...T)
	AddAll(collection Collection[T])
	Clear()
	Contains(element T) bool
	ContainsAll(collection Collection[T]) bool
	ContainsAny(collection Collection[T]) bool
	IsEmpty() bool
	Iterator() Iterator[T]
	Remove(element T, elements ...T)
	RemoveAll(collection Collection[T])
	RemoveIf(predicate Predicate[T])
	Retains(element T, elements ...T)
	RetainsAll(collection Collection[T])
	RetainsIf(predicate Predicate[T])
	Size() int
	ToArray() []T
}

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
