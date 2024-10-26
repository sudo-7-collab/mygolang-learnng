package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for out Pizza:")

	// comma ok or err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Type of this rating is %T", input)
}
