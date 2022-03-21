package main

import (
	"fmt"
)

// we create VALUES of a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

var x int

type person struct {
	first string
	last  string
}

type foo int

var y foo

func main() {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	p := person{
		first: "Hongbin",
		last:  "Lin",
	}
	fmt.Printf("%T\n", p)
	fmt.Println(p)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}
