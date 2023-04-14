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
