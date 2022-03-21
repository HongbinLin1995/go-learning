package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64

var a int8 = -128

func main() {
	x = 42
	y = 42.42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
