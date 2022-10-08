package stack

import (
	"fmt"
	"reflect"
)

// Stack is a simple implementation of a stack data structure.
type Stack[T any] struct {
	items []T
}

// New returns a new stack.
func New[T any]() Stack[T] {
	return Stack[T]{items: []T{}}
}

// Push adds items to a stack.
func (s *Stack[T]) Push(item ...T) {
	s.items = append(s.items, item...)
}

// Pop removes an item from the stack and returns it.
// Panics if the stack is empty. Use PopSafe for safer access to the Stack.
func (s *Stack[T]) Pop() T {
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// PopSafe removes an item from the stack and returns it.
// Returns the zero value of the type if the stack is empty.
// To make this function safe, it uses reflection and is therefore slower than Pop.
func (s *Stack[T]) PopSafe() T {
	if s.IsEmpty() {
		return reflect.Zero(reflect.TypeOf(*new(T))).Interface().(T)
	}

	return s.Pop()
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the size of the stack.
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Peek returns the top item of the stack without removing it.
func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

// Clear removes all items from the stack.
func (s *Stack[T]) Clear() {
	s.items = []T{}
}

// Contains returns true if the stack contains the item.
func (s *Stack[T]) Contains(item T) bool {
	for _, v := range s.items {
		if reflect.DeepEqual(v, item) {
			return true
		}
	}
	return false
}

// Values returns the values of the stack as a slice.
func (s *Stack[T]) Values() []T {
	return s.items
}

func (s Stack[T]) String() string {
	return fmt.Sprint(s.items)
}
