package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

// Playground for testing channels and goroutines for parallel programming

// We use this type to communicate search ranges
type SearchRange struct {
	lowerBound int
	upperBound int
}

func main() {

	// The idea of this app is demonstrate asynchronous parallel programming.
	// We do this by elaborating a bit on the "primefinder" example and distribute the
	// search for prime numbers to multiple threads.

	fmt.Printf("\nPRIMEFINDER - MultiCore Version\n\n")

	// Find out the number of CPUs we have to allocate work
	numCPUs := runtime.NumCPU()
	fmt.Printf("Found %d CPUs to work with.\n", numCPUs)

	// Create a channel to communicate search range per thread
	searchRanges := make(chan SearchRange)
	defer close(searchRanges)

	// Create a channel to communicate results
	primeNumbers := make(chan int)
	defer close(primeNumbers)

	// First we as the user to give us a range in which to search for prime numbers
	lowerBound := 0
	upperBound := 0
	fmt.Printf("Please enter lower bound for search: ")
	fmt.Scan(&lowerBound)
	fmt.Println()
	fmt.Printf("Please enter upper bound for search: ")
	fmt.Scan(&upperBound)
	fmt.Println()

	if upperBound < lowerBound {
		log.Fatal("Upper bound must be greater than lower bound.")
	}

	fmt.Println("Waiting for prime numbers...")
	start := time.Now()

	// We now subdivide the range into equal chunks per CPU
	searchSize := upperBound - lowerBound
	if numCPUs > 1 && searchSize%2 != 0 {
		// normalize range. we will only ever have an even number of CPUs (unless we have only one of course)
		searchSize++
	}

	chunkSize := searchSize / numCPUs
	rest := searchSize % numCPUs

	for i := 0; i < numCPUs; i++ {
		// start thread (will wait for incoming ranges)
		searchPrime(searchRanges, primeNumbers)
		// feed new range into channel
		var nextRange SearchRange
		nextRange.lowerBound = lowerBound + chunkSize*i
		nextRange.upperBound = nextRange.lowerBound + chunkSize
		if i == numCPUs-1 {
			// the last CPU gets the missing upper bound in case we can't divide evenly
			nextRange.upperBound += rest
		}
		searchRanges <- nextRange
	}

	// now read from results channel until we found all prime numbers
	for searchDone := 0; searchDone < numCPUs; {
		nextPrime := <-primeNumbers
		if nextPrime == -1 {
			searchDone++
		} else {
			fmt.Printf("%d ", nextPrime)
		}
	}

	finish := time.Now()
	duration := finish.Sub(start)

	fmt.Println()
	fmt.Printf("Found all prime numbers. Took %s\n", duration.String())
}

func searchPrime(ranges chan SearchRange, results chan int) {
	go func() {
		// read range from channel
		searchRange := <-ranges
		// search for prime numbers in this range
		for i := searchRange.lowerBound; i < searchRange.upperBound; i++ {
			if isPrime(i) {
				// write new prime to results channel
				results <- i
			}
		}
		// signal finish
		results <- -1
	}()
}

func isPrime(n int) bool {
	// easy cases
	if n < 2 {
		return false
	}
	if n%2 == 0 {
		return false
	}

	// hard search
	for i := 3; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
