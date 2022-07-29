package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("hello world!"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":5656", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
