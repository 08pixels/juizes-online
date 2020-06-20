package main

import "fmt"

var number int

func main() {
	fmt.Scan(&number)

	current := 0
	next := 1

	for i := 1; i <= number; i++ {
		if i == number {
			fmt.Printf("%d\n", current)
		} else {
			fmt.Printf("%d ", current)
		}

		current, next = next, current+next
	}
}
