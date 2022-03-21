package main

import (
	"fmt"
)

func main() {
	x := [5]int{10, 11, 12, 13, 14}
	for i, v := range x {
		fmt.Println(i, v, x[i])
	}
	fmt.Printf("%T\n", x)
}
