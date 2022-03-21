package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("ROUTINE:", runtime.NumGoroutine())

	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
	}()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("ROUTINE:", runtime.NumGoroutine())

	go func() {
		fmt.Println("hello from thing two")
		wg.Done()
	}()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("ROUTINE:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("ROUTINE:", runtime.NumGoroutine())
}
