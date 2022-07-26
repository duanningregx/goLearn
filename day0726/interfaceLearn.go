package main

import (
	"fmt"
	"time"
)

type animals interface {
	eat(something string)
	sleepAnimals(t time.Time)
}

type human interface {
	animals
	doCook(something string) bool
}

type cat struct {
	name string
}

func (c cat) eat(something string) {
	fmt.Println("cat is eat" + something)
}

func (c cat) sleepAnimals(t time.Time) {
	fmt.Println("cat is sleep" + t.String())
}

func printCat(ifc animals) {
	ifc.eat("food")
	ifc.sleepAnimals(time.Now())
}

type chen struct {
	name string
}

func (c chen) eat(something string) {
	fmt.Println("chen is eat" + something)
}

func (c chen) sleepAnimals(t time.Time) {
	fmt.Println("chen is sleep" + t.String())
}

func (c chen) doCook(something string) bool {
	fmt.Println("chen is doCook" + something)
	return true
}

func printChen(ifc human) {
	ifc.eat("food")
	ifc.sleepAnimals(time.Now())
	ifc.doCook("noodle")
}

func testInterface() {
	var chen = chen{"chen"}
	printChen(chen)

	var cat = cat{"sixi"}
	printCat(cat)
}
