package main

import "fmt"

func main() {
	welcome := "Defer"
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
	fmt.Println("Welcome to", welcome)
	myDefer()
}

func myDefer() {
	for i := 1; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
