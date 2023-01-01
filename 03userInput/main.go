package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content := "User Input"
	fmt.Println(content)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you Name: ")

	// comma err syntax

	name, err := reader.ReadString('\n')
	if err != nil { // Here we are checking the error.
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Your Name: ", name)
		fmt.Printf("Type of the variable: %T \n", name)
	}

	// comma ok syntax
	fmt.Println("Enter the Subject: ")
	subject, _ := reader.ReadString('\n') // Here we are ignoring the error by using underscore _ .
	fmt.Println("Name of the subject: ", subject)
	fmt.Printf("Type of the string: %T \n", subject)

}
