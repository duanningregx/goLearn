package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func testListSize() {
	var list = []int{1, 2, 3, 4, 5}
	fmt.Println(len(list))
	fmt.Println(cap(list))
	list = append(list, 10)
	fmt.Println(len(list))
	fmt.Println(cap(list))
}

func testListSlice() {
	var arr = []int{1, 2, 3, 4, 5}
	brr := arr[1:3] // 左闭右开
	fmt.Println(brr)
	brr[0] = 7 // 数据切片并不会发生数据拷贝。还是指向同一个地址
	fmt.Println(arr)
}

// slice 值拷贝会拷贝一个指针
/*
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/
func listPara(arr []int) {
	arr[len(arr)-1]++
}

func existsKeys(data map[int]string) {
	if value, exists := data[0]; exists {
		fmt.Println(value)
	}
}

func main() {
	fmt.Println("cpu is" + runtime.GOARCH)
	fmt.Println(strconv.IntSize)
	testListSize()
	testListSlice()

	var a = []int{1, 2, 3, 4, 5}
	listPara(a)
	fmt.Println(a)

	var data = make(map[int]string, 10)
	data[0] = "a"
	data[1] = "b"
	existsKeys(data)
}
