package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer
var signals = []string{"test"}

func main() {
	// go greeter("Hello")
	// greeter("World")
	urls := []string{
		"https://google.com",
		"https://fb.com",
		"https://go.dev",
		"https://github.com",
		"https://lco.dev",
	}
	for _, web := range urls {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatalln(err)
	}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	fmt.Printf("%d is the status code of %s\n", res.StatusCode, endpoint)
}
