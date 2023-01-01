package main

import "fmt"

// loggedIn :=   (cannot use this no var style declaration outside of methods or functions)
var LoggedTest string
var loggedIn = "dasdas"

const loggedInfo string = "sdasd"
const loggedOut = "sadadas"
const LoggedFrom string = "sdfgrtgdf" // Note: if we are declaring the variable starting letter with uppercase then it can used accross other files
const LoggedTo = "23423423"

func main() {

	fmt.Println("Variables")

	var username string = "Akashprabu A C"
	fmt.Println(username)
	fmt.Printf("Type of the vairable is %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Type of the vairable is %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Type of the vairable is %T \n", smallVal)

	var smallFloat float32 = 255.435345345345345345
	fmt.Println(smallFloat)
	fmt.Printf("Type of the vairable is %T \n", smallFloat)

	var bigFloat float64 = 255.435345345345345345
	fmt.Println(bigFloat)
	fmt.Printf("Type of the vairable is %T \n", bigFloat)

	// default values and some alias

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of the vairable is %T \n", anotherVariable)

	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("Type of the vairable is %T \n", defaultString)

	var defaultBool bool
	fmt.Println(defaultBool)
	fmt.Printf("Type of the vairable is %T \n", defaultBool)

	var defaultFloat float32
	fmt.Println(defaultFloat)
	fmt.Printf("Type of the vairable is %T \n", defaultFloat)

	// implicit type

	var name = "String"
	fmt.Println(name)
	fmt.Printf("Type of the vairable is %T \n", name)

	// no var style

	name2 := "String"
	fmt.Println(name2)
	fmt.Printf("Type of the vairable is %T \n", name2)

	fmt.Println(loggedInfo)
	fmt.Printf("Type of the vairable is %T \n", loggedInfo)
}
