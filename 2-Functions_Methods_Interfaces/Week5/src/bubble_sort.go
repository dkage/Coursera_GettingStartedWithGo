package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)

	// Asks the user to input a sequence of up to 10 integers
	fmt.Print("Please enter a sequence of up to 10 integers (separated by spaces): \n")
	scanner.Scan()

	// Transforms scanner content into a string
	userInput := scanner.Text()

	// separate integers by spaces
	integersList := strings.Split(userInput, " ")

	// Convert strings to integers
	intList := make([]int, len(integersList))
	for i, str := range integersList {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Failed to convert %s to an integer: %v\n", str, err)
			return
		}
		intList[i] = num
	}

	// Print unsorted list of integers
	fmt.Println("List of integers as sent by user:", intList)

	// Calls for bubble sort and returns the sorted list
	BubbleSort(intList)
	fmt.Println("Sorted integers:", intList)
}

// BubbleSort /*
// Sorts a list of integers using the bubble sort algorithm
func BubbleSort(integersList []int) {
	for i := 0; i < len(integersList); i++ {
		for j := 0; j < len(integersList)-1; j++ {
			if integersList[j] > integersList[j+1] {
				Swap(integersList, j)
			}
		}
	}
}

// Swap /*
// Swaps two integers in a list
func Swap(integersList []int, i int) {
	integersList[i], integersList[i+1] = integersList[i+1], integersList[i]
}
