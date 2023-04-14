package collections_test

import (
	"testing"

	"github.com/RageCage64/collections-go"
	"github.com/RageCage64/go-assert"
)

func TestSliceContainsDirect(t *testing.T) {
	x := []int{1, 2, 3}
	assert.Assert(t, collections.SliceContains(x, 1), "it didn't")
}
