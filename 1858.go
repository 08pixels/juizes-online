package main

import "fmt"

var cases int
var number int

func main() {
	position := 0
	value := 101

	fmt.Scan(&cases)

	for i := 1; i <= cases; i++ {
		fmt.Scan(&number)

		if value > number {
			position = i
			value = number
		}
	}
	fmt.Println(position)
}
