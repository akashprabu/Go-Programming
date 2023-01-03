package main

import "fmt"

func main() {
	welcome := "Maps"
	fmt.Println("Welcome to ", welcome)
	languages := make(map[string]string)
	languages["C++"] = "CPP"
	languages["RB"] = "Ruby"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"

	// Accessing the variables
	fmt.Println("List of all Languages: ", languages)
	fmt.Println("JS is the fullform of ", languages["JS"])

	// deleting a variable
	delete(languages, "JS")
	fmt.Println("List of all Languages: ", languages)
	fmt.Println("JS is the fullform of ", languages["JS"]) // will print the value as empty

	// loops
	for key, value := range languages {
		fmt.Println("Key: ", key, "  Value: ", value)
	}
}
