package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	cb := make(chan int, 1)
	cb <- 43
	fmt.Println(<-cb)
}
