package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	const gs = 100
	var incrementer int64 = 0
	var wg sync.WaitGroup
	wg.Add(gs)

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			temp := atomic.LoadInt64(&incrementer)
			fmt.Println("Count", temp)
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("\n")
	fmt.Println("Final count:", incrementer)
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())
}
