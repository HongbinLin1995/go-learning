package main

import "fmt"

func main() {

	f := foo()
	v := f()
	fmt.Println(v)
}

func foo() func() int {
	return func() int {
		return 123
	}
}
