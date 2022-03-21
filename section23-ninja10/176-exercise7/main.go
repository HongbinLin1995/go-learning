package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}
	fmt.Println("Solution 1 ends")
	fmt.Printf("\n")

	x := 10
	y := 10
	c2 := gen(x, y)
	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c2)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
	fmt.Println("Solution 2 ends")
	fmt.Printf("\n")

	var wg sync.WaitGroup
	c3 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(m int) {
				for i := 0; i < 10; i++ {
					c3 <- i*10 + m
				}
				wg.Done()
			}(i)
		}
		wg.Wait()
		close(c3)
	}()
	for v := range c3 {
		fmt.Println(v)
	}
	fmt.Println("Solution 3 ends")
	fmt.Printf("\n")
}

func gen(x, y int) <-chan int {
	c := make(chan int)
	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	return c
}
