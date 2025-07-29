package main

import "fmt"

// DemonstrateLoops shows the different kinds of loops available in Go.
func DemonstrateLoops() {
	fmt.Println("--- Loop Examples ---")
	// Basic for loop with a counter
	// This loop will iterate 5 times
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// For loop as a "while" loop
	// This loop will continue as long as the condition is true
	fmt.Println("\nFor loop as a 'while' loop:")
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println("Sum:", sum)

	// For-range loop to iterate over an array or slice
	// This is the most common and idiomatic way to loop over collections in Go
	fmt.Println("\nFor-range loop with a slice:")
	fruits := []string{"apple", "banana", "cherry"}
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Fruit: %s\n", index, fruit)
	}

	// For-range loop to iterate over a map
	fmt.Println("\nFor-range loop with a map:")
	person := map[string]string{"name": "Alice", "age": "30"}
	for key, value := range person {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// Infinite loop with a break statement
	// This loop will run forever until the break statement is executed
	fmt.Println("\nInfinite loop with a break:")
	count := 0
	for {
		fmt.Println("Infinite loop iteration:", count)
		count++
		if count == 3 {
			break // Exit the loop
		}
	}
	fmt.Println()
}
