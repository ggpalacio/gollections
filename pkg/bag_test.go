package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
)

func TestBag_Add(t *testing.T) {
	bag := gollections.NewBag[string]()
	assertCollection(t, bag)

	bag.Add("A")
	assertCollection(t, bag, "A")

	bag.Add("B", "C")
	assertCollection(t, bag, "A", "B", "C")

	bag.Add("A", "B", "C", "D")
	assertCollection(t, bag, "A", "A", "B", "B", "C", "C", "D")

	bag.Add("D")
	assertCollection(t, bag, "A", "A", "B", "B", "C", "C", "D", "D")
}

func TestBag_AddAll(t *testing.T) {
	bag := gollections.NewBag[string]()
	assertCollection(t, bag)

	bag.AddAll(gollections.NewBag("A", "B", "C"))
	assertCollection(t, bag, "A", "B", "C")

	bag.AddAll(gollections.NewBag("A", "D"))
	assertCollection(t, bag, "A", "A", "B", "C", "D")

	bag.AddAll(gollections.NewBag("A"))
	assertCollection(t, bag, "A", "A", "A", "B", "C", "D")
}

func TestBag_AddMany(t *testing.T) {
	bag := gollections.NewBag[string]()
	assertCollection(t, bag)

	bag.AddMany("A", 1)
	assertCollection(t, bag, "A")

	bag.AddMany("B", 2)
	assertCollection(t, bag, "A", "B", "B")

	bag.AddMany("C", 3)
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.AddMany("D", 0)
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.AddMany("E", -1)
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")
}

func TestBag_Clear(t *testing.T) {
	bag := gollections.NewBag("A", "B", "C")
	assertCollection(t, bag, "A", "B", "C")

	bag.Clear()
	assertCollection(t, bag)
}

func TestBag_Count(t *testing.T) {
	bag := gollections.NewBag("A", "B", "C")
	assert.Equal(t, 3, bag.Size())
	assert.Equal(t, 1, bag.Count("A"))
	assert.Equal(t, 1, bag.Count("B"))
	assert.Equal(t, 1, bag.Count("C"))

	bag.Add("B")
	assert.Equal(t, 4, bag.Size())
	assert.Equal(t, 1, bag.Count("A"))
	assert.Equal(t, 2, bag.Count("B"))
	assert.Equal(t, 1, bag.Count("C"))

	bag.AddMany("C", 2)
	assert.Equal(t, 6, bag.Size())
	assert.Equal(t, 1, bag.Count("A"))
	assert.Equal(t, 2, bag.Count("B"))
	assert.Equal(t, 3, bag.Count("C"))

	bag.Remove("A", "B", "C")
	assert.Equal(t, 3, bag.Size())
	assert.Equal(t, 0, bag.Count("A"))
	assert.Equal(t, 1, bag.Count("B"))
	assert.Equal(t, 2, bag.Count("C"))

	bag.RemoveEvery("C")
	assert.Equal(t, 1, bag.Size())
	assert.Equal(t, 0, bag.Count("A"))
	assert.Equal(t, 1, bag.Count("B"))
	assert.Equal(t, 0, bag.Count("C"))
}

func TestBag_Contains(t *testing.T) {
	bag := gollections.NewBag[string]()
	bag.Add("A", "B", "C")
	assert.True(t, bag.Contains("A"))
	assert.True(t, bag.Contains("B"))
	assert.True(t, bag.Contains("C"))
	assert.False(t, bag.Contains("D"))
	assert.False(t, bag.Contains("E"))
}

func TestBag_ContainsAll(t *testing.T) {
	bag := gollections.NewBag("A", "B", "C")
	assert.True(t, bag.ContainsAll(gollections.NewBag("A")))
	assert.True(t, bag.ContainsAll(gollections.NewBag("A", "B")))
	assert.True(t, bag.ContainsAll(gollections.NewBag("A", "B", "C")))
	assert.False(t, bag.ContainsAll(gollections.NewBag("A", "B", "C", "D")))
	assert.False(t, bag.ContainsAll(gollections.NewBag("A", "B", "C", "D", "E")))
}

func TestBag_ContainsAny(t *testing.T) {
	bag := gollections.NewBag("D", "E")
	assert.False(t, bag.ContainsAny(gollections.NewBag("A")))
	assert.False(t, bag.ContainsAny(gollections.NewBag("A", "B")))
	assert.False(t, bag.ContainsAny(gollections.NewBag("A", "B", "C")))
	assert.True(t, bag.ContainsAny(gollections.NewBag("A", "B", "C", "D")))
	assert.True(t, bag.ContainsAny(gollections.NewBag("A", "B", "C", "D", "E")))
}

func TestBag_IsEmpty(t *testing.T) {
	bag := gollections.NewBag[string]()
	assert.True(t, bag.IsEmpty())

	bag.Add("X")
	assert.False(t, bag.IsEmpty())

	bag.Remove("X")
	assert.True(t, bag.IsEmpty())
}

