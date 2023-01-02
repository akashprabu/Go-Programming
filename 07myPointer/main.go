package main

import "fmt"

func main() {
	welcome := "Pointers"
	fmt.Println("Welcome to class on ", welcome)

	var ptr *int
	fmt.Println("Value of the pointer is : ", ptr) // gives us nil by default
	fmt.Printf("Tyoe of the Variable %T\n", ptr)

	number := 23
	var ptr1 *int = &number
	fmt.Println("Value of the pointer is : ", *ptr1, "\nAnd the Address of the pointer is : ", ptr1)

	var ptr2 = &number
	fmt.Println("Value of the pointer is : ", *ptr2, "\nAnd the Address of the pointer is : ", ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("The new value is : ", number)

}
