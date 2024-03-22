package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	set := gollections.NewSet[string]()
	assert.Zero(t, set.Size())
	assert.True(t, set.IsEmpty())

	set.Add("A")
	assertSet(t, set, "A", 1)

	set.Add("B")
	assertSet(t, set, "B", 2)

	set.Add("C")
	assertSet(t, set, "C", 3)

	set.Add("A")
	assertSet(t, set, "A", 3)

	set.Add("B", "C", "D", "E")
	assertSet(t, set, "B", 5)
	assertSet(t, set, "C", 5)
	assertSet(t, set, "D", 5)
	assertSet(t, set, "E", 5)
}

func TestSet_Clear(t *testing.T) {
	set := gollections.NewSet[string]()
	assert.Zero(t, set.Size())
	assert.True(t, set.IsEmpty())

	set.Add("A", "B", "C")
	assertSet(t, set, "A", 3)
	assertSet(t, set, "B", 3)
	assertSet(t, set, "C", 3)

	set.Clear()
	assert.Zero(t, set.Size())
	assert.True(t, set.IsEmpty())
}

func assertSet[T comparable](t *testing.T, set gollections.Set[T], item T, size int) {
	assert.True(t, set.Contains(item))
	assert.Equal(t, set.Size(), size)
	assert.False(t, set.IsEmpty())
}
