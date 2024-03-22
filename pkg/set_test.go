package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	collection := gollections.NewSet[string]()
	assertSet(t, collection)

	collection.Add("A")
	assertSet(t, collection, "A")

	collection.Add("B")
	assertSet(t, collection, "A", "B")

	collection.Add("C")
	assertSet(t, collection, "A", "B", "C")

	collection.Add("A")
	assertSet(t, collection, "A", "B", "C")

	collection.Add("B", "C", "D")
	assertSet(t, collection, "A", "B", "C", "D")

	collection.Add("E", "F")
	assertSet(t, collection, "A", "B", "C", "D", "E", "F")
}

func TestSet_AddAll(t *testing.T) {
	source := gollections.NewSet[string]()
	target := gollections.NewSet[string]()
	assertSet(t, source)
	assertSet(t, target)

	source.Add("A", "B", "C")
	assertSet(t, source, "A", "B", "C")
	assertSet(t, target)

	target.AddAll(source)
	assertSet(t, source, "A", "B", "C")
	assertSet(t, target, "A", "B", "C")

	source.Clear()
	assertSet(t, source)
	assertSet(t, target, "A", "B", "C")
}

func TestSet_Clear(t *testing.T) {
	collection := gollections.NewSet[string]()
	assertSet(t, collection)

	collection.Add("A", "B", "C")
	assertSet(t, collection, "A", "B", "C")

	collection.Clear()
	assertSet(t, collection)
}

func TestSet_Contains(t *testing.T) {
	collection := gollections.NewSet[string]()
	collection.Add("A", "B", "C")
	assert.True(t, collection.Contains("A"))
	assert.True(t, collection.Contains("B"))
	assert.True(t, collection.Contains("C"))
	assert.False(t, collection.Contains("D"))
	assert.False(t, collection.Contains("E"))
}

func TestSet_ContainsAll(t *testing.T) {
	collection := gollections.NewSet("A", "B", "C")
	assert.True(t, collection.ContainsAll(gollections.NewSet("A")))
	assert.True(t, collection.ContainsAll(gollections.NewSet("A", "B")))
	assert.True(t, collection.ContainsAll(gollections.NewSet("A", "B", "C")))
	assert.False(t, collection.ContainsAll(gollections.NewSet("A", "B", "C", "D")))
	assert.False(t, collection.ContainsAll(gollections.NewSet("A", "B", "C", "D", "E")))
}

func TestSet_ContainsAny(t *testing.T) {
	collection := gollections.NewSet("D", "E")
	assert.False(t, collection.ContainsAny(gollections.NewSet("A")))
	assert.False(t, collection.ContainsAny(gollections.NewSet("A", "B")))
	assert.False(t, collection.ContainsAny(gollections.NewSet("A", "B", "C")))
	assert.True(t, collection.ContainsAny(gollections.NewSet("A", "B", "C", "D")))
	assert.True(t, collection.ContainsAny(gollections.NewSet("A", "B", "C", "D", "E")))
}

func TestSet_IsEmpty(t *testing.T) {
	collection := gollections.NewSet[string]()
	assert.True(t, collection.IsEmpty())

	collection.Add("X")
	assert.False(t, collection.IsEmpty())

	collection.Remove("X")
	assert.True(t, collection.IsEmpty())
}

func TestSet_Iterator(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	collection := gollections.NewSet(elements...)

	for it := collection.Iterator(); it.HasNext(); {
		assert.Contains(t, elements, it.Next())
	}
	assertSet(t, collection, "A", "B", "C", "D", "E")

	for it := collection.Iterator(); it.HasNext(); {
		it.Next()
		it.Remove()
	}
	assertSet(t, collection)
}

func TestSet_Iterator_NoSuchElement(t *testing.T) {
	collection := gollections.NewSet[string]()
	it := collection.Iterator()
	assert.PanicsWithValue(t, "has no more elements to iterate", func() {
		it.Next()
	})
}

