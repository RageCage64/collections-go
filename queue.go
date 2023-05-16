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

import "errors"

type LLNode[T any] struct {
	Value T
	Prev  *LLNode[T]
	Next  *LLNode[T]
}

func (head *LLNode[T]) Len() int {
	curr := head
	n := 0
	for curr != nil {
		n++
		curr = curr.Next
	}
	return n
}

func (head *LLNode[T]) Range() []T {
	llLen := head.Len()
	result := make([]T, llLen)
	curr := head
	for i := 0; i < llLen; i++ {
		result[i] = curr.Value
		curr = curr.Next
	}
	return result
}

func (head *LLNode[T]) AddToTail(val T) *LLNode[T] {
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	node := &LLNode[T]{Value: val, Prev: curr}
	curr.Next = node
	return node
}

func (head *LLNode[T]) AddToHead(val T) *LLNode[T] {
	node := &LLNode[T]{Value: val, Next: head}
	head.Prev = node
	return node
}

func (n *LLNode[T]) RemoveSelf() {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
}

// A FIFO queue.
type Queue[T any] struct {
	Head *LLNode[T]
}

func (q *Queue[T]) Enqueue(val T) {
	if q.Head == nil {
		q.Head = &LLNode[T]{Value: val}
	} else {
		q.Head.AddToTail(val)
	}
}

var ErrQueueEmpty = errors.New("collections: the queue is empty")

func (q *Queue[T]) Dequeue() (T, error) {
	if q.Head == nil {
		return *new(T), ErrQueueEmpty
	}
	result := q.Head
	q.Head = result.Next
	return result.Value, nil
}

// A Stack, aka FILO queue.
type Stack[T any] struct {
	Head *LLNode[T]
}

func (s *Stack[T]) Push(val T) {
	if s.Head == nil {
		s.Head = &LLNode[T]{Value: val}
	} else {
		node := s.Head.AddToHead(val)
		s.Head = node
	}
}

var ErrStackEmpty = errors.New("the stack is empty")

func (s *Stack[T]) Pop() (T, error) {
	if s.Head == nil {
		return *new(T), ErrStackEmpty
	}
	oldHead := s.Head
	s.Head = s.Head.Next
	oldHead.RemoveSelf()
	return oldHead.Value, nil
}
