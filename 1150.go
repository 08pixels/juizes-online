package main

import (
	"fmt"
	"math"
)

func main() {
	numberA, numberB, sum := 0, -1000000, 0
	fmt.Scan(&numberA)

	for numberB <= numberA {
		fmt.Scan(&numberB)
	}

	i, sum := numberA+1, numberA

	for sum <= numberB {
		sum += i
		i++
	}

	fmt.Println(math.Abs(float64(i - numberA)))
}
