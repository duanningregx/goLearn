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

func lockTest() {
	var lock sync.RWMutex //  跟C++基本一致
	var data = 10086
	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("lock testA started", data)
	}()

	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("lock testB started", data)
	}()

	go func() {
		lock.Lock()
		defer lock.Unlock()
		data = 1
		fmt.Println("lockTestC started", data)
	}()
}

func main() {
	fmt.Println(runtime.NumCPU())
	waitGroup()
	// go语言并发编程模型？ MPG ?

	lockTest()
}
