package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%T\n", x)
	fmt.Println(x)

	x = append(x[:3], x[5:]...)
	fmt.Printf("%T\n", x)
	fmt.Println(x)
}
