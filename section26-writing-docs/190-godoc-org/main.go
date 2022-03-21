package main

import (
	"fmt"

	"github.com/HongbinLin1995/golang/code_samples/section26-writing-docs/190-godoc-org/mymath"
)

func main() {
	fmt.Println("2 + 3 =", mymath.Gcd(2, 3))
	fmt.Println("4 + 7 =", mymath.Gcd(4, 7))
	fmt.Println("5 + 9 =", mymath.Gcd(5, 9))
}
