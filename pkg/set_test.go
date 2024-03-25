package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	set := gollections.NewSet[string]()
	assertCollection(t, set)

	set.Add("A")
	assertCollection(t, set, "A")

	set.Add("B")
	assertCollection(t, set, "A", "B")

	set.Add("C")
	assertCollection(t, set, "A", "B", "C")

	set.Add("A")
	assertCollection(t, set, "A", "B", "C")

	set.Add("B", "C", "D")
	assertCollection(t, set, "A", "B", "C", "D")

	set.Add("E", "F")
	assertCollection(t, set, "A", "B", "C", "D", "E", "F")
}

func TestSet_AddAll(t *testing.T) {
	set1 := gollections.NewSet[string]()
	set2 := gollections.NewSet[string]()
	assertCollection(t, set1)
	assertCollection(t, set2)

	set1.Add("A", "B", "C")
	assertCollection(t, set1, "A", "B", "C")
	assertCollection(t, set2)

	set2.AddAll(set1)
	assertCollection(t, set1, "A", "B", "C")
	assertCollection(t, set2, "A", "B", "C")

	set1.Clear()
	assertCollection(t, set1)
	assertCollection(t, set2, "A", "B", "C")
}

func TestSet_Clear(t *testing.T) {
	set := gollections.NewSet[string]()
	assertCollection(t, set)

	set.Add("A", "B", "C")
	assertCollection(t, set, "A", "B", "C")

	set.Clear()
	assertCollection(t, set)
}

func TestSet_Contains(t *testing.T) {
	set := gollections.NewSet("A", "B", "C")
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
	TestIterator(t)
}

func TestSet_Remove(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertCollection(t, set, "A", "B", "C", "D", "E")

	set.Remove("C")
	assertCollection(t, set, "A", "B", "D", "E")

	set.Remove("B", "D", "E")
	assertCollection(t, set, "A")

	set.Remove("A")
	assertCollection(t, set)
}

func TestSet_RemoveAll(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertCollection(t, set, "A", "B", "C", "D", "E")

	set.RemoveAll(gollections.NewSet("A", "C", "E"))
	assertCollection(t, set, "B", "D")

	set.RemoveAll(gollections.NewSet("A", "B", "C"))
	assertCollection(t, set, "D")

	set.RemoveAll(gollections.NewSet("D", "E"))
	assertCollection(t, set)
}

func TestSet_RemoveIf(t *testing.T) {
	set := gollections.NewSet(1, 2, 3, 4, 5, 6, 7, 8)
	assertCollection(t, set, 1, 2, 3, 4, 5, 6, 7, 8)

	set.RemoveIf(func(element int) bool {
		return element%2 != 0
	})
	assertCollection(t, set, 2, 4, 6, 8)

	set.RemoveIf(func(element int) bool {
		return element%4 == 0
	})
	assertCollection(t, set, 2, 6)
}

func TestSet_Rentains(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertCollection(t, set, "A", "B", "C", "D", "E")

	set.Retains("A", "B", "C")
	assertCollection(t, set, "A", "B", "C")

	set.Retains("B")
	assertCollection(t, set, "B")
}

func TestSet_RentainsAll(t *testing.T) {
	set := gollections.NewSet("A", "B", "C", "D", "E")
	assertCollection(t, set, "A", "B", "C", "D", "E")

	set.RetainsAll(gollections.NewSet("A", "B", "C"))
	assertCollection(t, set, "A", "B", "C")

	set.RetainsAll(gollections.NewSet[string]())
	assertCollection(t, set)
}

func TestSet_RentainsIf(t *testing.T) {
	set := gollections.NewSet(1, 2, 3, 4, 5, 6, 7, 8)
	assertCollection(t, set, 1, 2, 3, 4, 5, 6, 7, 8)

	set.RetainsIf(func(element int) bool {
		return element%2 == 0
	})
	assertCollection(t, set, 2, 4, 6, 8)

	set.RetainsIf(func(element int) bool {
		return element%4 != 0
	})
	assertCollection(t, set, 2, 6)
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

func TestSet_ToBag(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	set := gollections.NewSet(elements...)
	assert.Equal(t, gollections.NewBag(elements...), set.ToBag())
}