func TestBag_Iterator(t *testing.T) {
	elements := []string{"A", "B", "B", "C", "C", "C"}
	bag := gollections.NewBag(elements...)

	for it := bag.Iterator(); it.HasNext(); {
		assert.Contains(t, elements, it.Next())
	}
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	for it := bag.Iterator(); it.HasNext(); {
		it.Next()
		it.Remove()
	}
	assertCollection(t, bag)
}

func TestBag_Remove(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.Remove("A")
	assertCollection(t, bag, "B", "B", "C", "C", "C")

	bag.Remove("C", "B")
	assertCollection(t, bag, "B", "C", "C")

	bag.Remove("A", "B", "C")
	assertCollection(t, bag, "C")
}

func TestBag_RemoveAll(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.RemoveAll(gollections.NewBag("A", "C"))
	assertCollection(t, bag, "B", "B", "C", "C")

	bag.RemoveAll(gollections.NewBag("B", "B", "B"))
	assertCollection(t, bag, "C", "C")

	bag.RemoveAll(gollections.NewBag("C"))
	assertCollection(t, bag, "C")
}

func TestBag_RemoveEvery(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C", "D", "E")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C", "D", "E")

	bag.RemoveEvery("A")
	assertCollection(t, bag, "B", "B", "C", "C", "C", "D", "E")

	bag.RemoveEvery("B")
	assertCollection(t, bag, "C", "C", "C", "D", "E")

	bag.RemoveEvery("D", "E")
	assertCollection(t, bag, "C", "C", "C")

	bag.RemoveEvery("C")
	assertCollection(t, bag)
}

func TestBag_RemoveIf(t *testing.T) {
	bag := gollections.NewBag(1, 2, 3, 4, 5, 6, 7, 8)
	assertCollection(t, bag, 1, 2, 3, 4, 5, 6, 7, 8)

	bag.RemoveIf(func(element int) bool {
		return element%2 != 0
	})
	assertCollection(t, bag, 2, 4, 6, 8)

	bag.RemoveIf(func(element int) bool {
		return element%4 == 0
	})
	assertCollection(t, bag, 2, 6)
}

func TestBag_RemoveMany(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.RemoveMany("A", 3)
	assertCollection(t, bag, "B", "B", "C", "C", "C")

	bag.RemoveMany("B", 2)
	assertCollection(t, bag, "C", "C", "C")

	bag.RemoveMany("C", 1)
	assertCollection(t, bag, "C", "C")

	bag.RemoveMany("C", 0)
	assertCollection(t, bag, "C", "C")

	bag.RemoveMany("C", -1)
	assertCollection(t, bag, "C", "C")
}

func TestBag_Retains(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.Retains("A", "B", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.Retains("B")
	assertCollection(t, bag, "B", "B")

	bag.Retains("C")
	assertCollection(t, bag)
}

func TestBag_RetainsAll(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.RetainsAll(gollections.NewBag("A", "C"))
	assertCollection(t, bag, "A", "C", "C", "C")

	bag.RetainsAll(gollections.NewBag[string]())
	assertCollection(t, bag)
}

func TestBag_RetainsIf(t *testing.T) {
	bag := gollections.NewBag(1, 2, 2, 3, 3, 3, 4, 4, 4, 4)
	assertCollection(t, bag, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4)

	bag.RetainsIf(func(element int) bool {
		return element%2 == 0
	})
	assertCollection(t, bag, 2, 2, 4, 4, 4, 4)

	bag.RetainsIf(func(element int) bool {
		return element%4 != 0
	})
	assertCollection(t, bag, 2, 2)
}

func TestBag_Set(t *testing.T) {
	bag := gollections.NewBag("A", "B", "B", "C", "C", "C")
	assertCollection(t, bag, "A", "B", "B", "C", "C", "C")

	bag.Set("C", 1)
	assertCollection(t, bag, "A", "B", "B", "C")

	bag.Set("A", 3)
	assertCollection(t, bag, "A", "A", "A", "B", "B", "C")

	bag.Set("B", 0)
	assertCollection(t, bag, "A", "A", "A", "C")

	bag.Set("C", -1)
	assertCollection(t, bag, "A", "A", "A")
}

func TestBag_Size(t *testing.T) {
	bag := gollections.NewBag[string]()
	assert.Zero(t, bag.Size())

	bag.Add("A")
	assert.Equal(t, 1, bag.Size())

	bag.AddMany("B", 2)
	assert.Equal(t, 3, bag.Size())

	bag.Add("C", "C", "C")
	assert.Equal(t, 6, bag.Size())

	bag.Remove("B")
	assert.Equal(t, 5, bag.Size())

	bag.RemoveEvery("C")
	assert.Equal(t, 2, bag.Size())

	bag.Remove("A", "B")
	assert.Zero(t, bag.Size())
}

func TestBag_ToArray(t *testing.T) {
	elements := []string{"A", "B", "B", "C", "C", "C"}
	bag := gollections.NewBag(elements...)
	assert.ElementsMatch(t, elements, bag.ToArray())
}

func TestBag_ToSet(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	bag := gollections.NewBag(elements...)
	assert.Equal(t, gollections.NewSet(elements...), bag.ToSet())
}
