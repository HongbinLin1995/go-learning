package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	go send(c)

	receive(c)

	fmt.Println("about to exit")
}

// send channel
func send(c chan<- int) {
	c <- 42
	fmt.Println("the value sent to the channel")
}

// receive channel
func receive(c <-chan int) {
	fmt.Println("the value received from the channel:", <-c)
}
