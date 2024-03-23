package gollections

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
