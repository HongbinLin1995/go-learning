package main

import (
	"fmt"
)

func main() {
	a := 3
	fmt.Printf("%d\t%b\t%x\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%x\n", b, b, b)
}
