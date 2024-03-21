package gollections

type Iterator[T any] interface {
	HasNext() bool
	Next() T
	Remove()
}

type Predicate[T any] interface {
	Test(value T) bool
}

type Comparator[T any] interface {
	Compare(value1, value2 T) int
}

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

type List[T any] interface {
	Collection[T]
	AddAt(index int, element T, elements ...T)
	AddAllAt(index int, collection Collection[T])
	Get(index int) T
	IndexOf(element T) int
	LastIndexOf(element T) int
	RemoveAt(index int) T
	RemoveRange(fromIndex, toIndex int)
	Set(index int, element T)
	Sort(comparator Comparator[T])
	SubList(fromIndex, toIndex int)
}
