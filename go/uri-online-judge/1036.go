package main

import (
	"fmt"
	"math"
)

var (
	a, b, c float64
)

func main() {
	fmt.Scan(&a, &b, &c)

	delta := b*b - (4 * a * c)

	if delta <= 0 || a == 0 {
		fmt.Println("Impossivel calcular")
	} else {
		R1 := (-b + math.Sqrt(delta)) / (2 * a)
		R2 := (-b - math.Sqrt(delta)) / (2 * a)

		fmt.Printf("R1 = %.5f\n", R1)
		fmt.Printf("R2 = %.5f\n", R2)
	}
}
