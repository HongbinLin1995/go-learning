package main

import (
	"fmt"
)

var a = 42
var b string = `Hongbin Lin said "He is handsome!"`
var c string = `Hongbin Lin said
 "He is 

 handsome!"`

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
