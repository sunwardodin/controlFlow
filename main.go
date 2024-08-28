package main

import "fmt"

func main() {
	fmt.Println("This is the first bit of code")
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// If statements
	if x < 42 {
		fmt.Println("x is Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("x is Equal to the meaning of life")
	} else {
		fmt.Println("x is Greater than the meaning of life")
	}

	// Logical Operators

	if x < 42 && y < 42 {
		fmt.Println("Both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}
}
