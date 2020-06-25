package main

import "fmt"

var option string
var number, answer float64

func main() {

	fmt.Scan(&option)
	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Scan(&number)

			if j > i && j < (11-i) {
				answer += number
			}
		}
	}

	if option == "M" {
		answer /= 30
	}
	fmt.Printf("%.1f\n", answer)
}
