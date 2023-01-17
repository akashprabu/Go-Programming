package main

import "fmt"

func main() {
	welcome := "Loops"
	fmt.Println("Welcome to", welcome)
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}
	fmt.Println(days)

	// Syntax 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// Syntax 2
	for i := range days {
		fmt.Println(days[i])
	}

	// Syntax 3
	for i, day := range days {
		fmt.Println("Index:", i, "Day:", day)
	}

	// Syntax 4 (Print Statement)
	for i, day := range days {
		fmt.Printf("Index: %v Day: %v\n", i, day)
	}

	// Syntax 5 (comma ok)
	for _, day := range days {
		fmt.Println("Day:", day)
	}

	// Syntax 6 (break)
	n := 1
	for n <= 10 {
		if n == 5 {
			break
		}
		fmt.Println("Value is ", n)
		n++
	}

	// Syntax 6 (continue)
	k := 1
	for k <= 10 {
		if k == 5 {
			k++
			continue
		}
		fmt.Println("Value is ", k)
		k++
	}

	// Syntax 6 (goto)
	l := 1
	for l <= 10 {
		if l == 5 {
			l++
			goto here
		}
		fmt.Println("Value is ", l)
		l++
	}

here:
	fmt.Println("Goto worked ")
}
