package main

import "fmt"

var number int

func main() {
	fmt.Scan(&number)

	for i := 0; i < number-1; i++ {
		fmt.Printf("Ho ")
	}
	fmt.Printf("Ho!\n")
}
