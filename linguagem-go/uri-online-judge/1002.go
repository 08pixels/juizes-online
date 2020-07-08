package main

import (
	"fmt"
	"math"
)

const PI = 3.14159

var radius float64

func main() {
	fmt.Scanf("%f", &radius)
	fmt.Printf("A=%.4f\n", PI*math.Pow(radius, 2))
}
