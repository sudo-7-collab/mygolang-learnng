package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Peach"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Length of fruitList array is ", len(fruitList))

	var vegList = [5]string{"Potato", "Beans", "Raddish"}
	fmt.Println("Vegetable list is ", vegList)
	fmt.Println("Length of vegList is ", len(vegList))
}
