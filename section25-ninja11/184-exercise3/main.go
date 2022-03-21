package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(e error) {
	fmt.Println(e)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}
	foo(c1)
}
