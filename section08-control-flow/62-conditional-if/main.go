package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}
	a := 1
	b := 1
	if a == b {
		fmt.Println("005")
	}

	if x := 42; x == 42 {
		fmt.Println("006")
	}

	i := 1
	for {
		i++
		if i >= 100 {
			break
		}
		if j := i % 2; j != 0 {
			continue
		}
		fmt.Println(i)
	}

}
