package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Conversions"
	fmt.Println(welcome)

	fmt.Println("Welcome to our pizza app !!!")
	fmt.Println("Please enter your rating between 1 to 5.")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// String Conversion
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // we are using TrimSapce because the input contains for example 4\n so we are removing the \n
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Your Rating + 1 gives: ", numRating+1)
	}
	fmt.Println("Thank you for using our pizza app!!!")
}
