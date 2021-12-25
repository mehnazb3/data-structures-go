// Stack
// Works on LIFO
// Pushing an element
// Pop an element

package main

import (
	"fmt"
)

type stack []int

// push - adds element in the last
func (s *stack) push(data int) int {
	*s = append(*s, data)
	return data
}

// pop - removes element from the last
func (s *stack) pop() {
	*s = (*s)[:(len(*s) - 1)]
	return
}

func main() {
	var myStack stack
	// Adding items
	myStack.push(10)                  // 10
	myStack.push(15)                  // 10, 15
	myStack.push(25)                  // 10, 15, 25
	fmt.Println("My stack:", myStack) // 10, 15, 25
	// Removing items
	myStack.pop()
	fmt.Println("My stack:", myStack) // 10, 15
	myStack.pop()
	fmt.Println("My stack:", myStack) // 10
	myStack.pop()
	fmt.Println("My stack:", myStack) // []
}
