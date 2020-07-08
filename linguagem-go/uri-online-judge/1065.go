package main

import "fmt"

var number int
var even int

func main() {
	i := 0
	for i < 5 {
		fmt.Scan(&number)

		if number%2 == 0 {
			even++
		}
		i++
	}

	fmt.Printf("%d valores pares\n", even)
}
