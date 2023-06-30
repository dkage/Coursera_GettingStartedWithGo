package main

import "fmt"

func main() {
	var userFloat float64
	var truncatedFloat int32

	// Ask user for float number to convert
	fmt.Println("Please, enter a float number to truncate: ")

	// Prompt user for float and returns it as printed value
	fmt.Scan(&userFloat)
	fmt.Printf("You entered: %f\n", userFloat)

	// Truncate value casting to int32
	truncatedFloat = int32(userFloat)

	// Print truncated value
	fmt.Printf("Truncated float value to integer: %d\n", truncatedFloat)

}
