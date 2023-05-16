// Copyright (c) 2023 Braydon Kains
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package collections

/*****************
The arguments for higher-order function algorithms.
*****************/

// UnaryPredicate is a function that takes a single element of type
// T and returns a boolean. Used for filtering operations.
type UnaryPredicate[T any] func(T) bool

// UnaryOperator is a function that takes a single element of type
// T and returns an element of type T.
type UnaryOperator[T any] func(T) T

// UnaryReceiver is a function that takes a single element of type
// T and uses it to do something without returning anything.
type UnaryReceiver[T any] func(T)

// Reducer is a function that takes an accumulator and the current
// element of the slice.
type Reducer[Acc any, T any] func(accumulator Acc, current T) Acc

/*****************
The alias API, which allows for receiver-style calling convention.
*****************/

// Slice is an alias over a Go slice that enables a direct receiver-style API.
// For Slice, T must satisfy comparable.
type Slice[T comparable] []T

// Check if a slice contains an element. Calls SliceContains.
func (sl Slice[T]) Contains(needle T) bool {
	return SliceContains(sl, needle)
}

// Check if a slice contains each individual element from another slice. Order
// is not considered. To consider order, use Subset. Calls SliceContainsEach.
func (sl Slice[T]) ContainsEach(needles []T) bool {
	return SliceContainsEach(sl, needles)
}

// Check if another slice is a (non-strict) subset of the slice. Calls SliceSubset.
func (sl Slice[T]) Subset(sub []T) bool {
	return SliceSubset(sl, sub)
}

// Run an operator on every element in the slice, and return a slice that
// contains the result of every operation. Calls SliceMap.
func (sl Slice[T]) Map(op UnaryOperator[T]) Slice[T] {
	return SliceMap(sl, op)
}

// Run a predicate on every element in the slice, and return a slice
// that contains every element for which the predicate was true. Calls
// SliceFilter.
func (sl Slice[T]) Filter(pred UnaryPredicate[T]) Slice[T] {
	return SliceFilter(sl, pred)
}

// Run an operation using every element of the slice one at a time.
func (sl Slice[T]) ForEach(do UnaryReceiver[T]) {
	SliceForEach(sl, do)
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

// Run a predicate on every element in the slice, and return a slice
// that contains every element for which the predicate was true. Calls
// SliceFilter.
func (sl AnySlice[T]) Filter(pred UnaryPredicate[T]) AnySlice[T] {
	return SliceFilter(sl, pred)
}

// Run an operation using every element of the slice one at a time.
func (sl AnySlice[T]) ForEach(do UnaryReceiver[T]) {
	SliceForEach(sl, do)
}

/*****************
This section contains the direct API, which is a more traditional Go style
API.
*****************/

// Produces an array with n elements from 0 to n-1. Good for producing a slice
// to do things n times. Produces a Slice[int], great for calling into ForEach.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations: Slice of n elements
func Range(n int) Slice[int] {
	sl := make(Slice[int], n)
	for i := 0; i < n; i++ {
		sl[i] = i
	}
	return sl
}

// Check if a slice contains an element.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// Allocations: None
func SliceContains[T comparable](haystack []T, needle T) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

// Check if a slice contains each individual element of another slice. Order
// is not considered. To consider order, use SliceSubset.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations: 1 map, m (needles argument) elements
func SliceContainsEach[T comparable](haystack []T, needles []T) bool {
	// Allocating a map here is better for time complexity and allocations,
	// since the alternative is working with a slice that would need to be
	// searched through and in most cases resized when elements are found.
	needleSet := make(map[T]struct{}, len(needles))
	for i := 0; i < len(needles); i++ {
		needleSet[needles[i]] = struct{}{}
	}

	for i := 0; i < len(haystack); i++ {
		delete(needleSet, haystack[i])
	}
	return len(needleSet) == 0
}

// Check if another slice is a subset of the slice. This check is non-strict.
// For a strict subset, use SliceSubsetStrict.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// Allocations: None
func SliceSubset[T comparable](sl []T, sub []T) bool {
	if len(sub) > len(sl) {
		return false
	}
	return subset(sl, sub)
}

// Check if another slice is a strict subset of the slice. That is, the other
// slice is a subset but not equal to the main slice.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// Allocations: None
func SliceSubsetStrict[T comparable](sl []T, sub []T) bool {
	// Strict subset means the two slices can't be the same size.
	if len(sub) > len(sl)-1 {
		return false
	}
	return subset(sl, sub)
}

func subset[T comparable](sl []T, sub []T) bool {
	subIdx := 0
	for i := 0; i < len(sl); i++ {
		if sl[i] == sub[subIdx] {
			if subIdx == len(sub)-1 {
				return true
			}
			subIdx++
		} else {
			subIdx = 0
		}
	}
	return false
}

// Run an operator on every element in the slice, and return a slice that
// contains the result of every operation.
//
// Sometimes known by other names: Transform, Select
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
// Sometimes known by other names: Where
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

// With a starting accumulator, run the reducer operator with the accumulator
// and each element of the slice. The accumulator should be a slice or a pointer
// to a value so the change is reflected throughout the combination.
//
// Sometimes known by other names: Fold, Aggregate, Combine
//
// Time Complexity: O(n * m) (where m = complexity of reducer)
// Space Complexity: O(1)
// Allocations: None
func SliceReduce[Acc any, T any](
	sl []T,
	accumulator Acc,
	reducer Reducer[Acc, T],
) Acc {
	for i := 0; i < len(sl); i++ {
		accumulator = reducer(accumulator, sl[i])
	}
	return accumulator
}

// Run an operation using each element of the slice one at a time.
//
// Time Complexity: O(n * m) (where m is the complexity of the operation)
// Space Complexity: O(1)
// Allocations: None
func SliceForEach[T any](sl []T, do UnaryReceiver[T]) {
	for i := 0; i < len(sl); i++ {
		do(sl[i])
	}
}
