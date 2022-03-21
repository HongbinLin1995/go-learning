package main

import "fmt"

func main() {

	// Does not work
	// ca := make(chan int)
	// ca <- 41
	// fmt.Println(<-ca)

	// Work with goroutine
	cb := make(chan int)
	go func() {
		cb <- 42
		fmt.Println("test")
	}()
	fmt.Println(<-cb)
	// fmt.Println("pass")

	// Work with buffer
	cc := make(chan int, 2)
	cc <- 43
	fmt.Println(<-cc)

	// Work with buffers
	c := make(chan int, 3)
	c <- 44
	fmt.Println(<-c)

	c <- 45
	c <- 46
	c <- 47
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
