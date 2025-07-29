package main

import "fmt"

// DemonstrateSlices shows various ways to create and manipulate slices in Go.
func DemonstrateSlices() {
	fmt.Println("--- Slice Examples ---")
	// Creating a slice using the make function
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	// This creates a slice of strings with a length of 0 and a capacity of 5.
	s := make([]string, 0, 5)
	fmt.Println("Empty slice:", s)
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))

	// Appending elements to a slice
	// The append function returns a new slice containing the new elements.
	s = append(s, "Apple")
	s = append(s, "Banana", "Cherry")
	fmt.Println("\nSlice after appending:", s)
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))

	// Creating a slice literal
	// This creates a slice with initial values.
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("\nSlice literal:", numbers)

	// Slicing a slice to create a new slice
	// This creates a new slice that references a portion of the original slice's underlying array.
	// The new slice will include elements from index 1 up to (but not including) index 4.
	sliceOfNumbers := numbers[1:4]
	fmt.Println("Slice of numbers [1:4]:", sliceOfNumbers)

	// Modifying the new slice will affect the original slice
	sliceOfNumbers[0] = 99
	fmt.Println("Original numbers slice after modification:", numbers)

	// Copying a slice
	// To create an independent copy, use the copy() function.
	copiedSlice := make([]int, len(numbers))
	copy(copiedSlice, numbers)
	fmt.Println("\nCopied slice:", copiedSlice)

	// Modifying the copied slice does not affect the original
	copiedSlice[0] = 1000
	fmt.Println("Original numbers slice after modifying copy:", numbers)
	fmt.Println("Copied slice after modification:", copiedSlice)
	fmt.Println()
}
