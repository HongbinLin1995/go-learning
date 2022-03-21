package main

import (
	"fmt"
)

var a = 43
var b int
var c int = 1

func main() {
	d := 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	foo()
}

func foo() {
	fmt.Println("fool")
}
