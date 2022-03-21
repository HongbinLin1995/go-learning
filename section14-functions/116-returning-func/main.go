package main

import (
	"fmt"
)

func main() {
	fmt.Println(bar()())

}

func test() int {
	return 123
}

// Create a function bar() which returns a func() with the int type output
func bar() func() int {
	return test
}
