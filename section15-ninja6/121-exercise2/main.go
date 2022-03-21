package main

import (
	"fmt"
)

func main() {
	x1 := []int{1, 2, 3, 4, 5}
	sum := foo(x1...)
	fmt.Println(sum)

	x2 := []int{1, 2, 3, 4, 5}
	sum2 := bar(x2)
	fmt.Println(sum2)

}

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
