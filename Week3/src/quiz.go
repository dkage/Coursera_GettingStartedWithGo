package main

import (
	"fmt"
)

func main() {

	number1()
	number2()
	number3()
	number4()
	number5()
	number6()

}

func number1() {
	// What the following program outputs?

	fmt.Print("1.:\n")
	x := []int{4, 8, 5}
	y := -1

	for _, elt := range x {
		if elt > y {
			y = elt
		}
	}

	fmt.Print(y)
	fmt.Printf("\n\n")

	// This output 8 as it is the biggest value (the for runs through the array and sets the value of y of the current
	// element if the current element is bigger than the current value of y
}

func number2() {
	// What is printed when executed?
	fmt.Print("2.:\n")

	x := [...]int{4, 8, 5}
	y := x[0:2]
	z := x[1:3]

	y[0] = 1
	z[1] = 3

	fmt.Print(x)
	fmt.Printf("\n\n")
}

func number3() {
	// What is printed when executed?

	fmt.Print("3.:\n")

	x := [...]int{1, 2, 3, 4, 5}
	y := x[0:2]
	z := x[1:4]

	fmt.Print(len(y), cap(y), len(z), cap(z))
	fmt.Printf("\n\n")

}

func number4() {
	fmt.Print("4.:\n")

	x := map[string]int{
		"ian":    1,
		"harris": 2,
	}

	for i, j := range x {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}

	fmt.Printf("\n\n")
}

type P struct {
	x string
	y int
}

func number5() {
	fmt.Print("5.:\n")

	b := P{"x", -1}
	a := [...]P{
		{"a", 10},
		{"b", 2},
		{"c", 3},
	}

	for _, z := range a {
		if z.y > b.y {
			b = z
		}
	}

	fmt.Println(b.x)

	fmt.Printf("\n\n")

}

func number6() {
	fmt.Print("6.:\n")

	s := make([]int, 0, 3)

	s = append(s, 100)

	fmt.Println(len(s), cap(s))

}
