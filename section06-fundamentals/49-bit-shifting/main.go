package main

import (
	"fmt"
)

const (
	_   = iota
	ckb = 1 << (iota * 10)
	cmb = 1 << (iota * 10)
	cgb = 1 << (iota * 10)
)

func main() {
	x := 3
	fmt.Printf("%d\t%b\n", x, x)
	lx := x << 1
	fmt.Printf("%d\t%b\n", lx, lx)
	lx2 := x << 2
	fmt.Printf("%d\t%b\n", lx2, lx2)
	rx := x >> 1
	fmt.Printf("%d\t%b\n", rx, rx)
	fmt.Printf("\n\n\n")

	kb := 1024 // 1024 -> 2^10
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Printf("%d\t\t\t%b\n", ckb, ckb)
	fmt.Printf("%d\t\t\t%b\n", cmb, cmb)
	fmt.Printf("%d\t\t%b\n", cgb, cgb)
}
