package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Println(x)
	x++

	{
		y := 42
		fmt.Println(y)
	}
	// fmt.Println(y)

	fmt.Println(x)
	foo()
	fmt.Printf("\n")

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func foo() {
	fmt.Println("hello")
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
