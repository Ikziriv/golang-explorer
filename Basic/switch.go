package main

import (
	"fmt"
	"time"
)

// getDayOfWeek returns the current day of the week as a string.
func getDayOfWeek() string {
	// The time.Now() function returns the current local time.
	// The Weekday() method returns the day of the week (e.g., time.Sunday).
	// The String() method converts the weekday to its string representation (e.g., "Sunday").
	return time.Now().Weekday().String()
}

// checkType is a helper function for the type switch example.
// It takes an interface{} which can hold a value of any type.
func checkType(i interface{}) {
	fmt.Printf("Value: %v, ", i)
	// A type switch is used to discover the dynamic type of an interface variable.
	// The 'i.(type)' syntax is only allowed inside a type switch.
	switch v := i.(type) {
	case int:
		// This case runs if the interface 'i' holds an int.
		// The variable 'v' will have the type int.
		fmt.Printf("Type: int. It's a whole number. Square is %d.\n", v*v)
	case string:
		// This case runs if the interface 'i' holds a string.
		// The variable 'v' will have the type string.
		fmt.Printf("Type: string. Length is %d.\n", len(v))
	case bool:
		// This case runs if the interface 'i' holds a bool.
		// The variable 'v' will have the type bool.
		fmt.Printf("Type: bool. The opposite is %t.\n", !v)
	default:
		// This case runs if none of the other cases match.
		// The variable 'v' will have the same type as 'i' (interface{}).
		fmt.Printf("Type: unknown. I don't know what to do with a %T.\n", v)
	}
}

// DemonstrateSwitchCases provides detailed examples of switch statements in Go.
func DemonstrateSwitchCases() {
	fmt.Println("\n--- Detailed Switch-Case Examples ---")

	// --- Example 1: Basic Switch on a Value ---
	// This is the most common form of a switch statement.
	// It compares an expression's value against a series of case values.
	fmt.Println("\n1. Basic Switch:")
	day := "Wednesday"
	switch day {
	case "Monday":
		// This block executes if day == "Monday".
		fmt.Println("It's the start of the work week.")
	case "Tuesday", "Wednesday", "Thursday":
		// This block executes if day is one of these values.
		// Multiple values can be listed in a single case, separated by commas.
		fmt.Println("It's a mid-week day.")
	case "Friday":
		// This block executes if day == "Friday".
		fmt.Println("TGIF! The weekend is near.")
	case "Saturday", "Sunday":
		// This block executes if day is a weekend day.
		fmt.Println("It's the weekend! Time to relax.")
	default:
		// The default case is optional and executes if no other case matches.
		fmt.Println("That's not a valid day of the week.")
	}

	// --- Example 2: Switch with a Short Statement ---
	// A short variable declaration can be included before the expression.
	// The variable 'today' is only in scope within this switch statement.
	fmt.Println("\n2. Switch with a short statement:")
	switch today := getDayOfWeek(); today {
	case "Saturday", "Sunday":
		fmt.Printf("It's %s! Enjoy the weekend!\n", today)
	default:
		fmt.Printf("It's %s. Back to work!\n", today)
	}
	// The following line would cause a compile error because 'today' is not defined here.
	// fmt.Println(today)

	// --- Example 3: Tagless Switch (like a cleaner if-else chain) ---
	// A switch without an expression is an alternate way to express if/else logic.
	// It is more readable than a long chain of if-else if statements.
	fmt.Println("\n3. Tagless Switch:")
	hour := time.Now().Hour() // Get the current hour (0-23)
	switch {                  // Note: no expression after 'switch'
	case hour < 12:
		// The case expression is a boolean condition.
		// This block executes if hour < 12 is true.
		fmt.Println("Good morning!")
	case hour < 17:
		// This block executes if the first case was false and hour < 17 is true.
		fmt.Println("Good afternoon!")
	default:
		// This block executes if all previous cases were false.
		fmt.Println("Good evening!")
	}

	// --- Example 4: Type Switch ---
	// A type switch is used to determine the type of an interface value.
	// It's a powerful feature for handling values of unknown types.
	fmt.Println("\n4. Type Switch:")
	checkType(42)
	checkType("hello world")
	checkType(true)
	checkType(3.14) // This will hit the default case.
	fmt.Println()
}
