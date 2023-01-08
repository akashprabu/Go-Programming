package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	welcome := "Urls"
	fmt.Println("Welcome to", welcome)
	fmt.Println(myUrl)
	result, _ := url.Parse(myUrl)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qparam := result.Query()
	fmt.Printf("The type of the qparam: %T\n", qparam) // key and values which means maps
	fmt.Println(qparam)
	fmt.Println(qparam["coursename"])
	for key, val := range qparam {
		fmt.Println(key, " ", val)
	}

	partsofUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=hitesh",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}
