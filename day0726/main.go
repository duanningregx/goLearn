package main

import "fmt"

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

func main() {
	var data = []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 5, 6}
	rel := distinct(data)
	fmt.Println(rel)
}
