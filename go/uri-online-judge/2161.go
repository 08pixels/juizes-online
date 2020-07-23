package main

import "fmt"

func squareRoot(repeat int) float64 {

	if repeat == 0 {
		return 0.0
	}

	return (1.0 / (6 + squareRoot(repeat-1)))
}

func main() {
	var n int
	fmt.Scan(&n)

	fmt.Printf("%.10f\n", 3+squareRoot(n))
}
