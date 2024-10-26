package main

import "fmt"

func main() {
	fmt.Println("structs in golang")
	// no class, no inheritance, no super, no parent and no child concept in golang
	abhi := User{"Abhishek", "abhishek.com", true, 27}
	fmt.Println(abhi)
	fmt.Printf("Abhsihek's details are %+v\n", abhi)
	fmt.Printf("My name is %v and email is %v\n", abhi.Name, abhi.Email)
	abhi.GetStatus()
	abhi.NewEmail()
	fmt.Printf("My name is %v and email is %v\n", abhi.Name, abhi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test.com"
	fmt.Println("Email of this user is :", u.Email)
}