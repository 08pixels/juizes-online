package main

import (
	"fmt"
)

var message string

func main() {
	fmt.Scan(&message)

	for i, _ := range message {
		fmt.Print(string(message[len(message)-(i+1)]))
	}
	fmt.Printf("\n")
}
