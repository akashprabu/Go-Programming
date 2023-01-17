package main

import "fmt"

func main() {
	welcome := "Functions"
	fmt.Println("Welcome to", welcome)
	greeter()
	greeterTwo()
	var a, b int = 3, 4
	result := adder(a, b)
	fmt.Printf("Sum of %d and %d is %d\n", a, b, result)
	sum, flag := proAdder(1, 2, 3, 4, 5, 10, 20)
	fmt.Println(sum, " ", flag)
}

func greeterTwo() {
	fmt.Println("Vanakkam 2 from Golang")
}

func greeter() {
	fmt.Println("Vanakkam from Golang")
}

func adder(a int, b int) int { // return tyoe in golang is called signature
	return a + b
}

func proAdder(values ...int) (int, bool) { // used when passing more interger or other datatype values to the function and this should be always be present at the end of the parameter
	sum := 0
	for _, i := range values {
		sum += i
	}
	return sum, true
}
