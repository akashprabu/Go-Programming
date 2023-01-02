package main

import (
	"fmt"
	"sort"
)

func main() {
	welcome := "Slices"
	fmt.Println("Welcome to ", welcome)

	var fruitList = []string{} // if we are using slices then we don't have to define the size, it is automatically expandable.
	fmt.Printf("Type of the Variable : %T\n", fruitList)

	JuiceList := []string{"Apple", "Orange", "Grape"}
	fmt.Println("Juice list is : ", JuiceList)
	fmt.Println("Length of the Juice list is : ", len(JuiceList))
	fmt.Printf("Type of the Variable : %T\n", JuiceList)

	// If we want to add value into the slice then use append method
	JuiceList = append(JuiceList, "Banana", "Guva")
	fmt.Println(JuiceList)

	// JuiceList = append(JuiceList[1:]) // It will include the starting index and till ending of the slice and overwrites it in the existing slice
	// fmt.Println(JuiceList)

	// JuiceList = append(JuiceList[1:3]) // It will include the starting index and ignore ending index of the slice and overwrites it in the existing slice
	// fmt.Println(JuiceList)

	JuiceList = append(JuiceList[:3]) // It will include the starting index and ignore ending index of the slice and overwrites it in the existing slice
	fmt.Println(JuiceList)

	highScores := make([]int, 4)
	highScores[0] = 126
	highScores[1] = 125
	highScores[2] = 124
	highScores[3] = 123
	// highScores[4] = 127 // Index out of range error
	highScores = append(highScores, 121) // But this is will work
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) // It will give you a boolean value if the slice is sorted(true) or not(false)
	sort.Ints(highScores)                       // It will sort the slice
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"react.js", "javascript", "swift", "python", "c++", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
