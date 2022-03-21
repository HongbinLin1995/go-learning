package main

import (
	"fmt"
)

func main() {
	defer foo()
	a := bar()
	fmt.Println(a)
	fmt.Println(1)
}

func foo() {
	fmt.Println("foo")
}

func bar() int {
	defer fmt.Println("bar")
	fmt.Println(2)
	return 3
}
