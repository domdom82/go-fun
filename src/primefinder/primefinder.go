package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
Playground app to find all prime numbers between n and m
*/
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: primefinder <n> <m>")
		os.Exit(1)
	}
	var n, err = strconv.Atoi(os.Args[1])
	var m, err2 = strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Println("Expected positive integer for n.")
		os.Exit(2)
	}

	if err2 != nil {
		fmt.Printf("%s\n", err2.Error())
		fmt.Println("Expected positive integer for m.")
		os.Exit(2)
	}

	if n > m {
		fmt.Println("<m> must be bigger than <n>.")
		os.Exit(3)
	}

	un := uint(n)
	um := uint(m)

	for i := un; i <= um; i++ {
		start := time.Now()
		p := isPrime(i)
		finish := time.Now()
		duration := finish.Sub(start)
		if p {
			fmt.Printf("Found prime number: %d. Took %s\n", i, duration.String())
		}
	}

}

/*
Returns true if given n is prime
*/
func isPrime(n uint) bool {
	// easy cases
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}

	// hard search
	for i := uint(3); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
