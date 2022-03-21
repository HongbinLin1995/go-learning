package main

import (
	"fmt"
)

func main() {
	x := [5]int{0, 1, 2, 3, 4}
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	y := x[1:4]
	fmt.Println(y)

	for i, v := range x {
		fmt.Println(i, v)
	}
}