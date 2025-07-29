package main

import "fmt"

// DemonstrateNumbers shows how to work with integer and floating-point numbers.
func DemonstrateNumbers() {
	fmt.Println("--- Number Examples ---")

	// --- Integers (Whole Numbers) ---
	// Go offers integers of various sizes (e.g., int8, int16, int32, int64).
	// The 'int' type is the most common; its size is platform-dependent (32 or 64 bits).

	// Declare an integer using the 'var' keyword.
	var a int = 10
	// Declare another integer using short variable declaration (:=).
	b := 20

	// Perform a simple arithmetic operation.
	sum := a + b
	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)

	// --- Floating-Point Numbers (Decimals) ---
	// Go has two floating-point types: float32 and float64.
	// 'float64' is the default and generally recommended for better precision.
	var pi float64 = 3.14159
	radius := 5.5 // Go infers this as a float64.

	// Perform a floating-point calculation.
	area := pi * radius * radius
	fmt.Printf("The area of a circle with radius %.2f is approximately %.2f\n", radius, area)

	// --- Type Conversion ---
	// Go requires explicit type conversion; it does not automatically convert between types.
	var integerValue int = 100
	var floatValue float64 = float64(integerValue) // Explicitly convert the int to a float64.
	fmt.Printf("Integer %d converted to float: %.2f\n", integerValue, floatValue)

	fmt.Println() // Add a newline for better output formatting.
}
