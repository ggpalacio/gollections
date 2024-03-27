package gollections_test

import (
	"testing"

	gollections "github.com/ggpalacio/gollections/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var numbers = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

type CollectionTestSuite struct {
	suite.Suite
	collection gollections.Collection[int]
}

func TestSet(t *testing.T) {
	suite.Run(t, &CollectionTestSuite{
		collection: gollections.NewSet[int](),
	})
}

func TestBag(t *testing.T) {
	suite.Run(t, &CollectionTestSuite{
		collection: gollections.NewBag[int](),
	})
}

func (ref *CollectionTestSuite) SetupTest() {
	ref.collection.Clear()
}

func (ref *CollectionTestSuite) TestAdd() {
	for index, number := range numbers {
		assertCollection(ref.T(), ref.collection, numbers[:index]...)

		ref.collection.Add(number)
		assertCollection(ref.T(), ref.collection, numbers[:index+1]...)
	}
}

func (ref *CollectionTestSuite) TestAddAll() {
	assertCollection(ref.T(), ref.collection)

	ref.collection.AddAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection, numbers[:]...)
}

func (ref *CollectionTestSuite) TestClear() {
	assertCollection(ref.T(), ref.collection)

	ref.collection.Add(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	ref.collection.Clear()
	assertCollection(ref.T(), ref.collection)
}

func (ref *CollectionTestSuite) TestContains() {
	for _, number := range numbers {
		ref.False(ref.collection.Contains(number))

		ref.collection.Add(number)
		ref.True(ref.collection.Contains(number))
	}
}

func (ref *CollectionTestSuite) TestContainsAll() {
	for index, number := range numbers {
		expectedCollection := newCollection(numbers[0 : index+1]...)
		ref.False(ref.collection.ContainsAll(expectedCollection))

		ref.collection.Add(number)
		ref.True(ref.collection.ContainsAll(expectedCollection))
	}
}

func (ref *CollectionTestSuite) TestContainsAny() {
	for index, number := range numbers {
		expectedCollection := newCollection(numbers[index])
		ref.False(ref.collection.ContainsAny(expectedCollection))

		ref.collection.Add(number)
		ref.True(ref.collection.ContainsAny(expectedCollection))
	}
}

func (ref *CollectionTestSuite) TestIsEmpty() {
	ref.True(ref.collection.IsEmpty())

	for _, number := range numbers {
		ref.collection.Add(number)
		ref.False(ref.collection.IsEmpty())
	}

	for _, number := range numbers {
		ref.False(ref.collection.IsEmpty())
		ref.collection.Remove(number)
	}

	ref.True(ref.collection.IsEmpty())
}

func (ref *CollectionTestSuite) TestIterator() {
	ref.collection.AddAll(newCollection(numbers[:]...))

	for it := ref.collection.Iterator(); it.HasNext(); {
		ref.Contains(numbers, it.Next())
	}
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	for it := ref.collection.Iterator(); it.HasNext(); {
		it.Next()
		it.Remove()
	}
	assertCollection(ref.T(), ref.collection)
}

func (ref *CollectionTestSuite) TestIterator_NoSuchElement() {
	it := ref.collection.Iterator()
	assert.PanicsWithValue(ref.T(), "has no more elements to iterate", func() {
		it.Next()
	})
}

func (ref *CollectionTestSuite) TestIterator_RemoveWithoutCallNext() {
	it := ref.collection.Iterator()
	assert.PanicsWithValue(ref.T(), "next method has not been called", func() {
		it.Remove()
	})
}

func (ref *CollectionTestSuite) TestIterator_RemoveCalledTwice() {
	ref.collection.Add(numbers[0])

	it := ref.collection.Iterator()
	it.Next()
	it.Remove()
	assert.PanicsWithValue(ref.T(), "remove method has already been called after the last call to the next method", func() {
		it.Remove()
	})
}

func (ref *CollectionTestSuite) TestRemove() {
	ref.collection.AddAll(newCollection(numbers[:]...))

	for index, number := range numbers {
		assertCollection(ref.T(), ref.collection, numbers[index:]...)
		ref.collection.Remove(number)
		assertCollection(ref.T(), ref.collection, numbers[index+1:]...)
	}
}

func (ref *CollectionTestSuite) TestRemoveAll() {
	ref.collection.AddAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	ref.collection.RemoveAll(newCollection(1, 3, 5, 7, 9))
	assertCollection(ref.T(), ref.collection, 0, 2, 4, 6, 8)

	ref.collection.RemoveAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection)
}

func (ref *CollectionTestSuite) TestRemoveIf() {
	ref.collection.AddAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	ref.collection.RemoveIf(func(value int) bool {
		return value%2 != 0
	})
	assertCollection(ref.T(), ref.collection, 0, 2, 4, 6, 8)

	ref.collection.RemoveIf(func(value int) bool {
		return value%4 == 0
	})
	assertCollection(ref.T(), ref.collection, 2, 6)
}

func (ref *CollectionTestSuite) TestRetains() {
	ref.collection.AddAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	ref.collection.Retains(1, 3, 5, 7, 9)
	assertCollection(ref.T(), ref.collection, 1, 3, 5, 7, 9)

	ref.collection.Retains(0, 2, 4, 6, 8)
	assertCollection(ref.T(), ref.collection)
}

func (ref *CollectionTestSuite) TestRetainsAll() {
	ref.collection.AddAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	ref.collection.RetainsAll(newCollection(0, 1, 2, 3, 4))
	assertCollection(ref.T(), ref.collection, 0, 1, 2, 3, 4)

	ref.collection.RetainsAll(newCollection(5, 6, 7, 8, 9))
	assertCollection(ref.T(), ref.collection)
}

func (ref *CollectionTestSuite) TestRetainsIf() {
	ref.collection.AddAll(newCollection(numbers[:]...))
	assertCollection(ref.T(), ref.collection, numbers[:]...)

	ref.collection.RetainsIf(func(value int) bool {
		return value%2 == 0
	})
	assertCollection(ref.T(), ref.collection, 0, 2, 4, 6, 8)

	ref.collection.RetainsIf(func(value int) bool {
		return value%4 == 0
	})
	assertCollection(ref.T(), ref.collection, 0, 4, 8)
}

func (ref *CollectionTestSuite) TestSize() {
	for index, number := range numbers {
		ref.collection.Add(number)
		ref.Equal(index+1, ref.collection.Size())
	}
	for index, number := range numbers {
		ref.Equal(len(numbers)-index, ref.collection.Size())
		ref.collection.Remove(number)
	}
	ref.Zero(ref.collection.Size())
}

func (ref *CollectionTestSuite) TestToArray() {
	ref.collection.AddAll(newCollection(numbers[:]...))
	ref.ElementsMatch(numbers, ref.collection.ToArray())
}

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

func newCollection[T comparable](elements ...T) gollections.Collection[T] {
	return gollections.NewSet(elements...)
}
