package main

import (
	"fmt"
	"runtime"
	"sync"
)

func add(x, y int) int {
	return x + y
}

func waitGroup() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x, y int) {
			defer wg.Done()
			fmt.Println(x + y)
		}(i, i+1)
	}
	wg.Wait()
	fmt.Println("exit waitGroup")
}

func main() {
	fmt.Println(runtime.NumCPU())
	waitGroup()
}
