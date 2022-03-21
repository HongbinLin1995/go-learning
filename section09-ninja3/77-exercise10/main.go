package main

import (
	"fmt"
)

func main() {
	a := true
	fmt.Printf("true && true\t %v\n", a && true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", a || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}
