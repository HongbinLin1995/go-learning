package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	incrementer := 0
	const gs = 100
	var m sync.Mutex
	var wg sync.WaitGroup
	wg.Add(gs)

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			// runtime.Gosched()
			v++
			incrementer = v
			fmt.Println("Count", incrementer)
			m.Unlock()
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
