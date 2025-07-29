package main

import "fmt"

// DemonstrateStrings shows how to declare, concatenate, and inspect strings.
func DemonstrateStrings() {
	fmt.Println("--- String Examples ---")

	// --- String Declaration ---
	// Declare a string variable using the 'var' keyword and the 'string' type.
	var greeting string = "Hello"

	// Declare another string using short variable declaration (:=).
	// The compiler infers that 'name' is of type 'string'.
	name := "Go Developer"

	// --- String Concatenation ---
	// You can join (concatenate) strings together using the '+' operator.
	message := greeting + ", " + name + "!"
	fmt.Println(message) // Print the resulting combined string.

	// --- String Length ---
	// The built-in len() function returns the number of bytes in a string.
	// Note: For characters that use more than one byte (like emojis or some accented letters),
	// the length might not equal the number of visible characters.
	fmt.Printf("Length of message: %d bytes\n", len(message))

	// --- Accessing Characters (Runes) ---
	// Strings in Go are immutable, so you cannot change a character at a specific index.
	// The following line would cause a compile error: message[0] = 'h'

	// To properly iterate over characters (which are called 'runes' in Go), use a for...range loop.
	fmt.Println("Characters in 'Go':")
	for index, runeValue := range "Go" {
		// 'index' is the starting byte position of the rune.
		// 'runeValue' is the character itself (type rune is an alias for int32).
		fmt.Printf("  at index %d, the character is %c\n", index, runeValue)
	}

	// --- Multi-line Strings ---
	// You can create multi-line strings using backticks (`).
	// These are raw string literals, and they preserve all whitespace and newlines.
	multiLineString := `
	This is a
	multi-line string.
	It's useful for templates or long text blocks.`
	fmt.Println(multiLineString)

	fmt.Println() // Add a newline for better output formatting.
}
