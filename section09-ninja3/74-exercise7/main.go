package main

import (
	"fmt"
)

func main() {
	x := "Hongbin Lin"

	if x == "Hongbin Lin" {
		fmt.Println(x)
	} else if x == "Kaisi Zhou" {
		fmt.Println("Kaisi Zhou", x)
	} else {
		fmt.Println("neither")
	}
}
