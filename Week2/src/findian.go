package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userString string
	var scanner = bufio.NewScanner(os.Stdin)

	fmt.Print("Please, enter a string: ")

	// Prompt user for string with spaces in between and store it in userString
	scanner.Scan()
	fmt.Print("You entered: ", scanner.Text(), "\n")
	userString = strings.ToLower(scanner.Text())

	// Check if string starts with 'i'
	// Check if string contains with 'a'
	// Check if string ends with 'n'
	if strings.HasPrefix(userString, "i") &&
		strings.Contains(userString, "a") &&
		strings.HasSuffix(userString, "n") {
		fmt.Print("Found!")
	} else {
		fmt.Print("Not Found!")
	}

}
