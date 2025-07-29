package main

import (
	"fmt"
	"math"
)

func DemonstrateMath() {
	// Basic Arithmetic Operations
	a := 10
	b := 3
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println("Addition (a + b):", a+b)
	fmt.Println("Subtraction (a - b):", a-b)
	fmt.Println("Multiplication (a * b):", a*b)
	// Note: Integer division truncates the result
	fmt.Println("Integer Division (a / b):", a/b)
	// Modulo operator gives the remainder
	fmt.Println("Modulo (a %% b):", a%b)

	// Floating-point operations
	c := 10.0
	d := 3.0
	fmt.Printf("\nc = %.1f, d = %.1f\n", c, d)
	fmt.Println("Floating-point Division (c / d):", c/d)

	// Using the math package for more advanced operations
	fmt.Println("\n--- Using the 'math' package ---")

	// Square root
	fmt.Println("Square root of 25:", math.Sqrt(25))

	// Power
	fmt.Println("2 to the power of 8:", math.Pow(2, 8))

	// Ceiling and Floor
	// Ceil rounds up, Floor rounds down
	x := 4.7
	fmt.Printf("\nx = %.1f\n", x)
	fmt.Println("Ceiling of x:", math.Ceil(x))
	fmt.Println("Floor of x:", math.Floor(x))

	// Round
	// Rounds to the nearest integer
	y := 4.5
	fmt.Printf("\ny = %.1f\n", y)
	fmt.Println("Round of x:", math.Round(x))
	fmt.Println("Round of y:", math.Round(y))

	// Max and Min
	// Returns the larger or smaller of two numbers
	fmt.Println("\nMax of 10 and 20:", math.Max(10, 20))
	fmt.Println("Min of 10 and 20:", math.Min(10, 20))

	// Constants from the math package
	fmt.Println("\nValue of Pi:", math.Pi)
	fmt.Println("Value of E:", math.E)
}
