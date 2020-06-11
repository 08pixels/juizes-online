package main

import "fmt"

var x float32
var y float32
var z float32

func ponderedMean(scoreA float32, scoreB float32, scoreC float32) float32 {
	return (scoreA*2 + scoreB*3 + scoreC*5) / 10.0
}

func main() {
	fmt.Scanf("%f\n%f\n%f", &x, &y, &z)
	fmt.Printf("MEDIA = %.1f\n", ponderedMean(x, y, z))
}
