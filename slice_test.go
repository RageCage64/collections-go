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

func TestSliceContains(t *testing.T) {
	haystack := []int{1, 2, 3}
	needle := 1
	assert.Assert(
		t,
		collections.SliceContains(haystack, needle),
		"expected %v to contain %d", haystack, needle,
	)
}

func TestSliceContainsEach(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8}
	needles := []int{5, 3, 4, 6}
	assert.Assert(
		t,
		collections.SliceContainsEach(haystack, needles),
		"expected %v to contain each of %v", haystack, needles,
	)
}

func TestSliceSubset(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8}
	subset := []int{3, 4, 5}
	assert.Assert(
		t,
		collections.SliceSubset(haystack, subset),
		"expected %v to contain subset %v", haystack, subset,
	)
}

func TestSliceSubsetEqualSlices(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8}
	subset := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Assert(
		t,
		collections.SliceSubset(haystack, subset),
		"expected %v to contain subset %v", haystack, subset,
	)
}

func TestSliceSubsetOutOfOrderFails(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8}
	subset := []int{4, 3, 5}
	assert.Assert(
		t,
		!collections.SliceSubset(haystack, subset),
		"expected %v not to contain out of order subset %v", haystack, subset,
	)
}

func TestSliceSubsetStrict(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8}
	subset := []int{5, 6, 7}
	assert.Assert(
		t,
		collections.SliceSubsetStrict(haystack, subset),
		"expected %v to contain strict subset %v", haystack, subset,
	)
}

func TestSliceSubsetStrictEqualSetFails(t *testing.T) {
	haystack := []int{1, 2, 3, 4, 5, 6, 7, 8}
	subset := []int{1, 2, 3, 4, 5, 6, 7, 8}
	assert.Assert(
		t,
		!collections.SliceSubsetStrict(haystack, subset),
		"expected %v not to have strict subset %v", haystack, subset,
	)
}

func TestMap(t *testing.T) {
	sl := collections.Slice[int]{1, 2, 3}
	result := sl.Map(
		func(x int) int {
			return x + 1
		},
	)
	assert.SliceEqual(t, []int{2, 3, 4}, result)
}

func TestFilter(t *testing.T) {
	sl := collections.Slice[int]{1, 2, 3, 4}
	result := sl.Filter(
		func(x int) bool {
			return x%2 == 0
		},
	)
	assert.SliceEqual(t, []int{2, 4}, result)
}

func TestReduce(t *testing.T) {
	sl := collections.Slice[int]{1, 2, 3, 4, 5}
	result := collections.SliceReduce(sl, 0,
		func(accumulator int, current int) int {
			return accumulator + current
		},
	)
	assert.Equal(t, 15, result)
}

func TestForEach(t *testing.T) {
	result := 0
	collections.Range(5).ForEach(func(x int) {
		result++
	})
	assert.Equal(t, 5, result)
}
