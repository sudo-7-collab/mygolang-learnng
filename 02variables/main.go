package main

import "fmt"

// Public
const LoginToken string = "gfdshsjfjhd"

func main() {
	// fmt.Print("Variables")

	var userName string = "abhishek"
	fmt.Println(userName)
	fmt.Printf("Variable is of type %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T \n", smallVal)

	var smallFloat32 float32 = 255.327376812743127
	fmt.Println(smallFloat32)
	fmt.Printf("Variable is of type %T \n", smallFloat32)

	var smallFloat64 float64 = 255.327376812743127
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of type %T \n", smallFloat64)

	// default values and some aliases
	var anotherIntVariable int
	fmt.Println(anotherIntVariable)
	fmt.Printf("Variable is of type %T \n", anotherIntVariable)

	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Variable is of type %T \n", anotherStringVariable)

	// implicit type
	var website = "learngo.com"
	fmt.Println(website)

	// no var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type %T \n", LoginToken)
}
