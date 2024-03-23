package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	set := gollections.NewSet[string]()
	assertSet(t, set)

	set.Add("A")
	assertSet(t, set, "A")

	set.Add("B")
	assertSet(t, set, "A", "B")

	set.Add("C")
	assertSet(t, set, "A", "B", "C")

	set.Add("A")
	assertSet(t, set, "A", "B", "C")

	set.Add("B", "C", "D")
	assertSet(t, set, "A", "B", "C", "D")

	set.Add("E", "F")
	assertSet(t, set, "A", "B", "C", "D", "E", "F")
}

func TestSet_AddAll(t *testing.T) {
	set1 := gollections.NewSet[string]()
	set2 := gollections.NewSet[string]()
	assertSet(t, set1)
	assertSet(t, set2)

	set1.Add("A", "B", "C")
	assertSet(t, set1, "A", "B", "C")
	assertSet(t, set2)

	set2.AddAll(set1)
	assertSet(t, set1, "A", "B", "C")
	assertSet(t, set2, "A", "B", "C")

	set1.Clear()
	assertSet(t, set1)
	assertSet(t, set2, "A", "B", "C")
}

func TestSet_Clear(t *testing.T) {
	set := gollections.NewSet[string]()
	assertSet(t, set)

	set.Add("A", "B", "C")
	assertSet(t, set, "A", "B", "C")

	set.Clear()
	assertSet(t, set)
}

func TestSet_Contains(t *testing.T) {
	set := gollections.NewSet[string]()
	set.Add("A", "B", "C")
	assert.True(t, set.Contains("A"))
	assert.True(t, set.Contains("B"))
	assert.True(t, set.Contains("C"))
	assert.False(t, set.Contains("D"))
	assert.False(t, set.Contains("E"))
}

func TestSet_ContainsAll(t *testing.T) {
	set := gollections.NewSet("A", "B", "C")
	assert.True(t, set.ContainsAll(gollections.NewSet("A")))
	assert.True(t, set.ContainsAll(gollections.NewSet("A", "B")))
	assert.True(t, set.ContainsAll(gollections.NewSet("A", "B", "C")))
	assert.False(t, set.ContainsAll(gollections.NewSet("A", "B", "C", "D")))
	assert.False(t, set.ContainsAll(gollections.NewSet("A", "B", "C", "D", "E")))
}

func TestSet_ContainsAny(t *testing.T) {
	set := gollections.NewSet("D", "E")
	assert.False(t, set.ContainsAny(gollections.NewSet("A")))
	assert.False(t, set.ContainsAny(gollections.NewSet("A", "B")))
	assert.False(t, set.ContainsAny(gollections.NewSet("A", "B", "C")))
	assert.True(t, set.ContainsAny(gollections.NewSet("A", "B", "C", "D")))
	assert.True(t, set.ContainsAny(gollections.NewSet("A", "B", "C", "D", "E")))
}

func TestSet_IsEmpty(t *testing.T) {
	set := gollections.NewSet[string]()
	assert.True(t, set.IsEmpty())

	set.Add("X")
	assert.False(t, set.IsEmpty())

	set.Remove("X")
	assert.True(t, set.IsEmpty())
}

func TestSet_Iterator(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	set := gollections.NewSet(elements...)

	for it := set.Iterator(); it.HasNext(); {
		assert.Contains(t, elements, it.Next())
	}
	assertSet(t, set, "A", "B", "C", "D", "E")

	for it := set.Iterator(); it.HasNext(); {
		it.Next()
		it.Remove()
	}
	assertSet(t, set)
}

func TestSet_Iterator_NoSuchElement(t *testing.T) {
	set := gollections.NewSet[string]()
	it := set.Iterator()
	assert.PanicsWithValue(t, "has no more elements to iterate", func() {
		it.Next()
	})
}

func TestSet_Iterator_RemoveWithoutCallNext(t *testing.T) {
	set := gollections.NewSet[string]()
	it := set.Iterator()
	assert.PanicsWithValue(t, "next method has not been called", func() {
		it.Remove()
	})
}

func TestSet_Iterator_RemoveCalledTwice(t *testing.T) {
	set := gollections.NewSet("X")
	it := set.Iterator()
	it.Next()
	it.Remove()
	assert.PanicsWithValue(t, "remove method has already been called after the last call to the next method", func() {
		it.Remove()
	})
}

func TestSet_Remove(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, set, "A", "B", "C", "D", "E")

	set.Remove("C")
	assertSet(t, set, "A", "B", "D", "E")

	set.Remove("B", "D", "E")
	assertSet(t, set, "A")

	set.Remove("A")
	assertSet(t, set)
}

func TestSet_RemoveAll(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, set, "A", "B", "C", "D", "E")

	set.RemoveAll(gollections.NewSet("A", "C", "E"))
	assertSet(t, set, "B", "D")

	set.RemoveAll(gollections.NewSet("A", "B", "C"))
	assertSet(t, set, "D")

	set.RemoveAll(gollections.NewSet("D", "E"))
	assertSet(t, set)
}

func TestSet_RemoveIf(t *testing.T) {
	set := gollections.NewSet(1, 2, 3, 4, 5, 6, 7, 8)
	assertSet(t, set, 1, 2, 3, 4, 5, 6, 7, 8)

	set.RemoveIf(func(element int) bool {
		return element%2 != 0
	})
	assertSet(t, set, 2, 4, 6, 8)

	set.RemoveIf(func(element int) bool {
		return element%4 == 0
	})
	assertSet(t, set, 2, 6)
}

func TestSet_Rentains(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, set, "A", "B", "C", "D", "E")

	set.Retains("A", "B", "C")
	assertSet(t, set, "A", "B", "C")

	set.Retains("B")
	assertSet(t, set, "B")
}

func TestSet_RentainsAll(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertSet(t, set, "A", "B", "C", "D", "E")

	set.RetainsAll(gollections.NewSet("A", "B", "C"))
	assertSet(t, set, "A", "B", "C")

	set.RetainsAll(gollections.NewSet[string]())
	assertSet(t, set)
}

func TestSet_RentainsIf(t *testing.T) {
	set := gollections.NewSet(1, 2, 3, 4, 5, 6, 7, 8)
	assertSet(t, set, 1, 2, 3, 4, 5, 6, 7, 8)

	set.RetainsIf(func(element int) bool {
		return element%2 == 0
	})
	assertSet(t, set, 2, 4, 6, 8)

	set.RetainsIf(func(element int) bool {
		return element%4 != 0
	})
	assertSet(t, set, 2, 6)
}

func TestSet_Size(t *testing.T) {
	set := gollections.NewSet[string]()
	assert.Zero(t, set.Size())

	set.Add("A")
	assert.Equal(t, 1, set.Size())

	set.Add("B")
	assert.Equal(t, 2, set.Size())

	set.Add("C", "D")
	assert.Equal(t, 4, set.Size())

	set.Remove("C")
	assert.Equal(t, 3, set.Size())

	set.Remove("D", "A", "B")
	assert.Zero(t, set.Size())
}

func TestSet_ToArray(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	set := gollections.NewSet(elements...)
	assert.ElementsMatch(t, elements, set.ToArray())
}

func assertSet[T comparable](t *testing.T, set gollections.Set[T], elements ...T) {
	if len(elements) == 0 {
		assert.Zero(t, set.Size())
		assert.True(t, set.IsEmpty())
	} else {
		assert.ElementsMatch(t, elements, set.ToArray())
		assert.Equal(t, len(elements), set.Size())
		assert.False(t, set.IsEmpty())
	}
}
