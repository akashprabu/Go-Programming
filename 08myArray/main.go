package main

import "fmt"

func main() {
	welcome := "Arrays"
	fmt.Println("Welcome to ", welcome, " in Golang")

	var fruitList [4]string // In golang we have to declare how many indexes we are going to allocate for the array
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fmt.Println("Fruit list is : ", fruitList) // by default the index which is not initialized be empty ""
	fmt.Println("Length of the fruit list is : ", len(fruitList))
	fmt.Printf("Type of the variable %T\n ", fruitList)

	var vegList = [5]string{"Carrot", "Radish", "Potato"}
	fmt.Println("Vegy list is : ", vegList) // by default the index which is not initialized be empty ""
	fmt.Println("Length of the Vegy list is : ", len(vegList))

}
