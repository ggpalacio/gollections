package gollections

var empty any

type set[T comparable] struct {
	elements map[T]any
}

type setIterator[T comparable] struct {
	collection     *set[T]
	elements       []T
	currentElement T
	index          int
	removed        bool
}

func NewSet[T comparable](elements ...T) Set[T] {
	collection := &set[T]{
		elements: make(map[T]any),
	}
	for _, element := range elements {
		collection.Add(element)
	}
	return collection
}

func (ref *set[T]) Add(element T, elements ...T) {
	ref.elements[element] = empty
	for _, element := range elements {
		ref.elements[element] = empty
	}
}

func (ref *set[T]) AddAll(collection Collection[T]) {
	for it := collection.Iterator(); it.HasNext(); {
		ref.Add(it.Next())
	}
}

func (ref *set[T]) Clear() {
	ref.elements = make(map[T]any)
}

func (ref *set[T]) Contains(element T) bool {
	_, found := ref.elements[element]
	return found
}

func (ref *set[T]) ContainsAll(collection Collection[T]) bool {
	found := 0
	for element := range ref.elements {
		if collection.Contains(element) {
			found++
		}
	}
	return found == collection.Size()
}

func (ref *set[T]) ContainsAny(collection Collection[T]) bool {
	for element := range ref.elements {
		if collection.Contains(element) {
			return true
		}
	}
	return false
}

func (ref *set[T]) IsEmpty() bool {
	return ref.Size() == 0
}

func (ref *set[T]) Iterator() Iterator[T] {
	return &setIterator[T]{
		collection: ref,
		elements:   ref.ToArray(),
	}
}

func (ref *set[T]) Remove(element T, elements ...T) {
	delete(ref.elements, element)
	for _, element := range elements {
		delete(ref.elements, element)
	}
}

func (ref *set[T]) RemoveAll(collection Collection[T]) {
	for element := range ref.elements {
		if collection.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *set[T]) RemoveIf(predicate Predicate[T]) {
	for element := range ref.elements {
		if predicate(element) {
			ref.Remove(element)
		}
	}
}

func (ref *set[T]) Retains(element T, elements ...T) {
	retainedElements := NewSet[T]()
	retainedElements.Add(element, elements...)
	for element := range ref.elements {
		if !retainedElements.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *set[T]) RetainsAll(collection Collection[T]) {
	for element := range ref.elements {
		if !collection.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *set[T]) RetainsIf(predicate Predicate[T]) {
	for element := range ref.elements {
		if !predicate(element) {
			ref.Remove(element)
		}
	}
}

func (ref *set[T]) Size() int {
	return len(ref.elements)
}

func (ref *set[T]) ToArray() []T {
	array := make([]T, len(ref.elements))
	index := 0
	for element := range ref.elements {
		array[index] = element
		index++
	}
	return array
}

func (ref *setIterator[T]) HasNext() bool {
	return ref.index < len(ref.elements)
}

func (ref *setIterator[T]) Next() T {
	if !ref.HasNext() {
		panic("has no more elements to iterate")
	}
	ref.currentElement = ref.elements[ref.index]
	ref.index++
	ref.removed = false
	return ref.currentElement
}

func (ref *setIterator[T]) Remove() {
	if ref.index == 0 {
		panic("next method has not been called")
	}
	if ref.removed {
		panic("remove method has already been called after the last call to the next method")
	}
	ref.collection.Remove(ref.currentElement)
	ref.removed = true
}
