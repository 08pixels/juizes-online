package main

import (
	"fmt"
)

var numerator, denominator int

func abs(number int) int {
	if number < 0 {
		return -number
	}

	return number
}

func main() {
	fmt.Scan(&numerator, &denominator)

	for q := -1000; q <= 1000; q++ {
		for r := 0; r < abs(denominator); r++ {
			if denominator*q+r == numerator {
				fmt.Printf("%d %d\n", q, r)
				break
			}
		}
	}
}
