package main

import (
	"fmt"
	"reflect"
)

func testReflect(x, y int) bool {
	fmt.Println(x, y)
	return true
}

func TestReflectFunc() {
	typeFunc := reflect.TypeOf(testReflect)

	for i := 0; i < typeFunc.NumIn(); i++ {
		fmt.Println("入参")
		fmt.Println(typeFunc.In(i).String())
		fmt.Println(typeFunc.In(i).Kind())
	}

	for i := 0; i < typeFunc.NumOut(); i++ {
		fmt.Println("out")
		fmt.Println(typeFunc.Out(i).String())
	}
}
