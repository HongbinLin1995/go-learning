package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum("Hongbin Lin", xi...)
	fmt.Println("The total(x) is", x)

	y := sum("Hongbin Lin")
	fmt.Println("The total(y) is", y)
}

func sum(s string, x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println(s)
	return sum
}

/// func (r receiver) identifier(parameter(s)) (return(s)) { code}
