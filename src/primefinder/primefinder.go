package main

import (
	"fmt"
	"os"
)

/*
Playground app to find prime numbers between 0 and n
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: primefinder <n>")
	}
}
