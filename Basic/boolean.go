package main // Declare the package as main

import "fmt"

// DemonstrateBooleans shows how to use boolean types and logical operators.
func DemonstrateBooleans() {
	fmt.Println("--- Boolean Examples ---")

	// Declare a boolean variable named 'isGoFun' and initialize it to true.
	// The 'bool' keyword specifies the type.
	var isGoFun bool = true
	fmt.Println("Is Go fun?", isGoFun) // Print the value of the boolean variable.

	// Declare another boolean variable using short variable declaration (:=).
	// The type 'bool' is inferred from the value 'false'.
	isLearning := true
	// The '!' operator negates the boolean value (true becomes false, false becomes true).
	fmt.Println("Are we still learning?", isLearning)

	// Booleans are commonly used in control structures like if-else statements.
	if isGoFun {
		// This block executes if isGoFun is true.
		fmt.Println("Indeed, Go is fun!")
	} else {
		// This block executes if isGoFun is false.
		fmt.Println("Let's make it fun!")
	}

	// --- Logical Operations ---
	// Declare two boolean variables for logical operations.
	var hasCoffee bool = true
	var isTired bool = false

	// The '&&' (AND) operator returns true only if both operands are true.
	canCode := hasCoffee && !isTired // true AND !false -> true AND true -> true
	fmt.Println("Ready to code?", canCode)

	// The '||' (OR) operator returns true if at least one of the operands is true.
	needsBreak := !hasCoffee || isTired // true OR false -> false OR false -> false
	fmt.Println("Time for a break?", needsBreak)
	fmt.Println() // Add a newline for better output formatting.
}
