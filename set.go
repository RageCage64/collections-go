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

// Set is a type alias over a map to an empty struct. This allows a nice
// API over the hash map functionality already available, where you can instead
// view the set as a single dimension collection and save a lot of clutter of
// empty struct initialization any time you add to the set.
type Set[T comparable] map[T]struct{}

// Initialize a new Set with a set of arguments. Good for if you want to initialize
// a new set in place with values.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations: 1 slice, n elements (variadic function argument). 1 set, n elements.
func NewSet[T comparable](elements ...T) Set[T] {
	set := make(Set[T], len(elements))
	for i := 0; i < len(elements); i++ {
		set.Add(elements[i])
	}
	return set
}

// Initialize a new Set with a slice. Good for if you already have a slice and want a
// set with its values (saves an allocation over NewSet). This function will toss
// duplicate values, see NewSetSaveDuplicates if you need to retain duplicate info.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations: 1 slice, n elements (variadic function argument). 1 set, n elements.
func NewSetFromSlice[T comparable](sl []T) Set[T] {
	set := make(Set[T], len(sl))
	for i := 0; i < len(sl); i++ {
		set.Add(sl[i])
	}
	return set
}

// Initialize a new set while preserving the elements that were found to be duplicate.
// The duplicate slice will be nil if there are no duplicates to save on allocations.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations (best case): 1 set, n elements
// Allocations (worst case): 1 set, n/2 elements, 2 slices, one with n-1 space, the
// second is the first one resized if necessary.
func NewSetSaveDuplicates[T comparable](slice []T) (Set[T], []T) {
	set := make(Set[T], len(slice))
	var dupes []T = nil
	dupesIdx := 0
	for _, el := range slice {
		if set.Contains(el) {
			if dupes == nil {
				// Now that we know there is at least one dupe, we can make the allocation.
				// The max possible size of the duplicate array is n-1 (aka a slice where
				// every element is the same).
				dupes = make([]T, len(slice)-1)
			}
			dupes[dupesIdx] = el
			dupesIdx++
		} else {
			set[el] = struct{}{}
		}
	}
	if dupes != nil && dupesIdx < len(dupes)-1 {
		dupes = dupes[:dupesIdx]
	}
	return set, dupes
}

// ToSlice will take all of the elements of the set and create a new slice out of them.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations: 1 slice, n elements
func (set Set[T]) ToSlice() []T {
	slice := make([]T, len(set))
	i := 0
	for el := range set {
		slice[i] = el
		i++
	}
	return slice
}

// Add an element to the set.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
// Allocations: Set resize
func (s Set[T]) Add(el T) {
	s[el] = struct{}{}
}

// Remove an element from the set.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
// Allocations: Set resize
func (s Set[T]) Remove(el T) {
	delete(s, el)
}

// Check if the set contains a particular value.
//
// Time Complexity: O(1)
// Space Complexity: O(1)
// Allocations: None
func (s Set[T]) Contains(el T) bool {
	_, containsEl := s[el]
	return containsEl
}

// Creates a Clone of the set which contains all the same values.
//
// Time Complexity: O(n)
// Space Complexity: O(n)
// Allocations: 1 set of n elements
func (s Set[T]) Clone() Set[T] {
	clone := make(Set[T], len(s))
	for el := range s {
		clone.Add(el)
	}
	return clone
}

// Checks if the set is equal to another.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
// Allocations: None
func (s Set[T]) Equals(s2 Set[T]) bool {
	if len(s) != len(s2) {
		return false
	}
	for el := range s {
		if !s2.Contains(el) {
			return false
		}
	}
	return true
}
