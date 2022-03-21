package main

import "fmt"

const (
	a          = 42
	b          = 42.42
	c          = "Hongbin Lin"
	ta int     = 42
	tb float64 = 42.42
	tc string  = "Hongbin Lin"
)

const foo = "bar"
const typedFoo string = "bar"

var s string

type MyString string

var mys MyString

func main() {
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", mys)
	fmt.Printf("%T\n", foo)
	fmt.Printf("%T\n", typedFoo)
	fmt.Printf("\n")
	s = foo // This works just fine
	fmt.Println(s)
	s = typedFoo // As does this
	fmt.Println(s)

	mys = foo // This works just fine
	fmt.Println(mys)
	// mys = typedFoo // cannot use typedFoo (type string) as type MyString in assignment

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(ta)
	fmt.Printf("%T\n", ta)
	fmt.Println(tb)
	fmt.Printf("%T\n", tb)
	fmt.Println(tc)
	fmt.Printf("%T\n", tc)
}
