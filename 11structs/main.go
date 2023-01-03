package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	welcome := "Structs"
	fmt.Println("Welcome to ", welcome)

	// There is no inheritence in golang; no  parent or super
	akash := User{"akash", 12, "akashprabu.cs18@bitsathy.ac.in", true}
	fmt.Println(akash)
	fmt.Printf("akash details: %+v\n", akash) // %+v gives you the details
	fmt.Printf("Name is %v and Email is %v", akash.Name, akash.Email)

	temp := User{
		Name:   "Hari",
		Age:    23,
		Status: true,
		Email:  "hari.cs18@bitsathy.ac.in",
	}
	fmt.Println(temp)
	fmt.Printf("temp details: %+v\n", temp) // %+v gives you the details
	fmt.Printf("Name is %v and Email is %v", temp.Name, temp.Email)
}
