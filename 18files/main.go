package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file"
	file, err := os.Create("./myFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("The content of file is\n", dataByte)
	fmt.Println("The content of file is\n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
