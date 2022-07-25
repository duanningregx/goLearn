package main

import "fmt"
import _ "github.com/mailru/easyjson"
import "gonum.org/v1/gonum/stat"

func main() {
	fmt.Println("hello world")
	rel := []float64{1, 2, 3, 4}
	stat.Variance(rel, nil)
}
