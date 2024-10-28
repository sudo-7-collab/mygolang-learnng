package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to golang web requests verbs")
	PerformGetRequest()
}

func PerformGetRequest() {
	const url = "http://localhost:8000/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code : ", response.StatusCode)
	fmt.Println("Content Length is : ", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is :", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(content)
	// fmt.Println(string(content))
}
