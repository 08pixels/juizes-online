package main

import "fmt"

var row int
var answer, number float64
var option string

func main() {
	fmt.Scan(&row, &option)

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&number)
			if i == row {
				answer += number
			}
		}
	}

	if option == "M" {
		answer /= 12
	}

	fmt.Printf("%.1f\n", answer)
}
