package stack

import (
	"errors"
)

// Playground for defining new types. Exercise creates a stack type.

// Stack is a slice of anything, which oddly enough is something that implements the empty interface.
// could be read as void*
type Stack []interface{}

// this is how Go adds functions to types. by putting a "receiver" in front which essentially becomes "this" in the function.

// Push with a *stack receiver because we want to modify the stack
func (stack *Stack) Push(stuff interface{}) {
	*stack = append(*stack, stuff)
}

// Pop the last guy
func (stack *Stack) Pop() (stuff interface{}, err error) {
	theStack := *stack

	if len(theStack) == 0 {
		return nil, errors.New("Tried to pop an empty stack.")
	}

	//get last element
	last := theStack[len(theStack)-1]
	//reduce stack by 1
	*stack = theStack[:len(theStack)-1]

	return last, nil
}

// some convenience methods

// Len returns the size of the stack
func (stack *Stack) Len() int {
	return len(*stack)
}

// Cap returns the (current) capacity of the stack
func (stack *Stack) Cap() int {
	return cap(*stack)
}

// IsEmpty returns true if the stack is empty
func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}
