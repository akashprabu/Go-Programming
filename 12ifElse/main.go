package main

import (
	"fmt"
)

func main() {
	welcome := "If Else"
	fmt.Println("Welcome to ", welcome)

	loginCount := 10
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Something User"
	}
	fmt.Println(result)

	var ans string
	if 9%2 == 1 {
		ans = "Odd"
	} else {
		ans = "Even"
	}
	fmt.Println("Answer is ", ans)

	// mostly used in website development url. Assigning value from network or else where during condition checking in if else condition.
	var ans1 string
	if num := 23; num > 20 {
		ans1 = "Greater than 20"
	} else {
		ans1 = "Less than 20"
	}
	fmt.Println("Answer1 is ", ans1)
}
