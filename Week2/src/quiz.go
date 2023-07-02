package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a int32
	var b int32

	a = 10
	b = 20

	avg := 2 % (a + b)

	fmt.Printf("Which of the following expressions does NOT compute the average" +
		" of two integers a and b?\n\n")

	fmt.Printf("Value of %d\n", a)
	fmt.Printf("Value of %d\n", b)
	fmt.Printf("\n\n")

	fmt.Printf("avg := (a+b)\n")
	fmt.Print(avg)

	fmt.Printf("\n\n")

	avg2 := float64(a+b) / 2
	fmt.Printf("avg := float64(a+b) / 2 \n")
	fmt.Print(avg2)

	fmt.Printf("\n\n")

	avg3 := float64(a+b) / 2.0
	fmt.Printf("avg := float64(a+b) / 2.0 \n")
	fmt.Print(avg3)

	fmt.Printf("\n\n")

	avg4 := float64(float64(a+b) / 2.0)
	fmt.Printf("avg := float64(float64(a+b) / 2.0) \n")
	fmt.Print(avg4)

	fmt.Printf("\n\n\n")

	fmt.Printf("What is printed when the following program is executed?\n" +
		"func main() {\n" +
		"    i, _ := strconv.Atoi(\"10\")\n" +
		"    y := i * 2 \n" +
		"    fmt.Println(y)\n" +
		"}\n")

	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)

	fmt.Printf("\n\n\n")

	fmt.Printf("What is printed when the following program is executed?\n" +
		"func main() {\n" +
		"    s := strings.Replace(\"ianianian\", \"ni\", \"in\", 2)\n" +
		"    fmt.Println(s)\n" +
		"}\n")

	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)

	fmt.Printf("\n\n\n")

	fmt.Printf("What is printed by this code?\n" +
		"func main() {\n" +
		"   x := 7\n\n" +
		"   switch {\n" +
		"      case x > 3:\n" +
		"         fmt.Println(\"1\")\n" +
		"      case x > 5:\n" +
		"         fmt.Println(\"2\")\n" +
		"      case x == 7:\n" +
		"         fmt.Println(\"3\")\n" +
		"      default:\n" +
		"         fmt.Println(\"4\")\n" +
		"   }\n" +
		"}\n")

	x := 7
	switch {

	case x > 3:
		fmt.Println("1")
	case x > 5:
		fmt.Println("2")
	case x == 7:
		fmt.Println("3")
	default:
		fmt.Println("4")

	}

	fmt.Printf("\n\n\n")

	fmt.Printf("What is printed by this code?\n" +
		"func main() {\n" +
		"   var xtemp int\n" +
		"   x1 := 0\n\n" +
		"   x2 := 1\n\n" +
		"" +
		"   for x := 0; x<5; x++ {\n" +
		"      xtemp = x2\n" +
		"      x2 = x1 + x2\n" +
		"      x1 = xtemp\n" +
		"   }\n" +
		"   fmt.Println(x1)\n" +
		"}\n")

	var xtemp int
	x1 := 0
	x2 := 1

	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x1 + x2
		x1 = xtemp
	}
	fmt.Println(x2)

	fmt.Printf("\n\n\n")

	fmt.Printf("This code compiles correctly.\n" +
		"func main() {\n" +
		"   var x_int int\n" +
		"   var y_int *int\n" +
		"" +
		"   z := 3\n" +
		"   y = &z\n" +
		"   x = &y\n" +
		"}\n")

	//var x_int int
	//var y_int *int
	//
	//z_int := 3
	//y_int = &z_int
	//x_int = &y_int
	fmt.Print("False, it does not print as i can't assign a pointer to a pointer\n\n\n")

}
