package gollections

var empty any

type set[T comparable] struct {
	items map[T]any
}

func NewSet[T comparable]() Set[T] {
	return &set[T]{
		items: make(map[T]any),
	}
}

func (ref *set[T]) Add(element T, elements ...T) {
	ref.items[element] = empty
	for _, element := range elements {
		ref.items[element] = empty
	}
}

func (ref *set[T]) AddAll(collection Collection[T]) {

}

func (ref *set[T]) Clear() {
	ref.items = make(map[T]any)
}

func (ref *set[T]) Contains(element T) bool {
	_, found := ref.items[element]
	return found
}

func (ref *set[T]) ContainsAll(collection Collection[T]) bool {
	return false
}

func (ref *set[T]) ContainsAny(collection Collection[T]) bool {
	return false
}

func (ref *set[T]) IsEmpty() bool {
	return ref.Size() == 0
}

func (ref *set[T]) Iterator() Iterator[T] {
	return nil
}

func (ref *set[T]) Remove(element T, elements ...T) {

}

func (ref *set[T]) RemoveAll(collection Collection[T]) {

}

func (ref *set[T]) RemoveIf(predicate Predicate[T]) {

}

func (ref *set[T]) Retains(element T, elements ...T) {

}

func (ref *set[T]) RetainsAll(collection Collection[T]) {

}

func (ref *set[T]) RetainsIf(predicate Predicate[T]) {

}

func (ref *set[T]) Size() int {
	return len(ref.items)
}

func (ref *set[T]) ToArray() []T {
	return nil
}
