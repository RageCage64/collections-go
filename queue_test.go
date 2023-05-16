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

func TestLinkedList(t *testing.T) {
	head := &collections.LLNode[int]{Value: 1}
	head.AddToTail(2)
	head.AddToTail(3)
	llSlice := head.Range()
	assert.SliceEqual(t, []int{1, 2, 3}, llSlice)

	head = head.AddToHead(0)
	llSlice = head.Range()
	assert.SliceEqual(t, []int{0, 1, 2, 3}, llSlice)

	head.Next.RemoveSelf()
	llSlice = head.Range()
	assert.SliceEqual(t, []int{0, 2, 3}, llSlice)
}

func TestQueue(t *testing.T) {
	q := &collections.Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	assert.Equal(t, 3, q.Head.Len())

	for i := 1; i <= 3; i++ {
		x, err := q.Dequeue()
		assert.NilErr(t, err)
		assert.Equal(t, i, x)
	}
}

func TestStack(t *testing.T) {
	q := &collections.Stack[int]{}
	q.Push(1)
	q.Push(2)
	q.Push(3)
	assert.Equal(t, 3, q.Head.Len())

	for i := 3; i >= 1; i-- {
		x, err := q.Pop()
		assert.NilErr(t, err)
		assert.Equal(t, i, x)
	}
}
