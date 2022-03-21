package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

const (
	c0 = 1 << iota // a == 1  (iota == 0)
	c1 = 1 << iota // b == 2  (iota == 1)
	c2 = 3         // c == 3  (iota == 2, unused)
	c3 = 1 << iota // d == 8  (iota == 3)
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Printf("\n")

	fmt.Println(c0)
	fmt.Printf("%T\n", c0)
	fmt.Println(c1)
	fmt.Printf("%T\n", c1)
	fmt.Println(c2)
	fmt.Printf("%T\n", c2)
	fmt.Println(c3)
	fmt.Printf("%T\n", c3)
}
