package main

import (
	"fmt"
	"os"
	"time"
)

func countDown(countC chan int, num int, finishedC chan bool) {
	if num <= 0 {
		finishedC <- true
	}

	tk := time.NewTicker(time.Second)
	defer tk.Stop()
	for {
		countC <- num
		num--
		<-tk.C
		if num <= 0 {
			finishedC <- true
			break
		}
	}
}

func abort(abortC chan bool) {
	buf := make([]byte, 1024)
	_, err := os.Stdin.Read(buf)
	if err != nil {
		return
	}
	abortC <- true
}

func listenChan(countC chan int, abortC chan bool, finishedC chan bool) {
	for {
		isExist := false
		select {
		case n := <-countC:
			fmt.Println("倒计时", n)
		case <-abortC:
			fmt.Println("abort")
			isExist = true
		case <-finishedC:
			fmt.Println("finished")
			isExist = true
		}
		if isExist {
			break
		}
	}
	fmt.Println("over")
}

func selectTest() {
	var countChan = make(chan int, 1)
	var finishedChan = make(chan bool, 1)
	var abortChan = make(chan bool, 1)
	go abort(abortChan)
	num := 5
	go countDown(countChan, num, finishedChan)
	listenChan(countChan, abortChan, finishedChan)
}
