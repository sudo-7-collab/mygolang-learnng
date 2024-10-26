package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")

	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is", result)

	proResult, proMessage := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro result is : ", proResult)
	fmt.Println("Pro message is :", proMessage)
}

func greeter() {
	fmt.Println("Namstey from golang")
}

func greeterTwo() {
	fmt.Println("Another method")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}
	return total, "Hi Pro Adder Function"
}
