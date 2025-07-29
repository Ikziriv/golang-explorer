package main

import "fmt"

func init() {
	// Declare some variables for comparison
	a := 10
	b := 20
	c := 10

	fmt.Printf("a = %d, b = %d, c = %d\n\n", a, b, c)

	// Equal to (==)
	// Checks if two values are equal.
	fmt.Println("a == c:", a == c) // true
	fmt.Println("a == b:", a == b) // false

	// Not equal to (!=)
	// Checks if two values are not equal.
	fmt.Println("\na != b:", a != b) // true
	fmt.Println("a != c:", a != c)   // false

	// Less than (<)
	// Checks if the left value is less than the right value.
	fmt.Println("\na < b:", a < b) // true

	// Less than or equal to (<=)
	// Checks if the left value is less than or equal to the right value.
	fmt.Println("\na <= c:", a <= c) // true
	fmt.Println("a <= b:", a <= b)   // true

	// Greater than (>)
	// Checks if the left value is greater than the right value.
	fmt.Println("\nb > a:", b > a) // true

	// Greater than or equal to (>=)
	// Checks if the left value is greater than or equal to the right value.
	fmt.Println("\nb >= a:", b >= a) // true
	fmt.Println("a >= c:", a >= c)   // true

	// Comparing strings
	// Strings are compared lexicographically (based on their byte values).
	str1 := "apple"
	str2 := "banana"
	str3 := "apple"

	fmt.Printf("\nstr1 = %s, str2 = %s, str3 = %s\n\n", str1, str2, str3)
	fmt.Println("str1 == str3:", str1 == str3) // true
	fmt.Println("str1 < str2:", str1 < str2)   // true
}
