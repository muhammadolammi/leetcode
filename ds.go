package main

import "log"

func MakeStack[T any]() *Stack[T] {
	return &Stack[T]{
		Data: []T{},
	}
}

func (q *Stack[T]) Push(v T) {
	q.Data = append(q.Data, v)
}
func (q *Stack[T]) Pop() (T, bool) {
	size := q.Size()
	if size == 0 {
		log.Println("empty stack, returning zero value.")
		var zero T
		return zero, false

	}
	elem := q.Data[size-1]
	// delete the element
	q.Data = q.Data[:size-1]
	return elem, true
}

func (q *Stack[T]) Peek() T {
	size := q.Size()
	if size == 0 {
		log.Println("empty stack, returning zero value.")
		var zero T
		return zero

	}
	return q.Data[size-1]

}

func (q *Stack[T]) Size() int {
	return len(q.Data)
}

func MakeQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Data: []T{},
	}
}

func (q *Queue[T]) Push(v T) {
	q.Data = append(q.Data, v)
}
func (q *Queue[T]) Pop() (T, bool) {
	size := q.Size()
	if size == 0 {
		log.Println("empty queue, returning zero value.")
		var zero T
		return zero, false
	}
	elem := q.Data[0]
	// delete the element
	q.Data = q.Data[1:]
	return elem, true
}

func (q *Queue[T]) Peek() (T, bool) {
	size := q.Size()
	if size == 0 {
		log.Println("empty stack, returning zero value.")
		var zero T
		return zero, false
	}
	return q.Data[0], true

}

func (q *Queue[T]) Size() int {
	return len(q.Data)
}
