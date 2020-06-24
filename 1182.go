package main

import "fmt"

var column int
var answer, number float64
var option string

func main() {
	fmt.Scan(&column, &option)

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&number)
			if j == column {
				answer += number
			}
		}
	}

	if option == "M" {
		answer /= 12
	}

	fmt.Printf("%.1f\n", answer)
}
