package gollections

type List[E comparable] interface {
	Collection[E]

	// AddAt inserts the specified element at the specified position in this list.
	AddAt(index int, element E)

	// AddAllAt inserts all the elements in the specified collection into this list at the specified position.
	AddAllAt(index int, collection Collection[E]) bool

	// CopyOf returns an unmodifiable List containing the elements of the given Collection, in its iteration order.
	CopyOf(collection Collection[E]) List[E]

	// Get returns the element at the specified position in this list.
	Get(index int)

	// IndexOf returns true if this list contains no elements.
	IndexOf(element E) int

	// LastIndexOf returns the index of the last occurrence of the specified element in this list,
	// or -1 if this list does not contain the element.
	LastIndexOf(element E) int

	// ListIterator returns a list iterator over the elements in this list (in proper sequence).
	ListIterator() ListIterator[E]

	// ListIteratorAt returns a list iterator over the elements in this list (in proper sequence),
	// starting at the specified position in the list.
	ListIteratorAt(index int) ListIterator[E]

	// RemoveAt remove removes the element at the specified position in this list.
	RemoveAt(index int)

	// ReplaceAll replaces each element of this list with the result of applying the operator to that element.
	ReplaceAll(operator UnaryOperator[E])

	// Set replaces the element at the specified position in this list with the specified element.
	Set(index int, element E) E

	// Sort sorts this list according to the order induced by the specified Comparator.
	Sort(comparator Comparator[E])

	// SubList returns a view of the portion of this list between the specified fromIndex, inclusive, and toIndex, exclusive.
	SubList(fromIndex, toIndex int) List[E]
}
