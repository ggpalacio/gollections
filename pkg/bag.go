package gollections

type Bag[T comparable] interface {
	Collection[T]
	AddMany(element T, count int)
	Count(element T) int
	RemoveEvery(element T, elements ...T)
	RemoveMany(element T, count int)
	Set(element T, count int)
	ToSet() Set[T]
}

type hashBag[T comparable] struct {
	elements map[T]int
}

func NewBag[T comparable](elements ...T) Bag[T] {
	bag := &hashBag[T]{
		elements: make(map[T]int),
	}
	for _, element := range elements {
		bag.Add(element)
	}
	return bag
}

func (ref *hashBag[T]) Add(element T, elements ...T) {
	ref.AddMany(element, 1)
	for _, element := range elements {
		ref.AddMany(element, 1)
	}
}

func (ref *hashBag[T]) AddAll(collection Collection[T]) {
	for it := collection.Iterator(); it.HasNext(); {
		ref.Add(it.Next())
	}
}

func (ref *hashBag[T]) AddMany(element T, count int) {
	if count <= 0 {
		return
	}
	ref.elements[element] = ref.elements[element] + count
}

func (ref *hashBag[T]) Clear() {
	ref.elements = make(map[T]int)
}

func (ref *hashBag[T]) Count(element T) int {
	return ref.elements[element]
}

func (ref *hashBag[T]) Contains(element T) bool {
	_, found := ref.elements[element]
	return found
}

func (ref *hashBag[T]) ContainsAll(collection Collection[T]) bool {
	found := 0
	for it := collection.Iterator(); it.HasNext(); {
		if ref.Contains(it.Next()) {
			found++
		}
	}
	return found == collection.Size()
}

func (ref *hashBag[T]) ContainsAny(collection Collection[T]) bool {
	for it := collection.Iterator(); it.HasNext(); {
		if ref.Contains(it.Next()) {
			return true
		}
	}
	return false
}

func (ref *hashBag[T]) IsEmpty() bool {
	return ref.Size() == 0
}

func (ref *hashBag[T]) Iterator() Iterator[T] {
	return &collectionIterator[T]{
		collection: ref,
		elements:   ref.ToArray(),
	}
}

func (ref *hashBag[T]) Remove(element T, elements ...T) {
	ref.RemoveMany(element, 1)
	for _, element := range elements {
		ref.RemoveMany(element, 1)
	}
}

func (ref *hashBag[T]) RemoveAll(collection Collection[T]) {
	for it := collection.Iterator(); it.HasNext(); {
		element := it.Next()
		if collection.Contains(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashBag[T]) RemoveEvery(element T, elements ...T) {
	delete(ref.elements, element)
	for _, element := range elements {
		delete(ref.elements, element)
	}
}

func (ref *hashBag[T]) RemoveIf(predicate Predicate[T]) {
	for element := range ref.elements {
		if predicate(element) {
			ref.Remove(element)
		}
	}
}

func (ref *hashBag[T]) RemoveMany(element T, count int) {
	if count <= 0 {
		return
	}
	ref.elements[element] = ref.elements[element] - count
	if ref.Count(element) <= 0 {
		delete(ref.elements, element)
	}
}

func (ref *hashBag[T]) Retains(element T, elements ...T) {
	retainedElements := NewSet[T]()
	retainedElements.Add(element, elements...)
	for element := range ref.elements {
		if !retainedElements.Contains(element) {
			ref.RemoveEvery(element)
		}
	}
}

func (ref *hashBag[T]) RetainsAll(collection Collection[T]) {
	for element := range ref.elements {
		if !collection.Contains(element) {
			ref.RemoveEvery(element)
		}
	}
}

func (ref *hashBag[T]) RetainsIf(predicate Predicate[T]) {
	for element := range ref.elements {
		if !predicate(element) {
			ref.RemoveEvery(element)
		}
	}
}

func (ref *hashBag[T]) Set(element T, count int) {
	if count <= 0 && ref.Contains(element) {
		ref.RemoveEvery(element)
	} else {
		ref.elements[element] = count
	}
}

func (ref *hashBag[T]) Size() int {
	size := 0
	for _, count := range ref.elements {
		size += count
	}
	return size
}

func (ref *hashBag[T]) ToArray() []T {
	array := make([]T, 0)
	for element, count := range ref.elements {
		for ; count > 0; count-- {
			array = append(array, element)
		}
	}
	return array
}

func (ref *hashBag[T]) ToSet() Set[T] {
	set := NewSet[T]()
	for element := range ref.elements {
		set.Add(element)
	}
	return set
}
