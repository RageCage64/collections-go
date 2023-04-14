package collections

// A UnaryPredicate is a function that takes a single element of type
// T and returns a boolean.
type UnaryPredicate[T any] func(T) bool

// A UnaryOperator is a function that takes a single element of type
// T and returns an element of type T.
type UnaryOperator[T any] func(T) T
