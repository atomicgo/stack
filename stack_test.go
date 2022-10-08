package stack_test

import (
	"fmt"
	"testing"

	"atomicgo.dev/stack"
	"github.com/MarvinJWendt/testza"
)

func TestStack(t *testing.T) {
	s := stack.New[string]()

	t.Run("Push", func(t *testing.T) {
		s.Push("Hello")
		s.Push("World")
	})

	t.Run("Values", func(t *testing.T) {
		testza.AssertEqual(t, s.Values(), []string{"Hello", "World"})
	})

	t.Run("String", func(t *testing.T) {
		testza.AssertEqual(t, s.String(), fmt.Sprintf("%v", []string{"Hello", "World"}))
	})

	t.Run("Size", func(t *testing.T) {
		testza.AssertEqual(t, s.Size(), 2)
	})

	t.Run("IsNotEmpty", func(t *testing.T) {
		testza.AssertFalse(t, s.IsEmpty())
	})

	t.Run("Contains", func(t *testing.T) {
		testza.AssertTrue(t, s.Contains("Hello"))
		testza.AssertTrue(t, s.Contains("World"))
		testza.AssertFalse(t, s.Contains("Foo"))
	})

	t.Run("Pop", func(t *testing.T) {
		testza.AssertEqual(t, "World", s.Pop())
		testza.AssertEqual(t, "Hello", s.Pop())
	})

	t.Run("PopFails", func(t *testing.T) {
		testza.AssertPanics(t, func() {
			s.Pop()
		})
	})

	t.Run("IsEmpty", func(t *testing.T) {
		testza.AssertTrue(t, s.IsEmpty())
	})

	t.Run("PopSafe", func(t *testing.T) {
		s.Push("Hello")
		testza.AssertEqual(t, "Hello", s.PopSafe())
		testza.AssertEqual(t, "", s.PopSafe())
	})

	t.Run("Clear", func(t *testing.T) {
		s.Push("Hello")
		s.Push("World")
		s.Clear()
		testza.AssertTrue(t, s.IsEmpty())
	})

	t.Run("Peek", func(t *testing.T) {
		s.Push("Hello")
		s.Push("World")
		testza.AssertEqual(t, "World", s.Peek())
		testza.AssertEqual(t, 2, s.Size())
		s.Clear()
	})

	t.Run("PushMultiple", func(t *testing.T) {
		s.Push("Hello", "World", "How", "Are", "You", "?")
		testza.AssertEqual(t, 6, s.Size())
		s.Clear()
	})

}
