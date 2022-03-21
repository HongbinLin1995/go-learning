package main

import (
	"fmt"
)

func main() {
	x1 := []string{"James", "Bond", "Shaken, not stirred"}
	x2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(x1)
	fmt.Println(x2)
	xx := [][]string{x1, x2}
	fmt.Println(xx)

	for i, xs := range xx {
		for j, v := range xs {
			fmt.Printf("X position: %v.\t Y position: %v.\t Value: %v.\n", i, j, v)
		}
	}
}
