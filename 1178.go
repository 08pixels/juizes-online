package main

import "fmt"

var number float64

func main() {

	fmt.Scan(&number)
	for i := 0; i < 100; i++ {
		fmt.Printf("N[%d] = %.4f\n", i, number)
		number /= 2
	}
}
