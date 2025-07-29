package main

import "fmt"

// DemonstrateConditions shows how to use if-else and switch statements.
func DemonstrateConditions() {
	fmt.Println("--- Condition Examples ---")
	// Example of an if-else if-else statement
	// This block checks the value of the 'age' variable
	age := 18

	if age < 13 {
		fmt.Println("You are a child.")
	} else if age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are an adult.")
	}

	// Example of a switch statement
	// This block checks the value of the 'day' variable
	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("It's the start of the week.")
	case "Friday":
		fmt.Println("It's the end of the week.")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}
	fmt.Println()
}
