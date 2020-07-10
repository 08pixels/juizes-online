package main

import (
	"fmt"
	"math"
)

var radius float64

func main() {
	fmt.Scanf("%f", &radius)
	fmt.Printf("VOLUME = %.3f\n", (4.0/3)*3.14159*math.Pow(radius, 3))
}
