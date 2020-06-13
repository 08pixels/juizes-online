package main

import "fmt"

var value float64

func main() {
	fmt.Scan(&value)

	switch {
	case value < 0, value > 100:
		fmt.Println("Fora de intervalo")
	case value <= 25:
		fmt.Println("Intervalo [0,25]")
	case value <= 50:
		fmt.Println("Intervalo (25,50]")
	case value <= 75:
		fmt.Println("Intervalo (50,75]")
	default:
		fmt.Println("Intervalo (75,100]")
	}
}
