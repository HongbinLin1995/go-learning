package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%q\n", i)
		// fmt.Printf("%v\t%#x\t%#U\t%q\n", i, i, i, i)
	}
}
