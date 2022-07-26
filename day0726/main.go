package main

import (
	"fmt"
	"time"
)

func distinct(data []int) []int {
	var dataMap = make(map[int]int)
	for _, value := range data {
		if key, ok := dataMap[value]; ok {
			dataMap[key]++
		} else {
			dataMap[value] = 1
		}
	}

	var rel []int
	for k := range dataMap {
		rel = append(rel, k)
	}
	return rel
}

type User struct {
	id      int
	name    string
	address string
	age     time.Time
}

func (u User) sayHello() {
	fmt.Println("hello,my name is " + u.name)
}

// 方法接受者传指针和struct的不同
func (u User) setName(name string) {
	u.name = name
}

func (u *User) setNamePointer(name string) {
	u.name = name
}

func main() {
	var data = []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6}
	rel := distinct(data)
	fmt.Println(rel)

	var chen = User{1, "chen", "sx", time.Now()}
	fmt.Println(chen)

	chen.setName("duan")
	fmt.Println(chen.name)
	chen.setNamePointer("duan")
	fmt.Println(chen.name)

}
