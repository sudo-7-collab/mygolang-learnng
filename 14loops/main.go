package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is i and value is %v\n", day)
	// }

	rougeValue := 1
	for rougeValue < 10 {

		// if rougeValue == 5 {
		// 	break
		// }

		// if rougeValue == 5 {
		// 	rougeValue++
		// 	continue
		// }

		if rougeValue == 5 {
			goto jump
		}
		fmt.Println(rougeValue)
		rougeValue++
	}

jump:
	fmt.Println("Jump using goto")
}
