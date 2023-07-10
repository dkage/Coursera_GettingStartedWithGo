package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type person struct {
	fname string
	lname string
}

/**
This program reads a file with names inside it and prints the names in a slice of persons
the file should have name and surname separated by spaces, one name per line, like this:

James Bond
Miss Marvel
Tony Stark
Jennifer Lawrence

The file should not have anything else written on it. Only the names and ending with a new line.
*/

func main() {
	var personName person
	var separateBySpace []string
	var lineNumber int
	fileName := bufio.NewScanner(os.Stdin)
	sliceOfPersons := make([]person, 0, 3)

	// Ask user for file name with names inside
	fmt.Printf("Please enter the filename of the text file with person names inside it:\n")
	fileName.Scan()

	namesFile, err := os.Open(fileName.Text())
	if err != nil {
		fmt.Printf("Error reading file.")
		fmt.Printf("%s", err)
		return
	}

	// Add file content to scanner, so it can read line by line, instead of reading by X amount of bytes with .Read()
	fileScanner := bufio.NewScanner(namesFile)

	// Divide fileScanner into lines
	fileScanner.Split(bufio.ScanLines)

	// Read file line by line
	fmt.Printf("Reading file...\n")
	lineNumber = 0
	for fileScanner.Scan() {
		lineNumber++

		separateBySpace = strings.Split(fileScanner.Text(), " ")
		personName.fname = separateBySpace[0]
		personName.lname = separateBySpace[1]

		sliceOfPersons = append(sliceOfPersons, personName)
	}

	fmt.Printf("Finished reading file. %d lines were read\n\n", lineNumber)
	// Sleep 3 seconds so user can read before final output
	time.Sleep(2 * time.Second)

	// Print slice of persons
	fmt.Printf("Slice of persons has a total of %d persons names:\n", len(sliceOfPersons))
	for i, person := range sliceOfPersons {
		fmt.Printf("Index %d -> First name: %s, Last name: %s\n", i, person.fname, person.lname)
	}

	// Close file before end
	err = namesFile.Close()
	if err != nil {
		fmt.Printf("Error closing file.")
		return
	}
}
