package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	welcome := "WebRequests"
	fmt.Println("Welcome to", welcome)
	const url = "http://localhost:8000"
	//PerformGetRequest(url + "/get")
	//PerformPostRequest(url + "/post")
	PerformPostFormRequest(url + "/postform")
}

func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest(myUrl string) {
	response, err := http.Get(myUrl)
	checkPanic(err)
	defer response.Body.Close()
	fmt.Println("The Status Code: ", response.StatusCode)
	fmt.Println("The Content length: ", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("The Content: ", string(content))
	// or
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is:", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest(myUrl string) {
	// creating a payload for sending in post request
	data := strings.NewReader(`
	{
		"coursename":"Let's go ith golang",
		"price":0,
		"platform":"learnCodeOnline.in"
	}`)
	response, err := http.Post(myUrl, "application/json", data)
	checkPanic(err)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("The Response Data:", string(content))
}

func PerformPostFormRequest(myUrl string) {
	data := url.Values{}
	data.Add("firstname", "Akashprabu")
	data.Add("lastname", "A C")
	data.Add("email", "appleakash202@gmail.com")
	response, err := http.PostForm(myUrl, data)
	checkPanic(err)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("The Response Data:", string(content))
}
