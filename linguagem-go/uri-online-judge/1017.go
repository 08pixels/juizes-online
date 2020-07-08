package main

import "fmt"

var hours float64
var speed float64

func main() {
	fmt.Scanf("%f\n%f", &hours, &speed)
	fmt.Printf("%.3f\n", (hours*speed)/12)
}
