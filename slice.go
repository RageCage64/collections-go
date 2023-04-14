package collections

/*****************
This section contains the alias API, which allows for receiver-style calling
convention.
*****************/

// Slice is an alias over a Go slice that enables a direct receiver-style API.
// For Slice, T must satisfy comparable.
type Slice[T comparable] []T

// Check if a slice contains an element. Calls SliceContains.
func (sl Slice[T]) Contains(needle T) bool {
	return SliceContains(sl, needle)
}

// Run an operator on every element in the slice, and return a slice that
// contains the result of every operation. Calls SliceMap.
func (sl Slice[T]) Map(op UnaryOperator[T]) Slice[T] {
	return SliceMap(sl, op)
}

// AnySlice is an alias over a Go slice that does not enforce T satisfying
// comparable. Any method that relies on comparison of elements is not
// available for AnySlice.
type AnySlice[T any] []T

// Run an operator on every element in the slice, and return a slice that
// contains the result of every operation. Calls SliceMap.
func (sl AnySlice[T]) Map(op UnaryOperator[T]) AnySlice[T] {
	return SliceMap(sl, op)
}

/*****************
This section contains the direct API, which is a more traditional Go style
API.
*****************/

// Check if a slice contains an element.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// Allocations: None.
func SliceContains[T comparable](haystack []T, needle T) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

// Run an operator on every element in the slice, and return a slice that
// contains the result of every operation.
//
// Time Complexity: O(n * m) (where m = complexity of operator)
// Space Complexity: O(n)
// Allocations: 1 slice, n elements.
func SliceMap[T any](sl []T, op UnaryOperator[T]) []T {
	result := make([]T, len(sl))
	for i := 0; i < len(sl); i++ {
		result[i] = op(sl[i])
	}
	return result
}

// Run a predicate on every element in the slice, and return a slice
// that contains every element for which the predicate was true.
//
// Time Complexity: O(n * m) (where m = complexity of predicate)
// Space Complexity: O(n)
// Allocations: 2 slice, first with n elements, second is the first slice
// resized if necessary.
func SliceFilter[T any](sl []T, pred UnaryPredicate[T]) []T {
	result := make([]T, len(sl))
	resultIdx := 0
	for i := 0; i < len(sl); i++ {
		if pred(sl[i]) {
			result[resultIdx] = sl[i]
			resultIdx++
		}
	}
	if resultIdx < len(result)-1 {
		result = result[:resultIdx]
	}
	return result
}
