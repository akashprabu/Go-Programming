package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	readFile, err := os.Open("/Users/akashchandrasekaran/go/src/Go-Programming/readfromfile/data.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Printf("%T\n", fileScanner.Text())
	}

	readFile.Close()
}
