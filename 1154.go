package main

import "fmt"

func main() {
	age := 1
	counter, sum := 0, 0

	for age > 0 {
		fmt.Scan(&age)
		sum += age
		counter++
	}

	fmt.Printf("%.2f\n", float64(sum-age)/float64(counter-1))
}
