package main

import "fmt"

func main() {
	c := make(chan int)
	q := make(chan string)
	go send(c, q)
	receive(c, q)
}

func send(c chan<- int, q chan<- string) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	// close(c)
	q <- "About to exit"

}

func receive(c <-chan int, q <-chan string) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case o := <-q:
			fmt.Println(o)
			return
		}
	}
}
