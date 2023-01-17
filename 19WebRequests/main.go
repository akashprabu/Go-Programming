package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	welcome := "WebRequests"
	fmt.Println("Welcome to", welcome)

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The type of the response: %T\n", response)
	fmt.Println("The Response:", response)
	defer response.Body.Close() // caller's responsibility to close the request
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content: ", string(result))
}
