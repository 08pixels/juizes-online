package main

import (
	"fmt"
	"math"
)

var (
	valueA float64
	valueB float64
	valueC float64
)

const PI = 3.14159

var polygons = []string{
	"TRIANGULO",
	"CIRCULO",
	"TRAPEZIO",
	"QUADRADO",
	"RETANGULO",
}

func areaFunction(polygonType int) float64 {
	switch polygonType {
	case 0:
		base := valueA
		height := valueC
		return (base * height) / 2.0
	case 1:
		radius := valueC
		return PI * math.Pow(radius, 2)
	case 2:
		baseA := valueA
		baseB := valueB
		height := valueC
		return ((baseA + baseB) * height) / 2.0
	case 3:
		side := valueB
		return side * side
	case 4:
		sideA := valueA
		sideB := valueB
		return sideA * sideB
	}
	return 0.0
}

func main() {
	fmt.Scanf("%v %v %v", &valueA, &valueB, &valueC)

	for i, polygon := range polygons {
		fmt.Printf("%v: %.3f\n", polygon, areaFunction(i))
	}
}
