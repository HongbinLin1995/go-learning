package main

import "fmt"

func main() {

	// var x int
	// fmt.Println("x: ", x)
	// x++
	// fmt.Println("x: ", x)
	// i := c()
	// fmt.Println("i: ", i)
	// fmt.Printf("\n")

	// defer foo()
	// a := bar()
	// fmt.Println(a)
	// fmt.Println(1)
	// fmt.Printf("\n")

	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g. Won't run due to panic from g.")
}

func g(i int) {
	if i > 3 {
		// fmt.Println("Panicking!")
		// panic("12312312")
		// panic(fmt.Sprintf("%v", i))
		return
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

// func c() (i int) {
// 	// fmt.Println("i(in): ", i)
// 	defer func() { i++ }()
// 	return 2
// }

// func foo() {
// 	fmt.Println("foo")
// 	fmt.Printf("\n")
// }

// func bar() (i int) {
// 	defer fmt.Println("bar")
// 	fmt.Println(2)
// 	defer func() { i++ }()
// 	return 3
// }
