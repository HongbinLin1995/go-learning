package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Hongbin Lin")
	s := woo("Kaisi Zhou")
	fmt.Println(s)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello from bar,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b

}
