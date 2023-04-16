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

package collections_test

import (
	"testing"

	"github.com/RageCage64/collections-go"
	"github.com/RageCage64/go-assert"
)

func runNewSetTestCase[T comparable](t *testing.T, elements ...T) {
	t.Helper()

	compareSetAndSlice(t, collections.NewSet(elements...), elements)
}

func TestNewSet(t *testing.T) {
	runNewSetTestCase(t, []int{1, 2, 3}...)
}

func runNewSetSaveDuplicatesTestCase[T comparable](
	t *testing.T,
	slice []T,
	expectedSet collections.Set[T],
	expectedDupes []T,
) {
	t.Helper()

	newSet, dupes := collections.NewSetSaveDuplicates(slice)
	assert.Assert(
		t,
		expectedSet.Equals(newSet),
		"new set does not equal expectation:\nexpected: %v\ngot: %v",
		expectedSet,
		newSet,
	)
	assert.SliceEqual(t, expectedDupes, dupes)
}

func TestNewSetSaveDuplicates(t *testing.T) {
	t.Run("no duplicates", func(t *testing.T) {
		t.Parallel()
		runNewSetSaveDuplicatesTestCase(
			t,
			[]int{1, 2, 3},
			collections.Set[int]{
				1: {},
				2: {},
				3: {},
			},
			[]int{},
		)
	})

	t.Run("all duplicates", func(t *testing.T) {
		t.Parallel()
		runNewSetSaveDuplicatesTestCase(
			t,
			[]int{1, 1, 2, 2, 3, 3},
			collections.Set[int]{
				1: {},
				2: {},
				3: {},
			},
			[]int{1, 2, 3},
		)
	})

	t.Run("all elements the same", func(t *testing.T) {
		t.Parallel()
		runNewSetSaveDuplicatesTestCase(
			t,
			[]int{1, 1, 1, 1},
			collections.Set[int]{
				1: {},
			},
			[]int{1, 1, 1},
		)
	})
}

func runSetToSliceTestCase[T comparable](t *testing.T, set collections.Set[T], expected []T) {
	assert.SliceEqual(t, expected, set.ToSlice())
}

func TestSetToSlice(t *testing.T) {
	runSetToSliceTestCase(
		t,
		collections.Set[int]{
			1: {},
			2: {},
			3: {},
		},
		[]int{1, 2, 3},
	)
}

func compareSetAndSlice[T comparable](t *testing.T, set collections.Set[T], slice []T) {
	t.Helper()

	for _, el := range slice {
		if !set.Contains(el) {
			t.Fatalf("expected set to contain %v", el)
		}
	}
	for el := range set {
		found := false
		for _, sliceEl := range slice {
			if el == sliceEl {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("set contained unexpected element %v", el)
		}
	}
}
