package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var userMap = make(map[string]string)

	// Get user name and address
	fmt.Print("Please enter your name: \n")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Please enter your address: \n")
	scanner.Scan()
	address := scanner.Text()

	// Add name and address to map
	userMap["name"] = name
	userMap["address"] = address

	// Convert map to JSON
	jsonOut, err := json.Marshal(userMap)

	// in case there is an error during the json.Marshal process
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Print JSON output
	fmt.Printf("\nJSON Output as byte object:\n")
	fmt.Println(jsonOut)
	fmt.Printf("\nJSON Output as string:\n")
	fmt.Println(string(jsonOut))

	// Exit
	fmt.Printf("\n\nExiting.\n")
	os.Exit(0)
}
