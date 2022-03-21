package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

var a1 A1
var a2 A2
var b1 B1
var b2 B2
var b3 B3
var b4 B4

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Printf("\n")
	// a = b

	fmt.Printf("%T\n", a1)
	fmt.Printf("%T\n", a2)
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T\n", b2)
	fmt.Printf("%T\n", b3)
	fmt.Printf("%T\n", b4)
}
