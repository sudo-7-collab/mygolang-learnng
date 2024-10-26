package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 9
	var result string

	if loginCount < 10 {
		result = "regular user"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10 logincount"
	}

	fmt.Println(result)

	if num := 10; num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is not less than 10")
	}

	// if err != nil {

	// }
}
