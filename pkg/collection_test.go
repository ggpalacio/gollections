package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
)

func assertCollection[T comparable](t *testing.T, collection gollections.Collection[T], elements ...T) {
	if len(elements) == 0 {
		assert.Zero(t, collection.Size())
		assert.True(t, collection.IsEmpty())
	} else {
		assert.ElementsMatch(t, elements, collection.ToArray())
		assert.Equal(t, len(elements), collection.Size())
		assert.False(t, collection.IsEmpty())
	}
}

func TestIterator(t *testing.T) {
	elements := []string{"A", "B", "C", "D", "E"}
	set := gollections.NewSet(elements...)

	for it := set.Iterator(); it.HasNext(); {
		assert.Contains(t, elements, it.Next())
	}
	assertCollection(t, set, "A", "B", "C", "D", "E")

	for it := set.Iterator(); it.HasNext(); {
		it.Next()
		it.Remove()
	}
	assertCollection(t, set)
}

func TestIterator_NoSuchElement(t *testing.T) {
	set := gollections.NewSet[string]()
	it := set.Iterator()
	assert.PanicsWithValue(t, "has no more elements to iterate", func() {
		it.Next()
	})
}

func TestIterator_RemoveWithoutCallNext(t *testing.T) {
	set := gollections.NewSet[string]()
	it := set.Iterator()
	assert.PanicsWithValue(t, "next method has not been called", func() {
		it.Remove()
	})
}

func TestIterator_RemoveCalledTwice(t *testing.T) {
	set := gollections.NewSet("X")
	it := set.Iterator()
	it.Next()
	it.Remove()
	assert.PanicsWithValue(t, "remove method has already been called after the last call to the next method", func() {
		it.Remove()
	})
}
