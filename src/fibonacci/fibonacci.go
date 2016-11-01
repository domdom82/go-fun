package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
Playground app to find all fibonacci numbers between 0 and n
*/
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: fibonacci <n>")
		os.Exit(1)
	}
	var n, err = strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Println("Expected positive integer for n.")
		os.Exit(2)
	}

	un := uint(n)
	for i := uint(0); i <= un; i++ {
		start := time.Now()
		f := fibonacci(uint(i))
		finish := time.Now()
		duration := finish.Sub(start)
		fmt.Printf("F(%d) = %d. Took %s\n", i, f, duration.String())
	}
}

/*
Returns fibonacci number for given n by recursion
*/
func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
