package main

import "fmt"

var (
	distance float64
	liters   float64
)

func main() {
	fmt.Scanf("%f\n%f", &distance, &liters)
	fmt.Printf("%.3f km/l\n", distance/liters)
}
