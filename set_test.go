package collections_test

import (
	"testing"

	"github.com/RageCage64/collections-go"
	"github.com/RageCage64/go-assert"
)

func TestSetToSlice(t *testing.T) {
	set := collections.Set[int]{
		1: {},
		2: {},
		3: {},
	}
	slice := set.ToSlice()
	assert.SliceEqual(t, []int{1, 2, 3}, slice)
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
