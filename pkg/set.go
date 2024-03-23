package gollections

var empty any

type hashSet[T comparable] struct {
	elements map[T]any
}

type setIterator[T comparable] struct {
	set            *hashSet[T]
	elements       []T
	currentElement T
	index          int
	removed        bool
}

func NewSet[T comparable](elements ...T) Set[T] {
	set := &hashSet[T]{
		elements: make(map[T]any),
	}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

func (ref *hashSet[T]) Add(element T, elements ...T) {
	ref.elements[element] = empty
	for _, element := range elements {
		ref.elements[element] = empty
	}
}

func (ref *hashSet[T]) AddAll(collection Collection[T]) {
	for it := collection.Iterator(); it.HasNext(); {
		ref.Add(it.Next())
	}
}

func (ref *hashSet[T]) Clear() {
	ref.elements = make(map[T]any)
}

func (ref *hashSet[T]) Contains(element T) bool {
	_, found := ref.elements[element]
	return found
}

func (ref *hashSet[T]) ContainsAll(collection Collection[T]) bool {
	found := 0
	for element := range ref.elements {
		if collection.Contains(element) {
			found++
		}
	}
	return found == collection.Size()
}

func (ref *hashSet[T]) ContainsAny(collection Collection[T]) bool {
	for element := range ref.elements {
		if collection.Contains(element) {
			return true
		}
	}
	return false
}

func (ref *hashSet[T]) IsEmpty() bool {
	return ref.Size() == 0
}

func (ref *hashSet[T]) Iterator() Iterator[T] {
	return &setIterator[T]{
		set:      ref,
		elements: ref.ToArray(),
	}
}

func (ref *hashSet[T]) Remove(element T, elements ...T) {
	delete(ref.elements, element)
	for _, element := range elements {
		delete(ref.elements, element)
	}
}

func (ref *hashSet[T]) RemoveAll(collection Collection[T]) {
	for element := range ref.elements {
		if collection.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashSet[T]) RemoveIf(predicate Predicate[T]) {
	for element := range ref.elements {
		if predicate(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashSet[T]) Retains(element T, elements ...T) {
	retainedElements := NewSet[T]()
	retainedElements.Add(element, elements...)
	for element := range ref.elements {
		if !retainedElements.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashSet[T]) RetainsAll(collection Collection[T]) {
	for element := range ref.elements {
		if !collection.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashSet[T]) RetainsIf(predicate Predicate[T]) {
	for element := range ref.elements {
		if !predicate(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashSet[T]) Size() int {
	return len(ref.elements)
}

func (ref *hashSet[T]) ToArray() []T {
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
	ref.set.Remove(ref.currentElement)
	ref.removed = true
}
