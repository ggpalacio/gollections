package gollections

type Comparator[T any] func(value1, value2 T) int
