package main

import "fmt"

func main() {
	// fmt.Println("Welcome to Class on pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("pointer points to ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value of pointer is ", *ptr)
}
