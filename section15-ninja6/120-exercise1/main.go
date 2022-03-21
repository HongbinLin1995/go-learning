package main

import (
	"fmt"
)

func main() {
	a := foo()
	b, c := bar()
	fmt.Println(a)
	fmt.Println(b, c)
}

func foo() int {
	return 1
}

func bar() (int, string) {
	return 2, "Hongbin Lin"
}
