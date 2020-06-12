package main

import (
	"fmt"
	"math"
)

var (
	pointA float64
	pointB float64
	pointX float64
	pointY float64
)

func main() {
	fmt.Scanf("%f %f", &pointA, &pointB)
	fmt.Scanf("%f %f", &pointX, &pointY)

	fmt.Printf("%.4f\n", math.Hypot(pointX-pointA, pointY-pointB))
}
