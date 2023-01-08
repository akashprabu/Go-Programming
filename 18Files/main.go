package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	welcome := "Files"
	fmt.Println("Welcome to", welcome)

	content := "This needs to be inserted in the file - Learning GoLang"
	file, err := os.Create("./myDemo.txt")
	checkError(err)
	// if err != nil {
	// 	panic(err)
	// }
	length, err := io.WriteString(file, content)
	checkError(err)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("The length of the file is", length)
	defer file.Close()
	ReadFile("myDemo.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadFile(filename string) {
	file, err := ioutil.ReadFile(filename)
	checkError(err)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("The Text File Contains:\n", string(file))
}
