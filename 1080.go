package main

import "fmt"

var number int
var biggest, position int

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Scan(&number)

		if biggest < number {
			biggest = number
			position = i
		}
	}

	fmt.Println(biggest)
	fmt.Println(position)
}
