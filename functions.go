package gollections

// Consumer performs this operation on the given argument.
type Consumer[T any] func(t T)

// Predicate Evaluates this predicate on the given argument.
type Predicate[T any] func(t T) bool

// Comparator compares its two arguments for order. Returns a negative integer, zero, or a positive integer
// as the first argument is less than, equal to, or greater than the second.
type Comparator[T any] func(t1 T, t2 T) int

// UnaryOperator returns a unary operator that always returns its input argument.
type UnaryOperator[T any] func(t T) T
