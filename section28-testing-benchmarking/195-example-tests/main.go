package main

import (
	"fmt"

	"github.com/HongbinLin1995/golang/code_samples/section28-testing-benchmarking/195-example-tests/acdc"
)

func main() {
	fmt.Println(acdc.Sum(2, 3))
	fmt.Println(acdc.Sum(2, 3, 4, 5, 6, 7, 8, 9))
}
