package gollections

type Predicate[T any] func(value T) bool

type Comparator[T any] func(value1, value2 T) int
