package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://go.dev/doc/security/"

func main() {
	fmt.Println("Welcome to handling webrequests in golang")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type : %T\n", response)
	// caller's responsibility to close the connection
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Resposne is :", string(dataBytes))
}
