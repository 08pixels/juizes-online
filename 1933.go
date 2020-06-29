package main

import "fmt"

var cardA, cardB int

func main() {

	fmt.Scan(&cardA, &cardB)

	if cardA > cardB {
		fmt.Println(cardA)
	} else {
		fmt.Println(cardB)
	}
}
