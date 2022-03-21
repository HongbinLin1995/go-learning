package main

import (
	"fmt"
)

func main() {

	a := 43 % 40
	b := 43 / 40
	fmt.Println(a, b)
	fmt.Printf("\n")

	i := 1
	for {
		i++
		if i >= 100 {
			break
		}
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
