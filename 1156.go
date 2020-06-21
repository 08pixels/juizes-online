package main

import "fmt"

func main() {
	result := 1.0

	for i := 1; 2*i+1 <= 39; i++ {
		numerator := 2*i + 1
		denominator := (1 << uint32(i))

		result += float64(numerator) / float64(denominator)
	}
	fmt.Printf("%.2f\n", result)
}
