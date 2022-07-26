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

// 继承
type student struct {
	User
	score float64
}

func switchType(data interface{}) {
	switch data.(type) {
	case int:
		rel := data.(int)
		fmt.Println(rel)
	case float32:
		rel := data.(float32)
		fmt.Println(rel)
	case string:
		rel := data.(string)
		fmt.Println(rel)
	default:
		fmt.Println("not support type")
	}
}

// CallBack 钩子函数
type CallBack func(x, y int) int

func processData(x, y int, f CallBack) int {
	return f(x, y)
}

func add(x, y int) int {
	return x + y
}

func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recovered %s\n", r)
		}
	}()
	panic(fmt.Errorf("panic"))
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

	var st = student{score: 99}
	st.id = 2
	st.age = time.Now()
	fmt.Println(st)

	var typeData = 10
	switchType(typeData)

	var x, y = 1, 2
	z := processData(x, y, add)
	fmt.Println(z)

	// TODO: learn 闭包

	panicRecover()
}
