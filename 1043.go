package main

import (
	"fmt"
	"math"
)

var a, b, c float64

func canSide(a, b, c float64) bool {
	return math.Abs(c-b) < a && c+b > a
}

func main() {
	fmt.Scan(&a, &b, &c)

	if canSide(a, b, c) && canSide(b, a, c) && canSide(c, a, b) {
		fmt.Printf("Perimetro = %.1f\n", a+b+c)
	} else {
		fmt.Printf("Area = %.1f\n", (c*(a+b))/2)
	}
}
