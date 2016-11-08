package main

import (
	"fmt"
	"stack/stack"
)

// Playground app for the custom stack type
func main() {

	//make a stack
	var myStack stack.Stack

	//push some stuff on the stack (even another stack)
	myStack.Push("blubb")
	myStack.Push(42)
	myStack.Push(false)
	var otherStack stack.Stack
	otherStack.Push("this is on the other stack")
	otherStack.Push("this also")
	myStack.Push(otherStack)

	fmt.Println("myStack is", myStack)

	//get stack size
	fmt.Println("Stack size is", myStack.Len())

	//get stack capacity
	fmt.Println("Stack capacity is", myStack.Cap())

	//pop some stuff
	myStack.Pop()
	myStack.Pop()

	//we expect 42 now
	num, _ := myStack.Pop()

	fmt.Println("This should be 42 ==", num)

	//empty? I think not
	fmt.Println("Stack empty:", myStack.IsEmpty())

	//provoke an error by popping two more
	myStack.Pop()
	_, err := myStack.Pop()

	if err != nil {
		fmt.Println("Got expected error:", err)
	} else {
		fmt.Println("w00t?")
	}

	//empty now? I hope so...
	fmt.Println("Stack finally empty:", myStack.IsEmpty())

}
