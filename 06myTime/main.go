package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time")

	temp := time.Now()
	fmt.Println(temp)
	fmt.Println(temp.Format("01-02-2006"))
	fmt.Println(temp.Format("01-02-2006 Monday"))
	fmt.Println(temp.Format("01-02-2006 15:04:15 Monday"))

	check := temp.Format("01-02-2006 15:04:15 Monday")
	fmt.Printf("%T\n", check)

	createdDate := time.Date(2020, time.July, 15, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
