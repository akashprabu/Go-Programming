package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	welcome := "SwitchCase"
	fmt.Println("Welcome to ", welcome)

	rand.Seed(time.Now().UnixNano()) // if we use this seed function outside of this vscode. It will only execute 3 because of some hardcoding of the other's application, gdb or VMS for security reasons. So we need to setup for system or VMS.
	dice := rand.Intn(6) + 1
	fmt.Println("Value of dice: ", dice)

	switch dice {
	case 1:
		fmt.Println("Value of dice is 1 you can open")
	case 2:
		fmt.Println("Value of dice is 2 you can move 2 spot")
	case 3:
		fmt.Println("Value of dice is 3 you can move 3 spot")
		fallthrough // it will execute the first underlying case only. Use multiple fallthrough for below cases to execute till last case with one roll of dice value greater than or equal to 4.
	case 4:
		fmt.Println("Value of dice is 4 you can move 4 spot")
	case 5:
		fmt.Println("Value of dice is 5 you can move 5 spot")
	case 6:
		fmt.Println("Value of dice is 6 you can move 6 spot and roll again")
	default:
		fmt.Println("Invalid!")
	}
}
