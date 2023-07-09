package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Week 3 - Assignment:
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should
be written as a loop.
Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through
the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the
slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to
accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop)
when the user enters the character 'X' instead of an integer.

*/

func main() {
	var userInput string
	var scanner = bufio.NewScanner(os.Stdin)
	var sli = make([]int, 0, 3) // Empty 3 size length array

	// For method call without any conditionals, so it is "infinite".
	for {

		// Asks user for integer
		fmt.Print("Type an integer, or 'X' to exit.\n")
		scanner.Scan()

		// Sets to lowercase, so if capital 'X' is typed, it works as well as 'x' to exit the script.
		userInput = strings.ToLower(scanner.Text())

		// Checks if input is an integer a valid integer. If err == nil, no errors occurred and input is INT
		if num, err := strconv.Atoi(userInput); err == nil {
			fmt.Printf("You typed an integer. Adding to the slice array.\n")

			sli = append(sli, num)
			sort.Ints(sli)

			fmt.Printf("The current slice is:\n")
			fmt.Println(sli)
			fmt.Printf("\n\n")

		} else {

			// If input is not an integer, checks if it is an X, and exit the script.
			if userInput == "x" {
				fmt.Print("X typed. Exiting application.")

				// Sleep 3 seconds so user can read the message, as some consoles will close before the user can
				// understand what is happening
				time.Sleep(3)

				// Kills script.
				os.Exit(0)
			} else {
				// If input is not an int, and not an X, ignore it as it wasn't stated what exactly the script should do
				// in this case.
				fmt.Print("\n\nInput not recognized, as it should be an INT or character 'X'. Ignoring it and " +
					"returning to loop.\n\n")

			}

		}

	}

}
