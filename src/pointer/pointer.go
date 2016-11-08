package main

import (
	"fmt"
)

// Playground for pointers. Helps understand * and & operands
func main() {

	//this is an int
	myInt := 3

	//this is a pointer to an int
	var myIntPointer *int

	//what is the value of myInt? (duh)
	fmt.Println("myInt =", myInt)

	//what is the value of the reference to myInt? (should be mem address)
	fmt.Println("&myInt =", &myInt)

	//what is the value of myIntPointer? (should be nil)
	fmt.Println("myIntPointer =", myIntPointer)

	//what is the value of the reference to myIntPointer? (should be mem address)
	fmt.Println("&myIntPointer =", &myIntPointer)

	//now let myIntPointer point to myInt
	myIntPointer = &myInt

	//what is the value if we dereference myIntPointer? (should be 3)
	fmt.Println("*myIntPointer =", *myIntPointer)

	//what is the value of myIntPointer? (should be address of myInt)
	fmt.Println("myIntPointer =", myIntPointer)

	//now make a pointer to a pointer to an int
	var myPointerIntPointer **int

	//what is the value of myPointerIntPointer? (should be nil)
	fmt.Println("myPointerIntPointer =", myPointerIntPointer)

	//now let myPointerIntPointer point to myIntPointer
	myPointerIntPointer = &myIntPointer

	//what is the value if we dereference myPointerIntPointer? (should be address of myInt)
	fmt.Println("*myPointerIntPointer =", *myPointerIntPointer)

	//what is the value if we dereference myPointerIntPointer twice? (should be 3 again)
	fmt.Println("**myPointerIntPointer =", **myPointerIntPointer)

	//what is the value if we dereference and then reference myPointerIntPointer? (should be address of myIntPointer)
	fmt.Println("&*myPointerIntPointer =", &*myPointerIntPointer)

	//what is the value if we dereference myPointerIntPointer twice and then reference? (should be address of myInt)
	fmt.Println("&**myPointerIntPointer =", &**myPointerIntPointer)
}
