package main

import (
	"fmt"
	"reflect"
)

type human struct {
	name    string
	sex     bool   //
	address string //
}

func reflectTestCase() {
	var ct = human{"chen", false, "chengdu"}

	reflectType := reflect.TypeOf(ct)
	fmt.Println(reflectType.String())
	fmt.Println(reflectType.Name())
	fmt.Println(reflectType.Kind())
}

func main() {
	reflectTestCase()

	fmt.Println(" <---------------------------------------------------------------------------->")
	TestReflectFunc()
	fmt.Println(" <---------------------------------------------------------------------------->")
}
