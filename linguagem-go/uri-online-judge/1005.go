package main

import "fmt"

var x float32
var y float32

func ponderedMean(scoreA float32, scoreB float32) float32 {
	return (scoreA*3.5 + scoreB*7.5) / 11.0
}

func main() {
	fmt.Scanf("%f\n%f", &x, &y)
	fmt.Printf("MEDIA = %.5f\n", ponderedMean(x, y))
}
