package main

import "fmt"

// Playground for understanding the difference between slices and arrays
func main() {
	fmt.Printf("SLICE:\n\n")
	// Let's make a slice of animals
	var animals = []string{"cat", "dog", "mouse"}
	fmt.Printf("before: Address of animals=%p\n", &animals)
	fmt.Println("animals=", animals)

	// First difference: Slices are not Arrays. If we pass a slice by value and modify the contents of the copy, the
	// changes affect the original slice, even though we passed by value. This anomaly is rooted in the fact that only the slice is passed by value,
	// however its underlying array can be changed and is not copied.
	replace(animals)

	fmt.Printf("after: Address of animals=%p\n", &animals)
	fmt.Println("animals=", animals)

	// Now what happens if we pass a slice by reference instead?
	// We would expect the same behavior, except that the address of the passed slice will never change.
	replaceByRef(&animals)

	fmt.Printf("after: Address of animals=%p\n", &animals)
	fmt.Println("animals=", animals)

	// Now let's make an array, shall we?
	fmt.Printf("ARRAY:\n\n")

	var people = [...]string{"alice", "bob", "chuck"}
	fmt.Printf("before: Address of people=%p\n", &people)
	fmt.Println("people=", people)

	// Second difference: If we pass an array by value and modify the contents of the copy, the
	// changes do NOT affect the original array. This is how we would expect "passing by value" to work usually.
	replaceArr(people)

	// Notice how people did not change here.
	fmt.Printf("after: Address of people=%p\n", &people)
	fmt.Println("people=", people)

	// Now what happens if we pass an array by reference instead?
	// We should now see a change to the outside array.
	replaceArrByRef(&people)

	// Notice how people now did change.
	fmt.Printf("after: Address of people=%p\n", &people)
	fmt.Println("people=", people)
}

func replace(a []string) {
	// The address of a is going to be different from the address of animals
	fmt.Printf("replace: Address of a=%p\n", &a)
	a[2] = "house"
}

func replaceByRef(a *[]string) {
	// The address of a is going to be the same as the address of animals
	fmt.Printf("replaceByRef: Address of *a=%p\n", &*a)
	(*a)[2] = "blouse"
}

// [3]string and [4]string are two different TYPES in Go. Therefore we can not pass a generic []string - which would clash with generic slices.
// This makes using pure arrays basically pointless. Go doc says we should always use slices instead. People be crazy...
// The only usage for this would be with arrays that never change, e.g. vector or matrix algebra.
func replaceArr(a [3]string) {
	// The address of a is going to be different from the address of people
	fmt.Printf("replaceArr: Address of a=%p\n", &a)
	a[2] = "cathy"
}

func replaceArrByRef(a *[3]string) {
	// The address of a is going to be the same as the address of people
	fmt.Printf("replaceArrByRef: Address of *a=%p\n", &*a)
	(*a)[2] = "curtis"
	a[2] = "curtis" //this equal to the line above as a convenience. Go does smart dereferencing of pointers.
}