func TestSet_Iterator_RemoveWithoutCallNext(t *testing.T) {
	collection := gollections.NewSet[string]()
	it := collection.Iterator()
	assert.PanicsWithValue(t, "next method has not been called", func() {
		it.Remove()
	})
}

func TestSet_Iterator_RemoveCalledTwice(t *testing.T) {
	collection := gollections.NewSet("X")
	it := collection.Iterator()
	it.Next()
	it.Remove()
	assert.PanicsWithValue(t, "remove method has already been called after the last call to the next method", func() {
		it.Remove()
	})
}

func TestSet_Remove(t *testing.T) {
	collection := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, collection, "A", "B", "C", "D", "E")

	collection.Remove("C")
	assertSet(t, collection, "A", "B", "D", "E")

	collection.Remove("B", "D", "E")
	assertSet(t, collection, "A")

	collection.Remove("A")
	assertSet(t, collection)
}

func TestSet_RemoveAll(t *testing.T) {
	collection := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, collection, "A", "B", "C", "D", "E")

	collection.RemoveAll(gollections.NewSet("A", "C", "E"))
	assertSet(t, collection, "B", "D")

	collection.RemoveAll(gollections.NewSet("A", "B", "C"))
	assertSet(t, collection, "D")

	collection.RemoveAll(gollections.NewSet("D", "E"))
	assertSet(t, collection)
}

func TestSet_RemoveIf(t *testing.T) {
	collection := gollections.NewSet(1, 2, 3, 4, 5, 6, 7, 8)
	assertSet(t, collection, 1, 2, 3, 4, 5, 6, 7, 8)

	collection.RemoveIf(func(element int) bool {
		return element%2 != 0
	})
	assertSet(t, collection, 2, 4, 6, 8)

	collection.RemoveIf(func(element int) bool {
		return element%4 == 0
	})
	assertSet(t, collection, 2, 6)
}

func TestSet_Rentains(t *testing.T) {
	collection := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, collection, "A", "B", "C", "D", "E")

	collection.Retains("A", "B", "C")
	assertSet(t, collection, "A", "B", "C")

	collection.Retains("B")
	assertSet(t, collection, "B")
}

func TestSet_RentainsAll(t *testing.T) {
	collection := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, collection, "A", "B", "C", "D", "E")

	collection.RetainsAll(gollections.NewSet("A", "B", "C"))
	assertSet(t, collection, "A", "B", "C")

	collection.RetainsAll(gollections.NewSet[string]())
	assertSet(t, collection)
}

func TestSet_RentainsIf(t *testing.T) {
	collection := gollections.NewSet(1, 2, 3, 4, 5, 6, 7, 8)
	assertSet(t, collection, 1, 2, 3, 4, 5, 6, 7, 8)

	collection.RetainsIf(func(element int) bool {
		return element%2 == 0
	})
	assertSet(t, collection, 2, 4, 6, 8)

	collection.RetainsIf(func(element int) bool {
		return element%4 != 0
	})
	assertSet(t, collection, 2, 6)
}

func TestSet_Size(t *testing.T) {
	collection := gollections.NewSet[string]()
	assert.Zero(t, collection.Size())

	collection.Add("A")
	assert.Equal(t, 1, collection.Size())

	collection.Add("B")
	assert.Equal(t, 2, collection.Size())

	collection.Add("C", "D")
	assert.Equal(t, 4, collection.Size())

	collection.Remove("C")
	assert.Equal(t, 3, collection.Size())

	collection.Remove("D", "A", "B")
	assert.Zero(t, collection.Size())
}

func TestSet_ToArray(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	collection := gollections.NewSet(elements...)
	assert.ElementsMatch(t, elements, collection.ToArray())
}

func assertSet[T comparable](t *testing.T, collection gollections.Set[T], elements ...T) {
	if len(elements) == 0 {
		assert.Zero(t, collection.Size())
		assert.True(t, collection.IsEmpty())
	} else {
		assert.ElementsMatch(t, elements, collection.ToArray())
		assert.Equal(t, len(elements), collection.Size())
		assert.False(t, collection.IsEmpty())
	}
}
