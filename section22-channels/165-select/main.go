package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	fmt.Println(even)
	fmt.Println(odd)
	fmt.Println(quit)
	fmt.Printf("\n")

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Printf("\n")
	fmt.Println("about to exit")
}

// send channel
func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- true
}

// receive channel
// A select is only used with channels
// A switch is used with concrete types

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("the value received from the even channel:", v)
		case v := <-odd:
			fmt.Println("the value received from the odd channel:", v)
		case v := <-quit:
			fmt.Println("the value received from the quit channel:", v)
			return
		}
	}
}
