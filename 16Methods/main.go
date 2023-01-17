package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	welcome := "Methods"
	fmt.Println("Welcome to", welcome)

	// There is no inheritence in golang; no  parent or super
	akash := User{"akash", 12, "akashprabu.cs18@bitsathy.ac.in", true}
	fmt.Println(akash)
	fmt.Printf("akash details: %+v\n", akash) // %+v gives you the details
	fmt.Printf("Name is %v and Email is %v\n", akash.Name, akash.Email)
	result := akash.solve()
	fmt.Println(akash, " ", result)
	result = akash.solve2()
	fmt.Println(akash, " ", result)
}

func (u User) solve() bool {
	u.Age = 21
	fmt.Println(u)
	return true
}

func (u *User) solve2() bool { // We are using pointer as reference here
	u.Age = 21
	fmt.Println(u)
	return true
}
