package main

import "fmt"

var number float64

func main() {

	for i := 0; i < 100; i++ {
		fmt.Scan(&number)

		if number <= 10 {
			fmt.Printf("A[%d] = %.1f\n", i, number)
		}
	}
}
