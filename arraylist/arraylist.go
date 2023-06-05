package arraylist

import "github.com/ggpalacio/gollections"

type arrayList[E comparable] struct {
	elements []E
}

func (ref *arrayList[E]) ForEach(action gollections.Consumer[E]) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Iterator() gollections.Iterator[E] {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Add(element E) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) AddAll(collection gollections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Contains(element E) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) ContainsAll(collection gollections.Collection[E]) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) IsEmpty() bool {
	return ref.Size() == 0
}

func (ref *arrayList[E]) Remove(element E) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) RemoveAll(collection gollections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) RemoveIf(filter gollections.Predicate[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) RetainAll(collection gollections.Collection[E]) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Size() int {
	return len(ref.elements)
}

func (ref *arrayList[E]) ToArray() []E {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) AddAt(index int, element E) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) AddAllAt(index int, collection gollections.Collection[E]) bool {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) CopyOf(collection gollections.Collection[E]) gollections.List[E] {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Get(index int) (E, bool) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) IndexOf(element E) int {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) LastIndexOf(element E) int {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) ListIterator() gollections.ListIterator[E] {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) ListIteratorAt(index int) gollections.ListIterator[E] {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) RemoveAt(index int) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) ReplaceAll(operator gollections.UnaryOperator[E]) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Set(index int, element E) E {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) Sort(comparator gollections.Comparator[E]) {
	//TODO implement me
	panic("implement me")
}

func (ref *arrayList[E]) SubList(fromIndex, toIndex int) gollections.List[E] {
	//TODO implement me
	panic("implement me")
}
