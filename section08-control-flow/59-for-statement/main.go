package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Printf("\n")
	i = 0
	for {
		if i >= 5 {
			break
		}
		fmt.Println(i)
		i++
	}
}
