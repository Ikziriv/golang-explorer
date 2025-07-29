package main

import "fmt"

func DemonstrateMap() {
	// Creating a map using the make function
	// This creates a map with string keys and int values.
	ages := make(map[string]int)

	// Adding key-value pairs to the map
	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 35

	fmt.Println("Map of ages:", ages)

	// Accessing a value by its key
	fmt.Println("\nBob's age is:", ages["Bob"])

	// Creating a map with a map literal
	capitals := map[string]string{
		"France":  "Paris",
		"Japan":   "Tokyo",
		"Germany": "Berlin",
	}
	fmt.Println("\nMap of capitals:", capitals)

	// Checking if a key exists
	// The second return value is a boolean that is true if the key exists.
	capital, ok := capitals["USA"]
	if ok {
		fmt.Println("\nThe capital of the USA is:", capital)
	} else {
		fmt.Println("\nThe capital of the USA is not in our map.")
	}

	// Iterating over a map using a for-range loop
	// Note: The order of iteration over a map is not guaranteed.
	fmt.Println("\nIterating over the capitals map:")
	for country, capital := range capitals {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}

	// Deleting a key-value pair from a map
	fmt.Println("\nOriginal ages map:", ages)
	delete(ages, "Charlie")
	fmt.Println("Ages map after deleting Charlie:", ages)

	// Getting the number of items in a map
	fmt.Println("\nNumber of capitals:", len(capitals))
}
