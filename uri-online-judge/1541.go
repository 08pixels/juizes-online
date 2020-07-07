package main

import (
	"fmt"
	"math"
)

var sideA, sideB, available float64

func main() {

	for true {
		fmt.Scan(&sideA)

		if sideA == 0 {
			break
		}

		fmt.Scan(&sideB, &available)

		fmt.Printf("%d\n", int(math.Sqrt(sideA*sideB/(available/100))))
	}
}
