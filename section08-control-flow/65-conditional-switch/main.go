package main

import (
	"fmt"
)

func main() {
	a := 3
	b := 3
	c := 4
	d := 4
	e := 15
	f := 15
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (a != b):
		fmt.Println("prints")
		fallthrough
	case (c == d):
		fmt.Println("also True, does it print")
		fallthrough
	case (7 == 9):
		fmt.Println("not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (e != f):
		fmt.Println("true 15")
		fallthrough
	default:
		fmt.Println("default")
	}

	n := "Honey"
	switch n {
	case "Money":
		fmt.Println("Money")
		fallthrough
	case "Honey":
		fmt.Println("Honey")
		fallthrough
	case "Bond":
		fmt.Println("Bond")
	}

}
