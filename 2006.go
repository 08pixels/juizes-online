package main

import "fmt"

var target, guest int
var answer int

func main() {
	fmt.Scan(&target)

	for i := 0; i < 5; i++ {
		fmt.Scan(&guest)

		if guest == target {
			answer++
		}
	}

	fmt.Printf("%d\n", answer)
}
