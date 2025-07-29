package main

import (
	"fmt"
	// No need to import other files from the same 'main' package.
	// Go compiles them together automatically.
)

// The main function is the entry point of the program.
// It runs after all the init() functions in the package have completed.
func main() {
	fmt.Println("--- Go Basics Explorer ---")

	// Calling functions from other files in the same package.
	// This makes the execution order explicit and easier to follow.
	DemonstrateArrays()
	DemonstrateSlices()
	DemonstrateStrings()
	DemonstrateNumbers()
	DemonstrateBooleans()
	DemonstrateConditions()
	DemonstrateSwitchCases()
	DemonstrateLoops()
	DemonstrateBreakContinue()

	fmt.Println("\n--- All examples finished ---")
}
