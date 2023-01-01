package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Maths, Crypto and Rand")

	var num int = 1
	var num2 float64 = 4.5
	fmt.Println(float64(num) + num2) // type conversion

	fmt.Println("Enter the range to generate a random number: ")
	var n int
	fmt.Scanf("%d", &n)

	// using seed function to give the algorithm of the rand to pick or else sometimes it gives only 1 from the range.
	// rand.Seed(7)
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)

	// using crypto
	// crypto gives us always unique integer from the range.
	//value, _ := rand.Int(rand.Reader, big.NewInt(n))
	fmt.Println("The random value generated: ", value)
}
