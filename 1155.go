package main

import "fmt"

func main() {

	result := 0.0

	for i := 1; i <= 100; i++ {
		result += (1 / float64(i))
	}

	fmt.Printf("%.2f\n", result)
}
