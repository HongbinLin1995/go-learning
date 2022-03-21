package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n")

	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address
	fmt.Printf("\n")

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	fmt.Printf("\n")

	b := &a
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &b)
	fmt.Printf("%T\n", *b)
	fmt.Println(b)
	fmt.Println(*b) // * gives you the value stored at an address when you have the address
	fmt.Println(a)
	fmt.Printf("\n")

	*b = 43
	fmt.Println(a)
	a = 40
	fmt.Println(*b)

	// this is useful
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value

}
