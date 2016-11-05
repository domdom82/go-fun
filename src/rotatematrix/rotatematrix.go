package main

import (
	"fmt"
	"math/rand"
	"os"
)

/*
Playground app to rotate a random square matrix by 90 degrees clockwise
*/
func main() {

	var size = enterTheMatrix()

	var matrix = matrixReloaded(size)
	fmt.Println()
	fmt.Println("The Matrix Reloaded:")
	printMatrix(matrix)

	matrixRevolutions(matrix)
	fmt.Println()
	fmt.Println("Matrix Revolutions:")
	printMatrix(matrix)
}

// reads the size of the matrix from the user
func enterTheMatrix() uint {
	size := uint(0)
	fmt.Printf("Enter the Matrix: ")
	var _, err = fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		fmt.Println("Expected positive integer for matrix size.")
		os.Exit(2)
	}
	return size
}

// generates a random square matrix from given size
func matrixReloaded(size uint) [][]int {
	var matrix = make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	//fill the matrix with random numbers
	for y := range matrix {
		for x := range matrix[y] {
			matrix[y][x] = rand.Intn(100)
		}
	}
	return matrix
}

// rotates the given matrix in-place 90 degrees clockwise
func matrixRevolutions(matrix [][]int) {
	/*
			since the matrix is square, we can work our way inward shifting elements top-left, top-right, bottom-right, bottom-left
			then resume with the next inner layer:

			Original   Layer 0    Layer 1
			1 2 3 4	   4 3 2 1    4 3 2 1
		    2 3 4 5 => 5	 2 => 5 4 3 2
			3 4 5 6    6     3    6 5 4 3
			4 5 6 7    7 6 5 4    7 6 5 4

			The example shows that only size / 2 iterations are necessary for both x and y axis.
	*/

	size := len(matrix)

	for y := 0; y < size/2; y++ {
		for x := 0; x < size/2; x++ {
			//top-left moves to the right as x increases
			topLeft := matrix[y][x]
			//top-right moves to the bottom as x increases
			topRight := matrix[x][size-1-y]
			//bottom-right moves to the left as x increases
			bottomRight := matrix[size-1-y][size-1-x]
			//bottom-left moves to the top as x increases
			bottomLeft := matrix[size-1-x][y]

			//top-right -> top-left
			matrix[x][size-1-y] = topLeft
			//bottom-right -> top-right
			matrix[size-1-y][size-1-x] = topRight
			//bottom-left -> bottom-right
			matrix[size-1-x][y] = bottomRight
			//top-left -> bottom-left
			matrix[y][x] = bottomLeft
		}
	}
}

// prints the given matrix
func printMatrix(matrix [][]int) {
	for y := range matrix {
		for x := range matrix[y] {
			fmt.Printf("%2d ", matrix[y][x])
		}
		fmt.Println()
	}
}
