package main

import "fmt"

// DemonstrateArrays shows how to declare, initialize, and use arrays.
func DemonstrateArrays() {
	fmt.Println("--- Array Examples ---")
	// Declare an array of strings with a length of 3
	// This is an example of an array in Go
	var fruits [3]string

	// Assign values to the array elements
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Orange"

	// Access and print the array elements
	fmt.Println("Fruits:", fruits)
	fmt.Println("First fruit:", fruits[0])

	// Declare and initialize an array with values
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)

	// Get the length of an array
	fmt.Println("Length of numbers array:", len(numbers))
	fmt.Println()
}
