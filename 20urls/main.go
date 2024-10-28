package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://go.dev:2525/doc/security?progammingLang=golang&year=2024"

func main() {
	fmt.Println("Welcome to URLs in golang")
	fmt.Println(myUrl)

	// parsing
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)
	fmt.Println(qparams["progammingLang"])
	fmt.Println(qparams["year"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "go.dev",
		Path:    "/doc/security",
		RawPath: "user=abhi",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
