package main

import (
	"fmt"

	"math/rand"

	"time"
)

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
	fmt.Println()

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
	fmt.Println()

	// The comma ok idiom
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is Greater than or equal to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is less than to x which is %v\n", z, x)
	}

	// Switch statements
	// With default
	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("x is 41")
	case 42:
		fmt.Println("x is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("This is the default case for x")
	}
	switch {
	case x < 42:
		fmt.Println("x is less than 42")
	case x == 42:
		fmt.Println("x is 42")
	case x > 42:
		fmt.Println("x is greater than 42")
	default:
		fmt.Println("Another default case for 42")
	}
	// fallthrough makes the switch go to the next next even if the previous case was true and executed
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("x is 40")
		fallthrough
	case 42:
		fmt.Println("x is 40")
	default:
		fmt.Println("Just in case")
	}
	fmt.Println()

	// Select statements
	ch1 := make(chan int)
	ch2 := make(chan int)
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	// This is a go routine
	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()
	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()
	// This is a select statement
	select {
	case v1 := <-ch1:
		fmt.Println("Value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("Value from channel 2", v2)
	}
}
