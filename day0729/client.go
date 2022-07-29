package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func get() {
	response, err := http.Get("http://localhost:5656")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(response.Body)
	fmt.Println(response.Body)
	fmt.Println(response.Status)
}

func post() {
	str := strings.NewReader("hello server")
	response, err := http.Post("http://localhost:5656", "application/json", str)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	fmt.Println(response.Body)
}
