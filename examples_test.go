package stack_test

import (
	"fmt"

	"atomicgo.dev/stack"
)

func ExampleNew() {
	stack.New[string]()
}

func ExampleStack_Push() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s)

	// Output:
	// [Hello World]
}

func ExampleStack_Pop() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	// Output:
	// World
	// Hello
}

func ExampleStack_PopSafe() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.PopSafe())
	fmt.Println(s.PopSafe())
	fmt.Println(s.PopSafe())

	// Output:
	// World
	// Hello
	//
}

func ExampleStack_Values() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.Values())

	// Output:
	// [Hello World]
}

func ExampleStack_Size() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.Size())

	// Output:
	// 2
}

func ExampleStack_IsEmpty() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.IsEmpty())

	s.Clear()

	fmt.Println(s.IsEmpty())

	// Output:
	// false
	// true
}

func ExampleStack_Contains() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.Contains("Hello"))
	fmt.Println(s.Contains("Foo"))

	// Output:
	// true
	// false
}

func ExampleStack_Clear() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	s.Clear()

	fmt.Println(s)

	// Output:
	// []
}

func ExampleStack_String() {
	s := stack.New[string]()
	s.Push("Hello")
	s.Push("World")

	fmt.Println(s.String())

	// Output:
	// [Hello World]
}
