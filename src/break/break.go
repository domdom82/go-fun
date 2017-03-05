package main

import "fmt"

// break and continue with labels
func main() {

myLabel:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("(%d, %d)\n", i, j)
			if i > 6 {
				fmt.Println("break")
				break myLabel // this will END the outer loop
			}
			if j > 6 {
				fmt.Println("continue")
				continue myLabel // this will END the inner loop
			}
		}
	}

}
