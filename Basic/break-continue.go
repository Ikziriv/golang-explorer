package main

import "fmt"

// DemonstrateLoopControl provides detailed examples of 'break' and 'continue' statements in Go.
func DemonstrateBreakContinue() {
	fmt.Println("\n--- Detailed Loop Control (break/continue) Examples ---")

	// --- Example 1: Basic 'break' ---
	// The 'break' statement terminates the execution of the innermost for, switch, or select statement.
	fmt.Println("\n1. Basic 'break':")
	fmt.Println("Looping from 0 to 9, but will break at 5.")
	for i := 0; i < 10; i++ {
		if i == 5 {
			// When i is 5, the 'break' statement is executed.
			// This immediately terminates the for loop.
			fmt.Println("Breaking the loop!")
			break
		}
		// This line will not be printed for i = 5 or greater.
		fmt.Printf("Current number: %d\n", i)
	}
	fmt.Println("Loop finished.")

	// --- Example 2: Basic 'continue' ---
	// The 'continue' statement begins the next iteration of the innermost for loop.
	// It skips the rest of the loop body for the current iteration.
	fmt.Println("\n2. Basic 'continue':")
	fmt.Println("Looping from 0 to 9, but will skip even numbers.")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			// If i is an even number, 'continue' is executed.
			// This skips the fmt.Printf statement below and starts the next iteration (i++).
			continue
		}
		// This line only executes for odd numbers.
		fmt.Printf("Processing odd number: %d\n", i)
	}
	fmt.Println("Loop finished.")

	// --- Example 3: 'break' with a Label ---
	// A labeled break can be used to terminate an outer loop from within a nested loop.
	fmt.Println("\n3. 'break' with a Label:")
outerloop: // This is the label for the outer loop.
	for i := 0; i < 3; i++ {
		fmt.Printf("Outer loop (i=%d)\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("  Inner loop (j=%d)\n", j)
			if i == 1 && j == 1 {
				// This 'break' statement refers to the 'outerloop' label.
				// It will terminate both the inner and the outer loop.
				fmt.Println("  Breaking out of the outer loop from the inner loop!")
				break outerloop
			}
		}
	}
	fmt.Println("Finished after labeled break.")
	fmt.Println()
}
